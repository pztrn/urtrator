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
    "fmt"
    "runtime"
    "strconv"
    "strings"

    // Local
    "github.com/pztrn/urtrator/datamodels"
    "github.com/pztrn/urtrator/ioq3dataparser"

    // Other
    "github.com/mattn/go-gtk/gdkpixbuf"
    "github.com/mattn/go-gtk/glib"
    "github.com/mattn/go-gtk/gtk"
)

type MainWindow struct {
    // Gamemodes.
    gamemodes map[string]string

    text string

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
    // Quick connect: nickname
    qc_nickname *gtk.Entry
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

    // Resources.
    // Pixbufs.
    // For unavailable (e.g. offline) server.
    server_offline_pic *gdkpixbuf.Pixbuf
    // For online server.
    server_online_pic *gdkpixbuf.Pixbuf
    // For private (passworded) server.
    server_private_pic *gdkpixbuf.Pixbuf
    // For public server
    server_public_pic *gdkpixbuf.Pixbuf


    // Flags.
    // Window is hidden?
    hidden bool
}

func (m *MainWindow) addToFavorites() {
    fmt.Println("Adding server to favorites...")

    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))

    server_address := ""
    if !strings.Contains(current_tab, "Favorites") {
        server_address = m.getIpFromServersList(current_tab)
    }

    // Getting server from database.
    fd := &FavoriteDialog{}
    if len(server_address) > 0 {
        servers := []datamodels.Server{}
        address := strings.Split(server_address, ":")[0]
        port := strings.Split(server_address, ":")[1]
        err1 := ctx.Database.Db.Select(&servers, ctx.Database.Db.Rebind("SELECT * FROM servers WHERE ip=? AND port=?"), address, port)
        if err1 != nil {
            fmt.Println(err1.Error())
        }
        fd.InitializeUpdate(&servers[0])
    } else {
        fd.InitializeNew()
    }
}

// Executes when delimiter for two panes is moved, to calculate VALID
// position.
func (m *MainWindow) checkMainPanePosition() {
    glib.IdleAdd(func() bool {
        m.pane_negative_position = m.window_width - m.hpane.GetPosition()
        return false
    })
}

// Executes when main window is moved or resized.
// Also calculating pane delimiter position and set it to avoid
// widgets hell :).
func (m *MainWindow) checkPositionAndSize() {
    glib.IdleAdd(func() bool {
        m.window.GetPosition(&m.window_pos_x, &m.window_pos_y)
        m.window.GetSize(&m.window_width, &m.window_height)

        m.hpane.SetPosition(m.window_width - m.pane_negative_position)
        return false
    })
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

    // Additional actions should be taken on Windows.
    if runtime.GOOS == "windows" {
        m.closeWin()
    }

    ctx.Close()
}

func (m *MainWindow) copyServerCredentialsToClipboard() {
    fmt.Println("Copying server's credentials to clipboard...")
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    server_address := m.getIpFromServersList(current_tab)
    ctx.Clipboard.CopyServerData(server_address)
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
        // Temporary disable all these modals on Linux.
        // See https://github.com/mattn/go-gtk/issues/289.
        if runtime.GOOS != "linux" {
            mbox_string := "Cannot delete server from favorites.\n\nServer isn't favorited."
            d := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, mbox_string)
            d.Response(func() {
                d.Destroy()
            })
            d.Run()
        } else {
            ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">Server isn't favorited</span></markup>"})
        }
    }

    ctx.Eventer.LaunchEvent("loadFavoriteServers", map[string]string{})
}

// Drop database data.
// ToDo: extend so we should have an ability to decide what to drop.
func (m *MainWindow) dropDatabasesData() {
    fmt.Println("Dropping database data...")
    var will_continue bool = false
    // Temporary disable all these modals on Linux.
    // See https://github.com/mattn/go-gtk/issues/289.
    if runtime.GOOS != "linux" {
        mbox_string := "You are about to drop whole database data.\n\nAfter clicking \"YES\" ALL data in database (servers, profiles, settings, etc.)\nwill be lost FOREVER. Are you sure?"
        d := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_WARNING, gtk.BUTTONS_YES_NO, mbox_string)
        d.Connect("response", func(resp *glib.CallbackContext) {
            if resp.Args(0) == 4294967287 {
                will_continue = false
            } else {
                will_continue = true
            }
            d.Destroy()
        })
        d.Run()
    } else {
        ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">Remove ~/.config/urtrator/database.sqlite3 manually!</span></markup>"})
    }

    if will_continue {
        ctx.Database.Db.MustExec("DELETE FROM servers")
        ctx.Database.Db.MustExec("DELETE FROM settings")
        ctx.Database.Db.MustExec("DELETE FROM urt_profiles")

        ctx.Eventer.LaunchEvent("loadProfiles", map[string]string{})
        ctx.Eventer.LaunchEvent("loadAllServers", map[string]string{})
        ctx.Eventer.LaunchEvent("loadFavoriteServers", map[string]string{})
    }
}

// Executes on "Edit favorite server" click.
func (m *MainWindow) editFavorite() {
    fmt.Println("Editing favorite server...")

    server_address := m.getIpFromServersList("Favorites")

    if len(server_address) > 0 {
        srv := ctx.Cache.Servers[server_address].Server
        fd := FavoriteDialog{}
        fd.InitializeUpdate(srv)
    }
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
    ctx.Eventer.LaunchEvent("loadAllServers", map[string]string{})
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
    ctx.Eventer.LaunchEvent("loadFavoriteServers", map[string]string{})
}

