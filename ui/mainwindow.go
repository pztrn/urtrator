// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package ui

import (
    // stdlib
    "errors"
    "fmt"
    "strconv"
    "strings"

    // Local
    "github.com/pztrn/urtrator/datamodels"
    "github.com/pztrn/urtrator/ioq3dataparser"

    // Other
    "github.com/mattn/go-gtk/glib"
    "github.com/mattn/go-gtk/gtk"
)

type MainWindow struct {
    // Gamemodes.
    gamemodes map[string]string

    // Widgets.
    // The window itself.
    window *gtk.Window
    // Vertical Box.
    vbox *gtk.VBox
    // Main menu.
    menubar *gtk.MenuBar
    // Toolbar
    toolbar *gtk.Toolbar
    // Horizontal box for main window.
    hpane *gtk.HPaned
    // Tab widget.
    tab_widget *gtk.Notebook
    // Tabs list.
    tabs map[string]*gtk.Frame
    // All servers widget.
    all_servers *gtk.TreeView
    // Favorite servers widget.
    fav_servers *gtk.TreeView
    // Statusbar.
    statusbar *gtk.Statusbar
    // Statusbar context ID.
    statusbar_context_id uint
    // Profiles combobox.
    profiles *gtk.ComboBoxText
    // Checkbox for hiding/showing offline servers in 'Servers' tab list.
    all_servers_hide_offline *gtk.CheckButton
    // Checkbox for hiding/showing offline servers in 'Favorites' tab list.
    fav_servers_hide_offline *gtk.CheckButton
    // Game launch button.
    launch_button *gtk.Button
    // Server's information.
    server_info *gtk.TreeView
    // Quick connect: server address
    qc_server_address *gtk.Entry
    // Quick connect: password
    qc_password *gtk.Entry
    // Tray icon.
    tray_icon *gtk.StatusIcon
    // Tray menu.
    tray_menu *gtk.Menu
    // Toolbar's label.
    toolbar_label *gtk.Label

    // Storages.
    // All servers store.
    all_servers_store *gtk.ListStore
    // All servers sortable store.
    all_servers_store_sortable *gtk.TreeSortable
    // Favorites
    fav_servers_store *gtk.ListStore
    // Server's information store.
    server_info_store *gtk.ListStore

    // Dialogs.
    options_dialog *OptionsDialog
    // Favorite server editing.
    favorite_dialog *FavoriteDialog

    // Other
    // Old profiles count.
    old_profiles_count int
    // Window size.
    window_width int
    window_height int
    // Window position.
    window_pos_x int
    window_pos_y int
    // Main pane delimiter position. It is calculated like:
    //
    //     window_width - pane_position
    //
    // so we will get same right pane width even if we will resize
    // main window. On resize and restore it will be set like:
    //
    //     window_width - m.pane_negative_position
    pane_negative_position int
    // Columns names for servers tabs.
    column_names map[string]string
    // Real columns positions on servers tabs.
    column_pos map[string]map[string]int


    // Flags.
    // Window is hidden?
    hidden bool
}

func (m *MainWindow) addToFavorites() {
    fmt.Println("Adding server to favorites...")

    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))

    server_address := m.getIpFromServersList(current_tab)

    // Getting server from database.
    m.favorite_dialog = &FavoriteDialog{}
    if len(server_address) > 0 {
        servers := []datamodels.Server{}
        address := strings.Split(server_address, ":")[0]
        port := strings.Split(server_address, ":")[1]
        err1 := ctx.Database.Db.Select(&servers, ctx.Database.Db.Rebind("SELECT * FROM servers WHERE ip=? AND port=?"), address, port)
        if err1 != nil {
            fmt.Println(err1.Error())
        }
        m.favorite_dialog.InitializeUpdate(&servers[0])
    } else {
        m.favorite_dialog.InitializeNew()
    }
}

// Executes when delimiter for two panes is moved, to calculate VALID
// position.
func (m *MainWindow) checkMainPanePosition() {
    m.pane_negative_position = m.window_width - m.hpane.GetPosition()
}

// Executes when main window is moved or resized.
// Also calculating pane delimiter position and set it to avoid
// widgets hell :).
func (m *MainWindow) checkPositionAndSize() {
    m.window.GetPosition(&m.window_pos_x, &m.window_pos_y)
    m.window.GetSize(&m.window_width, &m.window_height)

    m.hpane.SetPosition(m.window_width - m.pane_negative_position)
}

