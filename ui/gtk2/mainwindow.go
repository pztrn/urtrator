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
    "sort"
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
    // Checkbox for hiding/showing passworded servers in 'Servers' tab list.
    all_servers_hide_private *gtk.CheckButton
    // Combobox for filtering server's versions.
    all_servers_version *gtk.ComboBoxText
    // Combobox for filtering by gamemode.
    all_servers_gamemode *gtk.ComboBoxText
    // Checkbox for hiding/showing offline servers in 'Favorites' tab list.
    fav_servers_hide_offline *gtk.CheckButton
    // Checkbox for hiding/showing passworded servers in 'Favorites' tab list.
    fav_servers_hide_private *gtk.CheckButton
    // Combobox for filtering server's versions.
    fav_servers_version *gtk.ComboBoxText
    // Combobox for filtering by gamemode.
    fav_servers_gamemode *gtk.ComboBoxText
    // Game launch button.
    launch_button *gtk.Button
    // Server's main information.
    server_info *gtk.TreeView
    // Players information.
    players_info *gtk.TreeView
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
    // Players information store.
    players_info_store *gtk.ListStore

    // Dialogs.
    options_dialog *OptionsDialog
    server_cvars_dialog *ServerCVarsDialog

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
    // Application is initialized?
    initialized bool
    // Window is hidden?
    hidden bool
    // Use other's tab information?
    // Used when user changed active tab, to show information about
    // server which is selected on activated tab.
    use_other_servers_tab bool
    // Does servers updating already in progress?
    // This helps to prevent random crashes when more than one
    // updating process in progress.
    servers_already_updating bool
}

func (m *MainWindow) addToFavorites() {
    fmt.Println("Adding server to favorites...")

    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))

    server_address := ""
    if !strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
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

func (m *MainWindow) allServersGamemodeFilterChanged() {
    ctx.Cfg.Cfg["/serverslist/all_servers/gamemode"] = strconv.Itoa(m.all_servers_gamemode.GetActive())
    ctx.Eventer.LaunchEvent("loadAllServers", nil)
}

func (m *MainWindow) allServersVersionFilterChanged() {
    ctx.Cfg.Cfg["/serverslist/all_servers/version"] = strconv.Itoa(m.all_servers_version.GetActive())
    ctx.Eventer.LaunchEvent("loadAllServers", nil)
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
            mbox_string := ctx.Translator.Translate("Cannot delete server from favorites.\n\nServer isn't favorited.", nil)
            d := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, mbox_string)
            d.Response(func() {
                d.Destroy()
            })
            d.Run()
        } else {
            ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">" + ctx.Translator.Translate("Server isn't favorited", nil) + "</span></markup>"})
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
        mbox_string := ctx.Translator.Translate("You are about to drop whole database data.\n\nAfter clicking \"YES\" ALL data in database (servers, profiles, settings, etc.)\nwill be lost FOREVER. Are you sure?", nil)
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
        ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">" + ctx.Translator.Translate("Remove ~/.config/urtrator/database.sqlite3 manually!", nil) + "</span></markup>"})
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

func (m *MainWindow) favServersGamemodeFilterChanged() {
    ctx.Cfg.Cfg["/serverslist/favorite/gamemode"] = strconv.Itoa(m.fav_servers_gamemode.GetActive())
    ctx.Eventer.LaunchEvent("loadFavoriteServers", nil)
}

