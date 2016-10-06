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
    "encoding/base64"
    "errors"
    "fmt"
    "runtime"
    "strconv"
    "strings"

    // Local
    "github.com/pztrn/urtrator/common"
    "github.com/pztrn/urtrator/datamodels"

    // Other
    "github.com/mattn/go-gtk/gdkpixbuf"
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


    // Flags.
    // Window is hidden?
    hidden bool
}

func (m *MainWindow) addToFavorites() {
    fmt.Println("Adding server to favorites...")

    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))

    // Getting server's address from list.
    sel := m.all_servers.GetSelection()
    model := m.all_servers.GetModel()
    if strings.Contains(current_tab, "Favorites") {
        // Getting server's address from list.
        sel = m.fav_servers.GetSelection()
        model = m.fav_servers.GetModel()
    }
    iter := new(gtk.TreeIter)
    _ = sel.GetSelected(iter)

    // Getting server address.
    var srv_addr string
    srv_addr_gval := glib.ValueFromNative(srv_addr)
    model.GetValue(iter, 7, srv_addr_gval)
    server_address := srv_addr_gval.GetString()

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

    ctx.Close()
}

// Deleting server from favorites.
func (m *MainWindow) deleteFromFavorites() {
    fmt.Println("Removing server from favorites...")
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))

    // Assuming that deletion was called from "Servers" tab by default.
    sel := m.all_servers.GetSelection()
    model := m.all_servers.GetModel()
    if strings.Contains(current_tab, "Favorites") {
        // Getting server's address from list.
        sel = m.fav_servers.GetSelection()
        model = m.fav_servers.GetModel()
    }

    iter := new(gtk.TreeIter)
    _ = sel.GetSelected(iter)

    // Getting server address.
    var srv_addr string
    srv_addr_gval := glib.ValueFromNative(srv_addr)
    model.GetValue(iter, 7, srv_addr_gval)
    server_address := srv_addr_gval.GetString()

    var not_favorited bool = false
    if len(server_address) > 0 {
        address := strings.Split(server_address, ":")[0]
        port := strings.Split(server_address, ":")[1]
        srv := []datamodels.Server{}
        err := ctx.Database.Db.Select(&srv, ctx.Database.Db.Rebind("SELECT * FROM servers WHERE ip=? AND port=?"), address, port)
        if err != nil {
            fmt.Println(err.Error())
        }
        if srv[0].Favorite == "1" {
            ctx.Database.Db.MustExec(ctx.Database.Db.Rebind("UPDATE servers SET favorite='0' WHERE ip=? AND port=?"), address, port)
            ctx.Eventer.LaunchEvent("loadFavoriteServers")
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

    sel := m.fav_servers.GetSelection()
    model := m.fav_servers.GetModel()
    iter := new(gtk.TreeIter)
    _ = sel.GetSelected(iter)

    // Getting server address.
    var srv_addr string
    srv_addr_gval := glib.ValueFromNative(srv_addr)
    model.GetValue(iter, 7, srv_addr_gval)
    server_address := srv_addr_gval.GetString()

    if len(server_address) > 0 {
        address := strings.Split(server_address, ":")[0]
        port := strings.Split(server_address, ":")[1]
        srv := []datamodels.Server{}
        err := ctx.Database.Db.Select(&srv, ctx.Database.Db.Rebind("SELECT * FROM servers WHERE ip=? AND port=?"), address, port)
        if err != nil {
            fmt.Println(err.Error())
        }
        m.favorite_dialog = &FavoriteDialog{}
        m.favorite_dialog.InitializeUpdate(&srv[0])
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

// Main window initialization.
func (m *MainWindow) Initialize() {
    m.initializeStorages()

    gtk.Init(nil)
    m.window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
    m.window.SetTitle("URTrator")
    m.window.Connect("destroy", m.Close)
    m.vbox = gtk.NewVBox(false, 0)

    // Load program icon from base64.
    icon_bytes, _ := base64.StdEncoding.DecodeString(common.Logo)
    icon_pixbuf := gdkpixbuf.NewLoader()
    icon_pixbuf.Write(icon_bytes)
    logo = icon_pixbuf.GetPixbuf()
    m.window.SetIcon(logo)

    m.window.Connect("configure-event", m.checkPositionAndSize)

    // Restoring window position.
    var win_pos_x_str string = "0"
    var win_pos_y_str string = "0"
    saved_win_pos_x_str, ok := ctx.Cfg.Cfg["/mainwindow/position_x"]
    if ok {
        win_pos_x_str = saved_win_pos_x_str
    }
    saved_win_pos_y_str, ok := ctx.Cfg.Cfg["/mainwindow/position_y"]
    if ok {
        win_pos_y_str = saved_win_pos_y_str
    }
    win_pos_x, _ := strconv.Atoi(win_pos_x_str)
    win_pos_y, _ := strconv.Atoi(win_pos_y_str)
    m.window.Move(win_pos_x, win_pos_y)

    // Restoring window size.
    var win_size_width_str string = "1000"
    var win_size_height_str string = "600"
    saved_win_size_width_str, ok := ctx.Cfg.Cfg["/mainwindow/width"]
    if ok {
        win_size_width_str = saved_win_size_width_str
    }
    saved_win_size_height_str, ok := ctx.Cfg.Cfg["/mainwindow/height"]
    if ok {
        win_size_height_str = saved_win_size_height_str
    }

    m.window_width, _ = strconv.Atoi(win_size_width_str)
    m.window_height, _ = strconv.Atoi(win_size_height_str)
    m.window.SetDefaultSize(m.window_width, m.window_height)

    // Dialogs initialization.
    m.options_dialog = &OptionsDialog{}

    // Main menu.
    m.InitializeMainMenu()

    // Toolbar.
    m.InitializeToolbar()

    m.hpane = gtk.NewHPaned()
    m.vbox.PackStart(m.hpane, true, true, 5)
    m.hpane.Connect("event", m.checkMainPanePosition)

    // Restore pane position.
    // We will restore saved thing, or will use "window_width - 150".
    saved_pane_pos, ok := ctx.Cfg.Cfg["/mainwindow/pane_negative_position"]
    if ok {
        pane_negative_pos, _ := strconv.Atoi(saved_pane_pos)
        m.hpane.SetPosition(m.window_width - pane_negative_pos)
    } else {
        var w, h int = 0, 0
        m.window.GetSize(&w, &h)
        m.hpane.SetPosition(w - 150)
    }

    // Tabs initialization.
    m.InitializeTabs()

    // Sidebar initialization.
    m.initializeSidebar()

    // Tray icon.
    if ctx.Cfg.Cfg["/general/show_tray_icon"] == "1" {
        m.initializeTrayIcon()
    }

    // Events.
    m.initializeEvents()

    // Game profiles and launch button.
    profile_and_launch_hbox := gtk.NewHBox(false, 0)
    m.vbox.PackStart(profile_and_launch_hbox, false, true, 5)

    // Separator
    sep := gtk.NewHSeparator()
    profile_and_launch_hbox.PackStart(sep, true, true, 5)

    // Profile selection.
    profiles_label := gtk.NewLabel("Game profile:")
    m.profiles = gtk.NewComboBoxText()
    m.profiles.SetTooltipText("Profile which will be used for launching")

    profile_and_launch_hbox.PackStart(profiles_label, false, true, 5)
    profile_and_launch_hbox.PackStart(m.profiles, false, true, 5)

    ctx.Eventer.AddEventHandler("loadProfiles", m.loadProfiles)

    // One more separator.
    sepp := gtk.NewVSeparator()
    profile_and_launch_hbox.PackStart(sepp, false, true, 5)

    // Game launching button.
    m.launch_button = gtk.NewButtonWithLabel("Launch!")
    m.launch_button.SetTooltipText("Launch Urban Terror")
    m.launch_button.Clicked(m.launchGame)
    launch_button_image := gtk.NewImageFromStock(gtk.STOCK_APPLY, 32)
    m.launch_button.SetImage(launch_button_image)
    profile_and_launch_hbox.PackStart(m.launch_button, false, true, 5)

    // Statusbar.
    m.statusbar = gtk.NewStatusbar()
    m.vbox.PackStart(m.statusbar, false, true, 0)

    m.statusbar_context_id = m.statusbar.GetContextId("Status Bar")
    m.statusbar.Push(m.statusbar_context_id, "URTrator is ready")

    m.window.Add(m.vbox)
    m.window.ShowAll()

    // Launch events.
    ctx.Eventer.LaunchEvent("loadProfiles")
    ctx.Eventer.LaunchEvent("loadServersIntoCache")
    ctx.Eventer.LaunchEvent("loadAllServers")
    ctx.Eventer.LaunchEvent("loadFavoriteServers")

    gtk.Main()
}

// Events initialization.
func (m *MainWindow) initializeEvents() {
    fmt.Println("Initializing events...")
    ctx.Eventer.AddEventHandler("loadAllServers", m.loadAllServers)
    ctx.Eventer.AddEventHandler("loadFavoriteServers", m.loadFavoriteServers)
    ctx.Eventer.AddEventHandler("serversUpdateCompleted", m.serversUpdateCompleted)
}

// Main menu initialization.
func (m *MainWindow) InitializeMainMenu() {
    m.menubar = gtk.NewMenuBar()
    m.vbox.PackStart(m.menubar, false, false, 0)

    // File menu.
    fm := gtk.NewMenuItemWithMnemonic("File")
    m.menubar.Append(fm)
    file_menu := gtk.NewMenu()
    fm.SetSubmenu(file_menu)

    // Options.
    options_menu_item := gtk.NewMenuItemWithMnemonic("_Options")
    file_menu.Append(options_menu_item)
    options_menu_item.Connect("activate", m.options_dialog.ShowOptionsDialog)

    // Separator.
    file_menu_sep1 := gtk.NewSeparatorMenuItem()
    file_menu.Append(file_menu_sep1)

    // Exit.
    exit_menu_item := gtk.NewMenuItemWithMnemonic("E_xit")
    file_menu.Append(exit_menu_item)
    exit_menu_item.Connect("activate", m.Close)

    // About menu.
    am := gtk.NewMenuItemWithMnemonic("_About")
    m.menubar.Append(am)
    about_menu := gtk.NewMenu()
    am.SetSubmenu(about_menu)

    // About app item.
    about_app_item := gtk.NewMenuItemWithMnemonic("About _URTrator...")
    about_menu.Append(about_app_item)
    about_app_item.Connect("activate", ShowAboutDialog)

    // Separator.
    about_menu_sep1 := gtk.NewSeparatorMenuItem()
    about_menu.Append(about_menu_sep1)

    // Drop databases thing.
    about_menu_drop_database_data_item := gtk.NewMenuItemWithMnemonic("Drop database data...")
    about_menu.Append(about_menu_drop_database_data_item)
    about_menu_drop_database_data_item.Connect("activate", m.dropDatabasesData)
}

// Sidebar (with quick connect and server's information) initialization.
func (m *MainWindow) initializeSidebar() {
    sidebar_vbox := gtk.NewVBox(false, 0)

    // Quick connect frame.
    quick_connect_frame := gtk.NewFrame("Quick connect")
    sidebar_vbox.PackStart(quick_connect_frame, true, true, 5)
    qc_vbox := gtk.NewVBox(false, 0)
    quick_connect_frame.Add(qc_vbox)

    // Server address.
    srv_tooltip := "Server address we will connect to"
    srv_label := gtk.NewLabel("Server address:")
    srv_label.SetTooltipText(srv_tooltip)
    qc_vbox.PackStart(srv_label, false, true, 5)

    m.qc_server_address = gtk.NewEntry()
    m.qc_server_address.SetTooltipText(srv_tooltip)
    qc_vbox.PackStart(m.qc_server_address, false, true, 5)

    // Password.
    pass_tooltip := "Password we will use for server"
    pass_label := gtk.NewLabel("Password:")
    pass_label.SetTooltipText(pass_tooltip)
    qc_vbox.PackStart(pass_label, false, true, 5)

    m.qc_password = gtk.NewEntry()
    m.qc_password.SetTooltipText(pass_tooltip)
    qc_vbox.PackStart(m.qc_password, false, true, 5)

    m.hpane.Add2(sidebar_vbox)
}

// Initializes internal storages.
func (m *MainWindow) initializeStorages() {
    // Gamemodes.
    m.gamemodes = make(map[string]string)
    m.gamemodes = map[string]string{
        "1": "Last Man Standing",
        "2": "Free For All",
        "3": "Team DM",
        "4": "Team Survivor",
        "5": "Follow The Leader",
        "6": "Cap'n'Hold",
        "7": "Capture The Flag",
        "8": "Bomb",
        "9": "Jump",
        "10": "Freeze Tag",
        "11": "Gun Game",
        "12": "Instagib",
    }

    // Frames storage.
    m.tabs = make(map[string]*gtk.Frame)
    m.tabs["dummy"] = gtk.NewFrame("dummy")
    delete(m.tabs, "dummy")

    // Servers tab list view storage.
    // Structure:
    // Server status icon|Server name|Mode|Map|Players|Ping|Version
    m.all_servers_store = gtk.NewListStore(gdkpixbuf.GetType(), glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING)
    m.all_servers_store_sortable = gtk.NewTreeSortable(m.all_servers_store)

    // Same as above, but for favorite servers.
    m.fav_servers_store = gtk.NewListStore(gdkpixbuf.GetType(), glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING)

    // Profiles count after filling combobox. Defaulting to 0.
    m.old_profiles_count = 0

    // Window hidden flag.
    m.hidden = false
}

// Tabs widget initialization, including all child widgets.
func (m *MainWindow) InitializeTabs() {
    // Create tabs widget.
    m.tab_widget = gtk.NewNotebook()

    tab_allsrv_hbox := gtk.NewHBox(false, 0)
    swin1 := gtk.NewScrolledWindow(nil, nil)

    m.all_servers = gtk.NewTreeView()
    swin1.Add(m.all_servers)
    tab_allsrv_hbox.PackStart(swin1, true, true, 5)
    m.tab_widget.AppendPage(tab_allsrv_hbox, gtk.NewLabel("Servers"))

    m.all_servers.SetModel(m.all_servers_store)
    m.all_servers.AppendColumn(gtk.NewTreeViewColumnWithAttributes("Status", gtk.NewCellRendererPixbuf(), "pixbuf", 0))

    all_srv_name_column := gtk.NewTreeViewColumnWithAttributes("Name", gtk.NewCellRendererText(), "markup", 1)
    all_srv_name_column.SetSortColumnId(1)
    m.all_servers.AppendColumn(all_srv_name_column)

    all_gamemode_column := gtk.NewTreeViewColumnWithAttributes("Mode", gtk.NewCellRendererText(), "text", 2)
    all_gamemode_column.SetSortColumnId(2)
    m.all_servers.AppendColumn(all_gamemode_column)

    all_map_column := gtk.NewTreeViewColumnWithAttributes("Map", gtk.NewCellRendererText(), "text", 3)
    all_map_column.SetSortColumnId(3)
    m.all_servers.AppendColumn(all_map_column)

    // ToDo: custom sorting function.
    all_players_column := gtk.NewTreeViewColumnWithAttributes("Players", gtk.NewCellRendererText(), "text", 4)
    all_players_column.SetSortColumnId(4)
    m.all_servers.AppendColumn(all_players_column)

    all_ping_column := gtk.NewTreeViewColumnWithAttributes("Ping", gtk.NewCellRendererText(), "text", 5)
    all_ping_column.SetSortColumnId(5)
    m.all_servers.AppendColumn(all_ping_column)

    all_version_column := gtk.NewTreeViewColumnWithAttributes("Version", gtk.NewCellRendererText(), "text", 6)
    all_version_column.SetSortColumnId(6)
    m.all_servers.AppendColumn(all_version_column)

    all_ip_column := gtk.NewTreeViewColumnWithAttributes("IP", gtk.NewCellRendererText(), "text", 7)
    all_ip_column.SetSortColumnId(7)
    m.all_servers.AppendColumn(all_ip_column)
    // Sorting.
    // By default we are sorting by server name.
    // ToDo: remembering it to configuration storage.
    m.all_servers_store_sortable.SetSortColumnId(1, gtk.SORT_ASCENDING)

    // VBox for some servers list controllers.
    tab_all_srv_ctl_vbox := gtk.NewVBox(false, 0)
    tab_allsrv_hbox.PackStart(tab_all_srv_ctl_vbox, false, true, 5)

    // Checkbox for hiding offline servers.
    m.all_servers_hide_offline = gtk.NewCheckButtonWithLabel("Hide offline servers")
    m.all_servers_hide_offline.SetTooltipText("Hide offline servers on Servers tab")
    tab_all_srv_ctl_vbox.PackStart(m.all_servers_hide_offline, false, true, 5)
    m.all_servers_hide_offline.Clicked(m.hideOfflineAllServers)
    if ctx.Cfg.Cfg["/serverslist/all_servers/hide_offline"] == "1" {
        m.all_servers_hide_offline.SetActive(true)
    }

    // Final separator.
    ctl_sep := gtk.NewVSeparator()
    tab_all_srv_ctl_vbox.PackStart(ctl_sep, true, true, 5)

    // Favorites servers
    // ToDo: sorting as in all servers list.
    tab_fav_srv_hbox := gtk.NewHBox(false, 0)
    m.fav_servers = gtk.NewTreeView()
    swin2 := gtk.NewScrolledWindow(nil, nil)
    swin2.Add(m.fav_servers)
    tab_fav_srv_hbox.PackStart(swin2, true, true, 5)
    m.tab_widget.AppendPage(tab_fav_srv_hbox, gtk.NewLabel("Favorites"))
    m.fav_servers.SetModel(m.fav_servers_store)
    m.fav_servers.AppendColumn(gtk.NewTreeViewColumnWithAttributes("Status", gtk.NewCellRendererPixbuf(), "pixbuf", 0))

    fav_name_column := gtk.NewTreeViewColumnWithAttributes("Name", gtk.NewCellRendererText(), "markup", 1)
    fav_name_column.SetSortColumnId(1)
    m.fav_servers.AppendColumn(fav_name_column)

    fav_mode_column := gtk.NewTreeViewColumnWithAttributes("Mode", gtk.NewCellRendererText(), "text", 2)
    fav_mode_column.SetSortColumnId(2)
    m.fav_servers.AppendColumn(fav_mode_column)

    fav_map_column := gtk.NewTreeViewColumnWithAttributes("Map", gtk.NewCellRendererText(), "text", 3)
    fav_map_column.SetSortColumnId(3)
    m.fav_servers.AppendColumn(fav_map_column)

    fav_players_column := gtk.NewTreeViewColumnWithAttributes("Players", gtk.NewCellRendererText(), "text", 4)
    fav_players_column.SetSortColumnId(4)
    m.fav_servers.AppendColumn(fav_players_column)

    fav_ping_column := gtk.NewTreeViewColumnWithAttributes("Ping", gtk.NewCellRendererText(), "text", 5)
    fav_ping_column.SetSortColumnId(5)
    m.fav_servers.AppendColumn(fav_ping_column)

    fav_version_column := gtk.NewTreeViewColumnWithAttributes("Version", gtk.NewCellRendererText(), "text", 6)
    fav_version_column.SetSortColumnId(6)
    m.fav_servers.AppendColumn(fav_version_column)

    fav_ip_column := gtk.NewTreeViewColumnWithAttributes("IP", gtk.NewCellRendererText(), "text", 7)
    fav_ip_column.SetSortColumnId(7)
    m.fav_servers.AppendColumn(fav_ip_column)

    // VBox for some servers list controllers.
    tab_fav_srv_ctl_vbox := gtk.NewVBox(false, 0)
    tab_fav_srv_hbox.PackStart(tab_fav_srv_ctl_vbox, false, true, 5)

    // Checkbox for hiding offline servers.
    m.fav_servers_hide_offline = gtk.NewCheckButtonWithLabel("Hide offline servers")
    m.fav_servers_hide_offline.SetTooltipText("Hide offline servers on Favorites tab")
    tab_fav_srv_ctl_vbox.PackStart(m.fav_servers_hide_offline, false, true, 5)
    m.fav_servers_hide_offline.Clicked(m.hideOfflineFavoriteServers)
    if ctx.Cfg.Cfg["/serverslist/favorite/hide_offline"] == "1" {
        m.fav_servers_hide_offline.SetActive(true)
    }

    // Final separator.
    ctl_fav_sep := gtk.NewVSeparator()
    tab_fav_srv_ctl_vbox.PackStart(ctl_fav_sep, true, true, 5)

    // Add tab_widget widget to window.
    m.hpane.Add1(m.tab_widget)
}

// Toolbar initialization.
func (m *MainWindow) InitializeToolbar() {
    m.toolbar = gtk.NewToolbar()
    m.vbox.PackStart(m.toolbar, false, false, 5)

    // Update servers button.
    button_update_all_servers := gtk.NewToolButtonFromStock(gtk.STOCK_REFRESH)
    button_update_all_servers.SetLabel("Update all servers")
    button_update_all_servers.SetTooltipText("Update all servers in all tabs")
    button_update_all_servers.OnClicked(m.UpdateServers)
    m.toolbar.Insert(button_update_all_servers, 0)

    // Separator.
    separator := gtk.NewSeparatorToolItem()
    m.toolbar.Insert(separator, 1)

    // Add server to favorites button.
    fav_button := gtk.NewToolButtonFromStock(gtk.STOCK_ADD)
    fav_button.SetLabel("Add to favorites")
    fav_button.SetTooltipText("Add selected server to favorites")
    fav_button.OnClicked(m.addToFavorites)
    m.toolbar.Insert(fav_button, 2)

    fav_edit_button := gtk.NewToolButtonFromStock(gtk.STOCK_EDIT)
    fav_edit_button.SetLabel("Edit favorite")
    fav_edit_button.SetTooltipText("Edit selected favorite server")
    fav_edit_button.OnClicked(m.editFavorite)
    m.toolbar.Insert(fav_edit_button, 3)

    // Remove server from favorites button.
    fav_delete_button := gtk.NewToolButtonFromStock(gtk.STOCK_REMOVE)
    fav_delete_button.SetLabel("Remove from favorites")
    fav_delete_button.SetTooltipText("Remove selected server from favorites")
    fav_delete_button.OnClicked(m.deleteFromFavorites)
    m.toolbar.Insert(fav_delete_button, 4)

    // Separator for toolbar's label and buttons.
    toolbar_separator_toolitem := gtk.NewToolItem()
    toolbar_separator_toolitem.SetExpand(true)
    m.toolbar.Insert(toolbar_separator_toolitem, 5)
    // Toolbar's label.
    m.toolbar_label = gtk.NewLabel("URTrator is ready")
    toolbar_label_toolitem := gtk.NewToolItem()
    toolbar_label_toolitem.Add(m.toolbar_label)
    m.toolbar.Insert(toolbar_label_toolitem, 6)
}

// Tray icon initialization.
func (m *MainWindow) initializeTrayIcon() {
    fmt.Println("Initializing tray icon...")

    icon_bytes, _ := base64.StdEncoding.DecodeString(common.Logo)
    icon_pixbuf := gdkpixbuf.NewLoader()
    icon_pixbuf.Write(icon_bytes)
    logo = icon_pixbuf.GetPixbuf()

    m.tray_icon = gtk.NewStatusIconFromPixbuf(logo)
    m.tray_icon.SetName("URTrator")
    m.tray_icon.SetTitle("URTrator")
    m.tray_icon.SetTooltipText("URTrator is ready")

    // Tray menu is still buggy on windows, so skipping initialization,
    // if OS is Windows.
    if runtime.GOOS != "windows" {
        m.tray_menu = gtk.NewMenu()

        // Open/Close URTrator menu item.
        open_close_item := gtk.NewMenuItemWithLabel("Show / Hide URTrator")
        open_close_item.Connect("activate", m.showHide)
        m.tray_menu.Append(open_close_item)

        // Separator
        sep1 := gtk.NewSeparatorMenuItem()
        m.tray_menu.Append(sep1)

        // Exit menu item.
        exit_item := gtk.NewMenuItemWithLabel("Exit")
        exit_item.Connect("activate", m.window.Destroy)
        m.tray_menu.Append(exit_item)

        // Connect things.
        m.tray_icon.Connect("activate", m.showHide)
        m.tray_icon.Connect("popup-menu", m.showTrayMenu)
        m.tray_menu.ShowAll()
    }
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
    model.GetValue(iter, 1, srv_name_gval)
    server_name := srv_name_gval.GetString()

    // Getting server address.
    var srv_addr string
    srv_address_gval := glib.ValueFromNative(srv_addr)
    model.GetValue(iter, 7, srv_address_gval)
    srv_address := srv_address_gval.GetString()

    // Getting server's game version.
    var srv_game_ver_raw string
    srv_game_ver_gval := glib.ValueFromNative(srv_game_ver_raw)
    model.GetValue(iter, 6, srv_game_ver_gval)
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
    m.statusbar.Push(m.statusbar_context_id, "Launching Urban Terror...")
    m.toolbar_label.SetMarkup("<markup><span foreground=\"red\" font_weight=\"bold\">Urban Terror is launched with profile </span><span foreground=\"blue\">" + profile[0].Name + "</span> <span foreground=\"red\" font_weight=\"bold\">and connected to </span><span foreground=\"orange\" font_weight=\"bold\">" + srv_address + "</span></markup>")
    m.launch_button.SetSensitive(false)
    // ToDo: handling server passwords.
    ctx.Launcher.Launch(&profile[0], srv_address, "", m.unlockInterface)

    return nil
}

func (m *MainWindow) loadAllServers() {
    fmt.Println("Loading all servers...")
    // ToDo: do it without clearing.
    for _, server := range ctx.Cache.Servers {
        if m.all_servers_hide_offline.GetActive() && server.Server.Name == "" && server.Server.Players == "" {
            continue
        }
        iter := new (gtk.TreeIter)
        if !server.AllServersIterSet {
            server.AllServersIter = iter
            m.all_servers_store.Append(iter)
            server.AllServersIterSet = true
        } else {
            iter = server.AllServersIter
        }
        if server.Server.Name == "" && server.Server.Players == "" {
            m.all_servers_store.Set(iter, 0, gtk.NewImage().RenderIcon(gtk.STOCK_NO, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
        } else {
            m.all_servers_store.Set(iter, 0, gtk.NewImage().RenderIcon(gtk.STOCK_OK, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
        }
        server_name := ctx.Colorizer.Fix(server.Server.Name)
        m.all_servers_store.SetValue(iter, 1, server_name)
        m.all_servers_store.SetValue(iter, 2, m.gamemodes[server.Server.Gamemode])
        m.all_servers_store.SetValue(iter, 3, server.Server.Map)
        m.all_servers_store.SetValue(iter, 4, server.Server.Players + "/" + server.Server.Maxplayers)
        m.all_servers_store.SetValue(iter, 5, server.Server.Ping)
        m.all_servers_store.SetValue(iter, 6, server.Server.Version)
        m.all_servers_store.SetValue(iter, 7, server.Server.Ip + ":" + server.Server.Port)
    }
}

func (m *MainWindow) loadFavoriteServers() {
    fmt.Println("Loading favorite servers...")
    for _, server := range ctx.Cache.Servers {
        if server.Server.Favorite != "1" {
            continue
        }
        if m.fav_servers_hide_offline.GetActive() && server.Server.Name == "" && server.Server.Players == "" {
            continue
        }
        iter := new(gtk.TreeIter)
        if !server.FavServersIterSet {
            server.FavServersIter = iter
            m.fav_servers_store.Append(iter)
            server.FavServersIterSet = true
        } else {
            iter = server.FavServersIter
        }
        if m.fav_servers_store.IterIsValid(iter) {
            if server.Server.Name == "" && server.Server.Players == "" {
                m.fav_servers_store.Set(iter, 0, gtk.NewImage().RenderIcon(gtk.STOCK_NO, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
            } else {
                m.fav_servers_store.Set(iter, 0, gtk.NewImage().RenderIcon(gtk.STOCK_OK, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
            }
            server_name := ctx.Colorizer.Fix(server.Server.Name)
            m.fav_servers_store.SetValue(iter, 1, server_name)
            m.fav_servers_store.SetValue(iter, 2, m.gamemodes[server.Server.Gamemode])
            m.fav_servers_store.SetValue(iter, 3, server.Server.Map)
            m.fav_servers_store.SetValue(iter, 4, server.Server.Players + "/" + server.Server.Maxplayers)
            m.fav_servers_store.SetValue(iter, 5, server.Server.Ping)
            m.fav_servers_store.SetValue(iter, 6, server.Server.Version)
            m.fav_servers_store.SetValue(iter, 7, server.Server.Ip + ":" + server.Server.Port)
        } else {
            fmt.Println("Invalid iter for server: " + server.Server.Ip + ":" + server.Server.Port)
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