// Executes on URTrator shutdown.
func (m *MainWindow) Close() {
    // Save window parameters.
    ctx.Cfg.Cfg["/mainwindow/width"] = strconv.Itoa(m.window_width)
    ctx.Cfg.Cfg["/mainwindow/height"] = strconv.Itoa(m.window_height)
    ctx.Cfg.Cfg["/mainwindow/position_x"] = strconv.Itoa(m.window_pos_x)
    ctx.Cfg.Cfg["/mainwindow/position_y"] = strconv.Itoa(m.window_pos_y)
    ctx.Cfg.Cfg["/mainwindow/pane_negative_position"] = strconv.Itoa(m.pane_negative_position)

    // Saving columns sizes and positions.
    all_servers_columns := m.all_servers.GetColumns()
    for i := range all_servers_columns {
        ctx.Cfg.Cfg["/mainwindow/all_servers/" + all_servers_columns[i].GetTitle() + "_position"] = strconv.Itoa(i)
        ctx.Cfg.Cfg["/mainwindow/all_servers/" + all_servers_columns[i].GetTitle() + "_width"] = strconv.Itoa(all_servers_columns[i].GetWidth())
    }
    fav_servers_columns := m.fav_servers.GetColumns()
    for i := range fav_servers_columns {
        ctx.Cfg.Cfg["/mainwindow/fav_servers/" + fav_servers_columns[i].GetTitle() + "_position"] = strconv.Itoa(i)
        ctx.Cfg.Cfg["/mainwindow/fav_servers/" + fav_servers_columns[i].GetTitle() + "_width"] = strconv.Itoa(fav_servers_columns[i].GetWidth())
    }

    ctx.Close()
}

// Deleting server from favorites.
func (m *MainWindow) deleteFromFavorites() {
    fmt.Println("Removing server from favorites...")
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))

    server_address := m.getIpFromServersList(current_tab)

    var not_favorited bool = false
    if len(server_address) > 0 {
        if ctx.Cache.Servers[server_address].Server.Favorite == "1" {
            ctx.Cache.Servers[server_address].Server.Favorite = "0"
        } else {
            not_favorited = true
        }
    } else {
        not_favorited = true
    }

    if not_favorited {
        mbox_string := "Cannot delete server from favorites.\n\nServer isn't favorited."
        d := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, mbox_string)
        d.Response(func() {
            d.Destroy()
        })
        d.Run()
    }

    ctx.Eventer.LaunchEvent("loadFavoriteServers")
}

// Drop database data.
// ToDo: extend so we should have an ability to decide what to drop.
func (m *MainWindow) dropDatabasesData() {
    fmt.Println("Dropping database data...")
    var will_continue bool = false
    mbox_string := "You are about to drop whole database data.\n\nAfter clicking \"YES\" ALL data in database (servers, profiles, settings, etc.)\nwill be lost FOREVER. Are you sure?"
    d := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_YES_NO, mbox_string)
    d.Connect("response", func(resp *glib.CallbackContext) {
        if resp.Args(0) == 4294967287 {
            will_continue = false
        } else {
            will_continue = true
        }
    })
    d.Response(func() {
        d.Destroy()
    })
    d.Run()

    if will_continue {
        ctx.Database.Db.MustExec("DELETE FROM servers")
        ctx.Database.Db.MustExec("DELETE FROM settings")
        ctx.Database.Db.MustExec("DELETE FROM urt_profiles")

        ctx.Eventer.LaunchEvent("loadProfiles")
        ctx.Eventer.LaunchEvent("loadAllServers")
        ctx.Eventer.LaunchEvent("loadFavoriteServers")
    }
}

// Executes on "Edit favorite server" click.
func (m *MainWindow) editFavorite() {
    fmt.Println("Editing favorite server...")

    server_address := m.getIpFromServersList("Favorites")

    if len(server_address) > 0 {
        srv := ctx.Cache.Servers[server_address].Server
        m.favorite_dialog.InitializeUpdate(srv)
    }
}

func (m *MainWindow) getGameModeName(name string) string {
    val, ok := m.gamemodes[name]

    if !ok {
        return "Unknown or custom"
    }

    return val
}