func (m *MainWindow) favServersVersionFilterChanged() {
    ctx.Cfg.Cfg["/serverslist/favorite/version"] = strconv.Itoa(m.fav_servers_version.GetActive())
    ctx.Eventer.LaunchEvent("loadFavoriteServers", nil)
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

// Executes when "Hide passworded servers" checkbox changed it's state on
// "Servers" tab.
func (m *MainWindow) hidePrivateAllServers() {
    fmt.Println("(Un)Hiding private servers in 'Servers' tab...")
    if m.all_servers_hide_private.GetActive() {
        ctx.Cfg.Cfg["/serverslist/all_servers/hide_private"] = "1"
    } else {
        ctx.Cfg.Cfg["/serverslist/all_servers/hide_private"] = "0"
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

// Executes when "Hide passworded servers" checkbox changed it's state on
// "Favorites" tab.
func (m *MainWindow) hidePrivateFavoriteServers() {
    fmt.Println("(Un)Hiding private servers in 'Favorite' tab...")
    if m.all_servers_hide_private.GetActive() {
        ctx.Cfg.Cfg["/serverslist/favorite/hide_private"] = "1"
    } else {
        ctx.Cfg.Cfg["/serverslist/favorite/hide_private"] = "0"
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

        // Hide offline servers?
        if m.all_servers_hide_offline.GetActive() && (server.Server.Players == "" && server.Server.Maxplayers == "" || ping > 9000) {
            if server.AllServersIterInList && server.AllServersIterSet {
                m.all_servers_store.Remove(iter)
                server.AllServersIterInList = false
            }
            continue
        }

        // Hide private servers?
        if m.all_servers_hide_private.GetActive() && server.Server.IsPrivate == "1" {
            if server.AllServersIterInList && server.AllServersIterSet {
                m.all_servers_store.Remove(iter)
                server.AllServersIterInList = false
            }
            continue
        }

        // Hide servers that using different version than selected in
        // filter?
        if m.all_servers_version.GetActiveText() != ctx.Translator.Translate("All versions", nil) && m.all_servers_version.GetActiveText() != server.Server.Version {
            if server.AllServersIterInList && server.AllServersIterSet {
                m.all_servers_store.Remove(iter)
                server.AllServersIterInList = false
            }
            continue
        }

        // Hide servers that using different gamemode than selected in
        // filter?
        gm_int_as_str := strconv.Itoa(m.all_servers_gamemode.GetActive())
        if m.all_servers_gamemode.GetActiveText() != ctx.Translator.Translate("All gamemodes", nil) && gm_int_as_str != server.Server.Gamemode {
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
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"][ctx.Translator.Translate("IP", nil)], server.Server.Ip + ":" + server.Server.Port)
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
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"][ctx.Translator.Translate("Name", nil)], server_name)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"][ctx.Translator.Translate("Mode", nil)], m.getGameModeName(server.Server.Gamemode))
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"][ctx.Translator.Translate("Map", nil)], server.Server.Map)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"][ctx.Translator.Translate("Players", nil)], server.Server.Players + "/" + server.Server.Bots + "/" + server.Server.Maxplayers)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"][ctx.Translator.Translate("Ping", nil)], server.Server.Ping)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"][ctx.Translator.Translate("Version", nil)], server.Server.Version)
            m.all_servers_store.SetValue(iter, m.column_pos["Servers"][ctx.Translator.Translate("IP", nil)], server.Server.Ip + ":" + server.Server.Port)
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

        // Hide offline servers?
        if m.fav_servers_hide_offline.GetActive() && (server.Server.Players == "" && server.Server.Maxplayers == "" || ping > 9000) {
            if server.FavServersIterInList {
                m.fav_servers_store.Remove(iter)
                server.FavServersIterInList = false
            }
            continue
        }

        // Hide private servers?
        if m.fav_servers_hide_private.GetActive() && server.Server.IsPrivate == "1" {
            if server.FavServersIterInList && server.FavServersIterSet {
                m.fav_servers_store.Remove(iter)
                server.FavServersIterInList = false
            }
            continue
        }

        // Hide servers that using different version than selected in
        // filter?
        if m.fav_servers_version.GetActiveText() != ctx.Translator.Translate("All versions", nil) && m.fav_servers_version.GetActiveText() != server.Server.Version {
            if server.FavServersIterInList && server.FavServersIterSet {
                m.fav_servers_store.Remove(iter)
                server.FavServersIterInList = false
            }
            continue
        }

        // Hide servers that using different gamemode than selected in
        // filter?
        gm_int_as_str := strconv.Itoa(m.fav_servers_gamemode.GetActive())
        if m.fav_servers_gamemode.GetActiveText() != ctx.Translator.Translate("All gamemodes", nil) && gm_int_as_str != server.Server.Gamemode {
            if server.FavServersIterInList && server.FavServersIterSet {
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
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"][ctx.Translator.Translate("IP", nil)], server.Server.Ip + ":" + server.Server.Port)
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
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"][ctx.Translator.Translate("Name", nil)], server_name)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"][ctx.Translator.Translate("Mode", nil)], m.getGameModeName(server.Server.Gamemode))
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"][ctx.Translator.Translate("Map", nil)], server.Server.Map)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"][ctx.Translator.Translate("Players", nil)], server.Server.Players + "/" + server.Server.Bots + "/" + server.Server.Maxplayers)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"][ctx.Translator.Translate("Ping", nil)], server.Server.Ping)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"][ctx.Translator.Translate("Version", nil)], server.Server.Version)
            m.fav_servers_store.SetValue(iter, m.column_pos["Favorites"][ctx.Translator.Translate("IP", nil)], server.Server.Ip + ":" + server.Server.Port)
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