func (m *MainWindow) loadAllServers(data map[string]string) {
    fmt.Println("Loading all servers...")
    for _, server := range ctx.Cache.Servers {
        iter := new(gtk.TreeIter)
        ping, _ := strconv.Atoi(server.Server.Ping)

        if !server.AllServersIterSet {
            server.AllServersIter = iter
            server.AllServersIterSet = true
        } else {
            iter = server.AllServersIter
        }

        if m.all_servers_hide_offline.GetActive() && (server.Server.Players == "" && server.Server.Maxplayers == "" || ping > 9000) {
            if server.AllServersIterInList && server.AllServersIterSet {
                m.all_servers_store.Remove(iter)
                server.AllServersIterInList = false
            }
            continue
        }

        if !server.AllServersIterInList && server.AllServersIterSet {
            m.all_servers_store.Append(iter)
            server.AllServersIterInList = true
        }

        if server.Server.Name == "" && server.Server.Players == "" {
            m.all_servers_store.SetValue(iter, 0, m.server_offline_pic.GPixbuf)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"]["IP"], server.Server.Ip + ":" + server.Server.Port)
        } else {
            if ping > 9000 {
                m.all_servers_store.SetValue(iter, 0, m.server_offline_pic.GPixbuf)
            } else {
                m.all_servers_store.SetValue(iter, 0, m.server_online_pic.GPixbuf)
            }
            if server.Server.IsPrivate == "1" {
                m.all_servers_store.SetValue(iter, 1, m.server_private_pic.GPixbuf)
            } else {
                m.all_servers_store.SetValue(iter, 1, m.server_public_pic.GPixbuf)
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

func (m *MainWindow) loadFavoriteServers(data map[string]string) {
    fmt.Println("Loading favorite servers...")
    for _, server := range ctx.Cache.Servers {
        iter := new(gtk.TreeIter)
        ping, _ := strconv.Atoi(server.Server.Ping)

        if !server.FavServersIterSet {
            server.FavServersIter = iter
            server.FavServersIterSet = true
        } else {
            iter = server.FavServersIter
        }

        if m.fav_servers_hide_offline.GetActive() && (server.Server.Players == "" && server.Server.Maxplayers == "" || ping > 9000) {
            if server.FavServersIterInList {
                m.fav_servers_store.Remove(iter)
                server.FavServersIterInList = false
            }
            continue
        }

        // If server on favorites widget, but not favorited (e.g. just
        // removed from favorites) - remove it from list.
        if server.Server.Favorite != "1" && server.FavServersIterSet && server.FavServersIterInList {
            m.fav_servers_store.Remove(server.FavServersIter)
            server.FavServersIterInList = false
            server.FavServersIterSet = false
        }

        // Server isn't in favorites and wasn't previously added to widget.
        if server.Server.Favorite != "1" {
            continue
        }

        if !server.FavServersIterInList && server.FavServersIterSet {
            m.fav_servers_store.Append(iter)
            server.FavServersIterInList = true
        }

        if server.Server.Name == "" && server.Server.Players == "" {
            m.fav_servers_store.SetValue(iter, 0, m.server_offline_pic.GPixbuf)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"]["IP"], server.Server.Ip + ":" + server.Server.Port)
        } else {
            if ping > 9000 {
                m.fav_servers_store.SetValue(iter, 0, m.server_offline_pic.GPixbuf)
            } else {
                m.fav_servers_store.SetValue(iter, 0, m.server_online_pic.GPixbuf)
            }
            if server.Server.IsPrivate == "1" {
                m.fav_servers_store.SetValue(iter, 1, m.server_private_pic.GPixbuf)
            } else {
                m.fav_servers_store.SetValue(iter, 1, m.server_public_pic.GPixbuf)
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

func (m *MainWindow) loadProfiles(data map[string]string) {
    fmt.Println("Loading profiles into combobox on MainWindow")
    for i := 0; i < m.old_profiles_count; i++ {
        // ComboBox indexes are shifting on element deletion, so we should
        // detele very first element every time.
        m.profiles.Remove(0)
    }

    for _, profile := range ctx.Cache.Profiles {
        m.profiles.AppendText(profile.Profile.Name)
    }

    m.old_profiles_count = len(ctx.Cache.Profiles)
    fmt.Println("Added " + strconv.Itoa(m.old_profiles_count) + " profiles")

    m.profiles.SetActive(0)
}

func (m *MainWindow) serversUpdateCompleted(data map[string]string) {
    ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "Servers updated."})
}

func (m *MainWindow) setQuickConnectDetails(data map[string]string) {
    fmt.Println("Setting quick connect data...")
    m.qc_server_address.SetText(data["server"])
    m.qc_password.SetText(data["password"])
}

func (m *MainWindow) setToolbarLabelText(data map[string]string) {
    fmt.Println("Setting toolbar's label text...")
    if strings.Contains(data["text"], "<markup>") {
        m.toolbar_label.SetMarkup(data["text"])
    } else {
        m.toolbar_label.SetLabel(data["text"])
    }
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
    ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "URTrator is ready."})
}

func (m *MainWindow) updateOneServer() {
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    srv_address := m.getIpFromServersList(current_tab)

    if len(srv_address) > 0 {
        go ctx.Requester.UpdateOneServer(srv_address)
    }
}

// Triggered when "Update all servers" button is clicked.
func (m *MainWindow) UpdateServers() {
    ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">Updating servers...</span></markup>"})
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    fmt.Println("Updating servers on tab '" + current_tab + "'...")

    if strings.Contains(current_tab, "Servers") {
        go ctx.Requester.UpdateAllServers()
    } else if strings.Contains(current_tab, "Favorites") {
        go ctx.Requester.UpdateFavoriteServers()
    }
}