// Executes when "Hide offline servers" checkbox changed it's state on
// "Servers" tab.
func (m *MainWindow) hideOfflineAllServers() {
    fmt.Println("(Un)Hiding offline servers in 'Servers' tab...")
    if m.all_servers_hide_offline.GetActive() {
        ctx.Cfg.Cfg["/serverslist/all_servers/hide_offline"] = "1"
    } else {
        ctx.Cfg.Cfg["/serverslist/all_servers/hide_offline"] = "0"
    }
    ctx.Eventer.LaunchEvent("loadAllServers")
}

// Executes when "Hide offline servers" checkbox changed it's state on
// "Favorites" tab.
func (m *MainWindow) hideOfflineFavoriteServers() {
    fmt.Println("(Un)Hiding offline servers in 'Favorite' tab...")
    if m.fav_servers_hide_offline.GetActive() {
        ctx.Cfg.Cfg["/serverslist/favorite/hide_offline"] = "1"
    } else {
        ctx.Cfg.Cfg["/serverslist/favorite/hide_offline"] = "0"
    }
    ctx.Eventer.LaunchEvent("loadFavoriteServers")
}

func (m *MainWindow) launchGame() error {
    fmt.Println("Launching Urban Terror...")

    // Getting server's name from list.
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    sel := m.all_servers.GetSelection()
    model := m.all_servers.GetModel()
    if strings.Contains(current_tab, "Favorites") {
        sel = m.fav_servers.GetSelection()
        model = m.fav_servers.GetModel()
    }
    iter := new(gtk.TreeIter)
    _ = sel.GetSelected(iter)

    // Getting server name.
    var srv_name string
    srv_name_gval := glib.ValueFromNative(srv_name)
    if strings.Contains(current_tab, "Servers") {
        model.GetValue(iter, m.column_pos["Servers"]["Name"], srv_name_gval)
    } else if strings.Contains(current_tab, "Favorites") {
        model.GetValue(iter, m.column_pos["Favorites"]["Name"], srv_name_gval)
    }
    server_name := srv_name_gval.GetString()

    // Getting server address.
    var srv_addr string
    srv_address_gval := glib.ValueFromNative(srv_addr)
    if strings.Contains(current_tab, "Servers") {
        model.GetValue(iter, m.column_pos["Servers"]["IP"], srv_address_gval)
    } else if strings.Contains(current_tab, "Favorites") {
        model.GetValue(iter, m.column_pos["Favorites"]["IP"], srv_address_gval)
    }
    srv_address := srv_address_gval.GetString()

    // Getting server's game version.
    var srv_game_ver_raw string
    srv_game_ver_gval := glib.ValueFromNative(srv_game_ver_raw)
    if strings.Contains(current_tab, "Servers") {
        model.GetValue(iter, m.column_pos["Servers"]["Version"], srv_game_ver_gval)
    } else if strings.Contains(current_tab, "Favorites") {
        model.GetValue(iter, m.column_pos["Favorites"]["Version"], srv_game_ver_gval)
    }
    srv_game_ver := srv_game_ver_gval.GetString()

    // Check for proper server name. If length == 0: server is offline,
    // we should show notification to user.
    if len(server_name) == 0 {
        var will_continue bool = false
        mbox_string := "Selected server is offline.\n\nWould you still want to launch Urban Terror?\nIt will just launch a game, without connecting to\nany server."
        m := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_YES_NO, mbox_string)
        m.Connect("response", func(resp *glib.CallbackContext) {
            if resp.Args(0) == 4294967287 {
                will_continue = false
            } else {
                will_continue = true
            }
        })
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
        if !will_continue {
            return errors.New("User declined to connect to offline server")
        }
    }

    // Getting selected profile's name.
    profile_name := m.profiles.GetActiveText()
    if strings.Contains(current_tab, "Servers") {
        // Checking profile name length. If 0 - then stop executing :)
        // This check only relevant to "Servers" tab, favorite servers
        // have profiles defined (see next).
        if len(profile_name) == 0 {
            mbox_string := "Invalid game profile selected.\n\nPlease, select profile and retry."
            m := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
            m.Response(func() {
                m.Destroy()
            })
            m.Run()
            return errors.New("User didn't select valid profile.")
        }
    } else if strings.Contains(current_tab, "Favorites") {
        // For favorite servers profile specified in favorite server
        // information have higher priority, so we just override it :)
        server := []datamodels.Server{}
        // All favorites servers should contain IP and Port :)
        ip := strings.Split(srv_address, ":")[0]
        port := strings.Split(srv_address, ":")[1]
        err := ctx.Database.Db.Select(&server, ctx.Database.Db.Rebind("SELECT * FROM servers WHERE ip=? AND port=?"), ip, port)
        if err != nil {
            fmt.Println(err.Error())
        }
        profile_name = server[0].ProfileToUse
    }

    // Getting profile data from database.
    // ToDo: cache profiles data in runtime.
    profile := []datamodels.Profile{}
    err := ctx.Database.Db.Select(&profile, ctx.Database.Db.Rebind("SELECT * FROM urt_profiles WHERE name=?"), profile_name)
    if err != nil {
        fmt.Println(err.Error())
    }

    // Check if profile version is valid for selected game server.
    if profile[0].Version != srv_game_ver {
        mbox_string := "Invalid game profile selected.\n\nSelected profile have different game version than server.\nPlease, select valid profile and retry."
        m := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
        return errors.New("User didn't select valid profile, mismatch with server's version.")
    }

    // Hey, we're ok here! :) Launch Urban Terror!
    // Crear server name from "<markup></markup>" things.
    srv_name_for_label := string([]byte(server_name)[8:len(server_name)-9])
    fmt.Println(srv_name_for_label)
    // Show great coloured label.
    m.statusbar.Push(m.statusbar_context_id, "Launching Urban Terror...")
    m.toolbar_label.SetMarkup("<markup><span foreground=\"red\" font_weight=\"bold\">Urban Terror is launched with profile </span><span foreground=\"blue\">" + profile[0].Name + "</span> <span foreground=\"red\" font_weight=\"bold\">and connected to </span><span foreground=\"orange\" font_weight=\"bold\">" + srv_name_for_label + "</span></markup>")
    m.launch_button.SetSensitive(false)
    // ToDo: handling server passwords.
    ctx.Launcher.Launch(&profile[0], srv_address, "", m.unlockInterface)

    return nil
}