func (m *MainWindow) tabChanged() {
    if !m.initialized {
        return
    }

    fmt.Println("Active tab changed...")
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    fmt.Println(current_tab)

    m.use_other_servers_tab = true
    if strings.Contains(current_tab, ctx.Translator.Translate("Servers", nil)) {
        m.fav_servers.Emit("cursor-changed")
    } else if strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
        m.all_servers.Emit("cursor-changed")
    }
    m.use_other_servers_tab = false
}

func (m *MainWindow) serversUpdateCompleted(data map[string]string) {
    ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": ctx.Translator.Translate("Servers updated.", nil)})
    // Trigger "selection-changed" events on currently active tab's
    // servers list.
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))

    if strings.Contains(current_tab, ctx.Translator.Translate("Servers", nil)) {
        m.all_servers.Emit("cursor-changed")
    } else if strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
        m.fav_servers.Emit("cursor-changed")
    }

    m.servers_already_updating = false

}

func (m *MainWindow) setQuickConnectDetails(data map[string]string) {
    fmt.Println("Setting quick connect data...")
    m.qc_server_address.SetText(data["server"])
    m.qc_password.SetText(data["password"])
}

func (m *MainWindow) setToolbarLabelText(data map[string]string) {
    fmt.Println("Setting toolbar's label text...")
    if strings.Contains(data["text"], "<markup>") {
        fmt.Println("With markup")
        m.toolbar_label.SetMarkup(data["text"])
    } else {
        fmt.Println("Without markup")
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

func (m *MainWindow) showServerCVars() {
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    if m.use_other_servers_tab {
        if strings.Contains(current_tab, ctx.Translator.Translate("Servers", nil)) {
            current_tab = "Favorites"
        } else if strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
            current_tab = "Servers"
        }
    }
    srv_address := m.getIpFromServersList(current_tab)
    if len(srv_address) > 0 {
        m.server_cvars_dialog.Initialize(m.window, srv_address)
    }
}

func (m *MainWindow) showShortServerInformation() {
    fmt.Println("Server selection changed, updating server's information widget...")
    m.server_info_store.Clear()
    m.players_info_store.Clear()
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    if m.use_other_servers_tab {
        if strings.Contains(current_tab, ctx.Translator.Translate("Servers", nil)) {
            current_tab = "Favorites"
        } else if strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
            current_tab = "Servers"
        }
    }
    srv_address := m.getIpFromServersList(current_tab)

    // Getting server information from cache.
    if len(srv_address) > 0 && ctx.Cache.Servers[srv_address].Server.Players != "" {
        server_info := ctx.Cache.Servers[srv_address].Server
        parsed_general_data := ioq3dataparser.ParseInfoToMap(server_info.ExtendedConfig)
        parsed_players_info := ioq3dataparser.ParsePlayersInfoToMap(server_info.PlayersInfo)
        // Append to treeview generic info first. After appending it
        // will be deleted from map.

        // Server's name.
        iter := new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, ctx.Translator.Translate("Server's name", nil))
        m.server_info_store.SetValue(iter, 1, ctx.Colorizer.Fix(parsed_general_data["sv_hostname"]))
        delete(parsed_general_data, "sv_hostname")

        // Game version.
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, ctx.Translator.Translate("Game version", nil))
        m.server_info_store.SetValue(iter, 1, parsed_general_data["g_modversion"])
        delete(parsed_general_data, "g_modversion")

        // Players.
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, ctx.Translator.Translate("Players", nil))
        m.server_info_store.SetValue(iter, 1, server_info.Players + " of " + parsed_general_data["sv_maxclients"] + " (" + server_info.Bots + " bots)")
        delete(parsed_general_data, "sv_maxclients")

        // Ping
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, ctx.Translator.Translate("Ping", nil))
        m.server_info_store.SetValue(iter, 1, server_info.Ping + " ms")

        // Game mode
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, ctx.Translator.Translate("Game mode", nil))
        m.server_info_store.SetValue(iter, 1, m.gamemodes[server_info.Gamemode])

        // Map name
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, ctx.Translator.Translate("Current map", nil))
        m.server_info_store.SetValue(iter, 1, server_info.Map)

        // Private or public?
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, ctx.Translator.Translate("Passworded", nil))
        passworded_status := "<markup><span foreground=\"green\">" + ctx.Translator.Translate("No", nil) + "</span></markup>"
        if server_info.IsPrivate == "1" {
            passworded_status = "<markup><span foreground=\"red\">" + ctx.Translator.Translate("Yes", nil) + "</span></markup>"
        }
        m.server_info_store.SetValue(iter, 1, passworded_status)

        // Sorting keys of map.
        players_map_keys := make([]string, 0, len(parsed_players_info))
        for k := range parsed_players_info {
            // ToDo: figure out how to do this properly without
            // append().
            players_map_keys = append(players_map_keys, k)
        }

        sort.Strings(players_map_keys)

        for k := range players_map_keys {
            iter = new(gtk.TreeIter)
            nick := ctx.Colorizer.Fix(parsed_players_info[players_map_keys[k]]["nick"])
            m.players_info_store.Append(iter)
            m.players_info_store.SetValue(iter, 0, nick)
            m.players_info_store.SetValue(iter, 1, parsed_players_info[players_map_keys[k]]["frags"])
            m.players_info_store.SetValue(iter, 2, parsed_players_info[players_map_keys[k]]["ping"])
        }

        /*
        // Just a separator.
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)

        // Other parameters :).
        iter = new(gtk.TreeIter)
        m.server_info_store.Append(iter)
        m.server_info_store.SetValue(iter, 0, "<markup><span font_weight=\"bold\">OTHER PARAMETERS</span></markup>")

        // Sort it!
        general_data_keys := make([]string, 0, len(parsed_general_data))
        for k := range parsed_general_data {
            general_data_keys = append(general_data_keys, k)
        }

        sort.Strings(general_data_keys)

        for k := range general_data_keys {
            iter = new(gtk.TreeIter)
            m.server_info_store.Append(iter)
            m.server_info_store.SetValue(iter, 0, general_data_keys[k])
            m.server_info_store.SetValue(iter, 1, parsed_general_data[general_data_keys[k]])
        }
        */
    }
}

