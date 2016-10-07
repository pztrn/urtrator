package ui

import (
    // stdlib
    "encoding/base64"
    "fmt"
    "runtime"
    "strconv"

    // local
    "github.com/pztrn/urtrator/common"

    // Other
    "github.com/mattn/go-gtk/gdkpixbuf"
    "github.com/mattn/go-gtk/glib"
    "github.com/mattn/go-gtk/gtk"
)

// Main window initialization.
func (m *MainWindow) Initialize() {
    gtk.Init(nil)

    m.initializeStorages()

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

    // One more separator.
    sepp := gtk.NewVSeparator()
    profile_and_launch_hbox.PackStart(sepp, false, true, 5)

    // Game launching button.
    m.launch_button = gtk.NewButtonWithLabel("Launch!")
    m.launch_button.SetTooltipText("Launch Urban Terror")
    m.launch_button.Clicked(m.launchGame)
    launch_button_image := gtk.NewImageFromStock(gtk.STOCK_APPLY, 24)
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
    ctx.Eventer.AddEventHandler("loadProfiles", m.loadProfiles)
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

    server_info_frame := gtk.NewFrame("Server information")
    sidebar_vbox.PackStart(server_info_frame, true, true, 5)
    si_vbox := gtk.NewVBox(false, 0)
    server_info_frame.Add(si_vbox)

    // Scrolled thing.
    si_scroll := gtk.NewScrolledWindow(nil, nil)
    si_vbox.PackStart(si_scroll, true, true, 5)

    // Server's information.
    m.server_info = gtk.NewTreeView()
    m.server_info.SetModel(m.server_info_store)

    key_column := gtk.NewTreeViewColumnWithAttributes("Key", gtk.NewCellRendererText(), "markup", 0)
    m.server_info.AppendColumn(key_column)

    value_column := gtk.NewTreeViewColumnWithAttributes("Value", gtk.NewCellRendererText(), "markup", 1)
    m.server_info.AppendColumn(value_column)

    si_scroll.Add(m.server_info)

    // Button to view additional server info.
    additional_srv_info_button := gtk.NewButtonWithLabel("Additional information")
    additional_srv_info_button.Clicked(m.showServerInformation)
    si_vbox.PackStart(additional_srv_info_button, false, true, 5)

    // Quick connect frame.
    quick_connect_frame := gtk.NewFrame("Quick connect")
    sidebar_vbox.PackStart(quick_connect_frame, false, true, 5)
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

    // Columns names.
    // Key - default position in lists.
    m.column_names = map[string]string{
        "2": "Name",
        "3": "Mode",
        "4": "Map",
        "5": "Players",
        "6": "Ping",
        "7": "Version",
        "8": "IP",
    }
    // Real columns positions.
    m.column_pos = make(map[string]map[string]int)
    m.column_pos["Servers"] = make(map[string]int)
    m.column_pos["Favorites"] = make(map[string]int)

    // Frames storage.
    m.tabs = make(map[string]*gtk.Frame)
    m.tabs["dummy"] = gtk.NewFrame("dummy")
    delete(m.tabs, "dummy")

    // Servers tab list view storage.
    // Structure:
    // Server status icon|Server name|Mode|Map|Players|Ping|Version
    m.all_servers_store = gtk.NewListStore(gdkpixbuf.GetType(), gdkpixbuf.GetType(), glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING)
    m.all_servers_store_sortable = gtk.NewTreeSortable(m.all_servers_store)

    // Same as above, but for favorite servers.
    m.fav_servers_store = gtk.NewListStore(gdkpixbuf.GetType(), gdkpixbuf.GetType(), glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING)

    // Server's information store. Used for quick preview in main window.
    m.server_info_store = gtk.NewListStore(glib.G_TYPE_STRING, glib.G_TYPE_STRING)

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
    // These columns are static.
    m.all_servers.AppendColumn(gtk.NewTreeViewColumnWithAttributes("Status", gtk.NewCellRendererPixbuf(), "pixbuf", 0))
    m.all_servers.AppendColumn(gtk.NewTreeViewColumnWithAttributes("Public", gtk.NewCellRendererPixbuf(), "pixbuf", 1))

    // ...aand lets do dynamic generation :)
    // +2 because we have 2 static columns.
    all_servers_columns_to_append := make([]*gtk.TreeViewColumn, len(m.column_names) + 2)
    for pos, name := range m.column_names {
        fmt.Println(pos, name)
        // Check if we have column position saved. If so - use it.
        // Otherwise use default position.
        // Should be actual only for first launch.
        position := ctx.Cfg.Cfg["/mainwindow/all_servers/" + name + "_position"]
        if len(position) == 0 {
            position = pos
        }
        position_int, _ := strconv.Atoi(position)
        // Same for width.
        width := ctx.Cfg.Cfg["/mainwindow/all_servers/" + name + "_width"]
        if len(width) == 0 {
            width = "-1"
        }
        width_int, _ := strconv.Atoi(width)

        col := gtk.NewTreeViewColumnWithAttributes(name, gtk.NewCellRendererText(), "markup", position_int)
        col.SetSortColumnId(position_int)
        col.SetReorderable(true)
        col.SetResizable(true)
        col.SetSizing(gtk.TREE_VIEW_COLUMN_FIXED)
        col.SetFixedWidth(width_int)
        m.column_pos["Servers"][name] = position_int
        all_servers_columns_to_append[position_int] = col
    }

    for i := range all_servers_columns_to_append {
        if i < 2 {
            continue
        }
        m.all_servers.AppendColumn(all_servers_columns_to_append[i])
    }

    // Sorting.
    // By default we are sorting by server name.
    // ToDo: remembering it to configuration storage.
    m.all_servers_store_sortable.SetSortColumnId(m.column_pos["Servers"]["Name"], gtk.SORT_ASCENDING)

    // Selection changed signal, which will update server's short info pane.
    m.all_servers.Connect("cursor-changed", m.showShortServerInformation)

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
    m.fav_servers.AppendColumn(gtk.NewTreeViewColumnWithAttributes("Public", gtk.NewCellRendererPixbuf(), "pixbuf", 1))

    // +2 because we have 2 static columns.
    fav_servers_columns_to_append := make([]*gtk.TreeViewColumn, len(m.column_names) + 2)
    for pos, name := range m.column_names {
        fmt.Println(pos, name)
        // Check if we have column position saved. If so - use it.
        // Otherwise use default position.
        // Should be actual only for first launch.
        position := ctx.Cfg.Cfg["/mainwindow/fav_servers/" + name + "_position"]
        if len(position) == 0 {
            position = pos
        }
        position_int, _ := strconv.Atoi(position)
        // Same for width.
        width := ctx.Cfg.Cfg["/mainwindow/fav_servers/" + name + "_width"]
        if len(width) == 0 {
            width = "-1"
        }
        width_int, _ := strconv.Atoi(width)

        col := gtk.NewTreeViewColumnWithAttributes(name, gtk.NewCellRendererText(), "markup", position_int)
        col.SetSortColumnId(position_int)
        col.SetReorderable(true)
        col.SetResizable(true)
        col.SetSizing(gtk.TREE_VIEW_COLUMN_FIXED)
        col.SetFixedWidth(width_int)
        m.column_pos["Favorites"][name] = position_int
        fav_servers_columns_to_append[position_int] = col
    }

    for i := range fav_servers_columns_to_append {
        if i < 2 {
            continue
        }
        m.fav_servers.AppendColumn(fav_servers_columns_to_append[i])
    }

    // Selection changed signal, which will update server's short info pane.
    m.fav_servers.Connect("cursor-changed", m.showShortServerInformation)

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

    button_update_one_server := gtk.NewToolButtonFromStock(gtk.STOCK_UNDO)
    button_update_one_server.SetLabel("Update all servers")
    button_update_one_server.SetTooltipText("Update only selected server")
    button_update_one_server.OnClicked(m.updateOneServer)
    m.toolbar.Insert(button_update_one_server, 1)

    // Separator.
    separator := gtk.NewSeparatorToolItem()
    m.toolbar.Insert(separator, 2)

    // Add server to favorites button.
    fav_button := gtk.NewToolButtonFromStock(gtk.STOCK_ADD)
    fav_button.SetLabel("Add to favorites")
    fav_button.SetTooltipText("Add selected server to favorites")
    fav_button.OnClicked(m.addToFavorites)
    m.toolbar.Insert(fav_button, 3)

    fav_edit_button := gtk.NewToolButtonFromStock(gtk.STOCK_EDIT)
    fav_edit_button.SetLabel("Edit favorite")
    fav_edit_button.SetTooltipText("Edit selected favorite server")
    fav_edit_button.OnClicked(m.editFavorite)
    m.toolbar.Insert(fav_edit_button, 4)

    // Remove server from favorites button.
    fav_delete_button := gtk.NewToolButtonFromStock(gtk.STOCK_REMOVE)
    fav_delete_button.SetLabel("Remove from favorites")
    fav_delete_button.SetTooltipText("Remove selected server from favorites")
    fav_delete_button.OnClicked(m.deleteFromFavorites)
    m.toolbar.Insert(fav_delete_button, 5)

    // Separator for toolbar's label and buttons.
    toolbar_separator_toolitem := gtk.NewToolItem()
    toolbar_separator_toolitem.SetExpand(true)
    m.toolbar.Insert(toolbar_separator_toolitem, 6)
    // Toolbar's label.
    m.toolbar_label = gtk.NewLabel("URTrator is ready")
    toolbar_label_toolitem := gtk.NewToolItem()
    toolbar_label_toolitem.Add(m.toolbar_label)
    m.toolbar.Insert(toolbar_label_toolitem, 7)
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