func (m *MainWindow) loadAllServers() {
    fmt.Println("Loading all servers...")
    // ToDo: do it without clearing.
    for _, server := range ctx.Cache.Servers {
        iter := new(gtk.TreeIter)

        if !server.AllServersIterSet {
            server.AllServersIter = iter
            server.AllServersIterSet = true
        } else {
            iter = server.AllServersIter
        }

        if !server.AllServersIterInList {
            m.all_servers_store.Append(iter)
            server.AllServersIterInList = true
        }

        if m.all_servers_hide_offline.GetActive() && server.Server.Players == "" && server.Server.Maxplayers == "" && server.AllServersIterInList {
            m.all_servers_store.Remove(iter)
            server.AllServersIterInList = false
            continue
        }

        if server.Server.Name == "" && server.Server.Players == "" && server.Server.Maxplayers == "" {
            m.all_servers_store.SetValue(iter, 0, gtk.NewImage().RenderIcon(gtk.STOCK_NO, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"]["IP"], server.Server.Ip + ":" + server.Server.Port)
        } else {
            m.all_servers_store.SetValue(iter, 0, gtk.NewImage().RenderIcon(gtk.STOCK_OK, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
            if server.Server.IsPrivate == "1" {
                m.all_servers_store.SetValue(iter, 1, gtk.NewImage().RenderIcon(gtk.STOCK_CLOSE, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
            } else {
                m.all_servers_store.SetValue(iter, 1, gtk.NewImage().RenderIcon(gtk.STOCK_OK, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
            }
            server_name := ctx.Colorizer.Fix(server.Server.Name)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"]["Name"], server_name)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"]["Mode"], m.getGameModeName(server.Server.Gamemode))
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"]["Map"], server.Server.Map)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"]["Players"], server.Server.Players + "/" + server.Server.Maxplayers)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"]["Ping"], server.Server.Ping)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"]["Version"], server.Server.Version)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"]["IP"], server.Server.Ip + ":" + server.Server.Port)
        }
    }
}

func (m *MainWindow) loadFavoriteServers() {
    fmt.Println("Loading favorite servers...")
    for _, server := range ctx.Cache.Servers {
        if server.Server.Favorite != "1" && server.FavServersIterSet && server.FavServersIterInList {
            m.fav_servers_store.Remove(server.FavServersIter)
            server.FavServersIterInList = false
        }

        if server.Server.Favorite != "1" {
            continue
        }

        iter := new(gtk.TreeIter)

        if !server.FavServersIterSet {
            server.FavServersIter = iter
            server.FavServersIterSet = true
        } else {
            iter = server.FavServersIter
        }

        if !server.FavServersIterInList {
            m.fav_servers_store.Append(iter)
            server.FavServersIterInList = true
        }

        if m.fav_servers_hide_offline.GetActive() && server.Server.Players == "" && server.Server.Maxplayers == "" && server.FavServersIterInList {
            m.fav_servers_store.Remove(iter)
            server.FavServersIterInList = false
            continue
        }

        if server.Server.Name == "" && server.Server.Players == "" {
            m.fav_servers_store.SetValue(iter, 0, gtk.NewImage().RenderIcon(gtk.STOCK_NO, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"]["IP"], server.Server.Ip + ":" + server.Server.Port)
        } else {
            m.fav_servers_store.SetValue(iter, 0, gtk.NewImage().RenderIcon(gtk.STOCK_OK, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
            if server.Server.IsPrivate == "1" {
                m.fav_servers_store.SetValue(iter, 1, gtk.NewImage().RenderIcon(gtk.STOCK_CLOSE, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
            } else {
                m.fav_servers_store.SetValue(iter, 1, gtk.NewImage().RenderIcon(gtk.STOCK_OK, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
            }
            server_name := ctx.Colorizer.Fix(server.Server.Name)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"]["Name"], server_name)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"]["Mode"], m.getGameModeName(server.Server.Gamemode))
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"]["Map"], server.Server.Map)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"]["Players"], server.Server.Players + "/" + server.Server.Maxplayers)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"]["Ping"], server.Server.Ping)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"]["Version"], server.Server.Version)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"]["IP"], server.Server.Ip + ":" + server.Server.Port)
        }
    }
}

func (m *MainWindow) loadProfiles() {
    fmt.Println("Loading profiles into combobox on MainWindow")
    for i := 0; i < m.old_profiles_count; i++ {
        // ComboBox indexes are shifting on element deletion, so we should
        // detele very first element every time.
        m.profiles.Remove(0)
    }

    profiles := []datamodels.Profile{}
    err := ctx.Database.Db.Select(&profiles, "SELECT * FROM urt_profiles")
    if err != nil {
        fmt.Println(err.Error())
    }
    for p := range profiles {
        m.profiles.AppendText(profiles[p].Name)
    }

    m.old_profiles_count = len(profiles)
    fmt.Println("Added " + strconv.Itoa(m.old_profiles_count) + " profiles")
}

func (m *MainWindow) serversUpdateCompleted() {
    m.statusbar.Push(m.statusbar_context_id, "Servers updated.")
    m.toolbar_label.SetLabel("Servers updated.")
}

func (m *MainWindow) showHide() {
    if m.hidden {
        m.window.Show()
        m.hidden = false
        // Set window position on restore. Window loosing it on
        // multimonitor configurations.
        m.window.Move(m.window_pos_x, m.window_pos_y)
    } else {
        m.window.Hide()
        m.hidden = true
    }
}

func (m *MainWindow) showServerInformation() {
    fmt.Println("Showing server's information...")
}

func (m *MainWindow) showShortServerInformation() {
    fmt.Println("Server selection changed, updating server's information widget...")
    m.server_info_store.Clear()
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    srv_address := m.getIpFromServersList(current_tab)

    // Getting server information from cache.
    if len(srv_address) > 0 && ctx.Cache.Servers[srv_address].Server.Players != "" {
        server_info := ctx.Cache.Servers[srv_address].Server
        parsed_general_data := ioq3dataparser.ParseInfoToMap(server_info.ExtendedConfig)
        parsed_players_info := ioq3dataparser.ParsePlayersInfoToMap(server_info.PlayersInfo)
        // Append to treeview generic info first. After appending it
        // will be deleted from map.

        iter := new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, "<markup><span font_weight=\"bold\">GENERAL INFO</span></markup>")

        // Server's name.
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, "Server's name")
        m.server_info_store.SetValue(iter, 1, ctx.Colorizer.Fix(parsed_general_data["sv_hostname"]))
        delete(parsed_general_data, "sv_hostname")

        // Game version.
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, "Game version")
        m.server_info_store.SetValue(iter, 1, parsed_general_data["g_modversion"])
        delete(parsed_general_data, "g_modversion")

        // Players.
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, "Players")
        m.server_info_store.SetValue(iter, 1, server_info.Players + " of " + parsed_general_data["sv_maxclients"])
        delete(parsed_general_data, "sv_maxclients")

        // Ping
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, "Ping")
        m.server_info_store.SetValue(iter, 1, server_info.Ping + " ms")

        // Game mode
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, "Game mode")
        m.server_info_store.SetValue(iter, 1, m.gamemodes[server_info.Gamemode])

        // Map name
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, "Current map")
        m.server_info_store.SetValue(iter, 1, server_info.Map)

        // Private or public?
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, "Passworded")
        passworded_status := "<markup><span foreground=\"green\">No</span></markup>"
        if server_info.IsPrivate == "1" {
            passworded_status = "<markup><span foreground=\"red\">Yes</span></markup>"
        }
        m.server_info_store.SetValue(iter, 1, passworded_status)

        // Just a separator.
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)

        // Players information
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, "<markup><span font_weight=\"bold\">PLAYERS</span></markup>")

        for _, value := range parsed_players_info {
            iter = new(gtk.TreeIter)
            nick := ctx.Colorizer.Fix(value["nick"])
            m.server_info_store.Append(iter)
            m.server_info_store.SetValue(iter, 0, nick)
            m.server_info_store.SetValue(iter, 1, "(frags: " + value["frags"] + " | ping: " + value["ping"] + ")")
        }

        // Just a separator.
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)

        // Other parameters :).
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, "<markup><span font_weight=\"bold\">OTHER PARAMETERS</span></markup>")

        for key, value := range parsed_general_data {
            iter = new(gtk.TreeIter)
            m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, key)
        m.server_info_store.SetValue(iter, 1, value)
        }
    }
}