// Show tray menu on right-click on tray icon.
func (m *MainWindow) showTrayMenu(cbx *glib.CallbackContext) {
    m.tray_menu.Popup(nil, nil, gtk.StatusIconPositionMenu, m.tray_icon,  uint(cbx.Args(0)), uint32(cbx.Args(1)))
}

// Unlocking interface after game shut down.
func (m *MainWindow) unlockInterface() {
    m.launch_button.SetSensitive(true)
    ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": ctx.Translator.Translate("URTrator is ready.", nil)})
}

func (m *MainWindow) updateOneServer() {
    if m.servers_already_updating {
        return
    }
    m.servers_already_updating = true

    ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">" + ctx.Translator.Translate("Updating selected server...", nil) + "</span></markup>"})
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    srv_address := m.getIpFromServersList(current_tab)

    if len(srv_address) > 0 {
        go ctx.Requester.UpdateOneServer(srv_address)
    }
}

// Triggered when "Update all servers" button is clicked.
func (m *MainWindow) UpdateServers() {
    if m.servers_already_updating {
        return
    }
    m.servers_already_updating = true

    ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">" + ctx.Translator.Translate("Updating servers...", nil) + "</span></markup>"})
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    fmt.Println("Updating servers on tab '" + current_tab + "'...")

    if strings.Contains(current_tab, ctx.Translator.Translate("Servers", nil)) {
        go ctx.Requester.UpdateAllServers(false)
    } else if strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
        go ctx.Requester.UpdateFavoriteServers()
    }
}

func (m *MainWindow) UpdateServersEventHandler(data map[string]string) {
    ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">" + ctx.Translator.Translate("Updating servers...", nil) + "</span></markup>"})

    go ctx.Requester.UpdateAllServers(true)
}