// Show tray menu on right-click on tray icon.
func (m *MainWindow) showTrayMenu(cbx *glib.CallbackContext) {
    m.tray_menu.Popup(nil, nil, gtk.StatusIconPositionMenu, m.tray_icon,  uint(cbx.Args(0)), uint32(cbx.Args(1)))
}

// Unlocking interface after game shut down.
func (m *MainWindow) unlockInterface() {
    m.launch_button.SetSensitive(true)
    m.statusbar.Push(m.statusbar_context_id, "URTrator is ready.")
    m.toolbar_label.SetLabel("URTrator is ready.")
}

func (m *MainWindow) updateOneServer() {
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    sel := m.all_servers.GetSelection()
    model := m.all_servers.GetModel()
    iter := new(gtk.TreeIter)
    _ = sel.GetSelected(iter)

    // Getting server address.
    var srv_addr string
    srv_address_gval := glib.ValueFromNative(srv_addr)
    model.GetValue(iter, m.column_pos["Servers"]["IP"], srv_address_gval)
    if strings.Contains(current_tab, "Favorites") {
        sel = m.fav_servers.GetSelection()
        model = m.fav_servers.GetModel()
        model.GetValue(iter, m.column_pos["Favorites"]["IP"], srv_address_gval)
    }

    srv_address := srv_address_gval.GetString()

    if len(srv_address) > 0 {
        go ctx.Requester.UpdateOneServer(srv_address)
    }
}

// Triggered when "Update all servers" button is clicked.
func (m *MainWindow) UpdateServers() {
    m.statusbar.Push(m.statusbar_context_id, "Updating servers...")
    m.toolbar_label.SetMarkup("<markup><span foreground=\"red\" font_weight=\"bold\">Updating servers...</span></markup>")
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    fmt.Println("Updating servers on tab '" + current_tab + "'...")

    if strings.Contains(current_tab, "Servers") {
        go ctx.Requester.UpdateAllServers()
    } else if strings.Contains(current_tab, "Favorites") {
        go ctx.Requester.UpdateFavoriteServers()
    }
}
