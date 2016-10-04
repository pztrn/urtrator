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

func (m *MainWindow) Close() {
    ctx.Close()
}

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

func (m *MainWindow) hideOfflineAllServers() {
    fmt.Println("(Un)Hiding offline servers in 'Servers' tab...")
    ctx.Eventer.LaunchEvent("loadAllServers")
}

func (m *MainWindow) hideOfflineFavoriteServers() {
    fmt.Println("(Un)Hiding offline servers in 'Favorite' tab...")
    ctx.Eventer.LaunchEvent("loadFavoriteServers")
}

func (m *MainWindow) Initialize() {
    m.initializeStorages()

    gtk.Init(nil)
    m.window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
    m.window.SetTitle("URTrator")
    m.window.Connect("destroy", m.Close)
    m.vbox = gtk.NewVBox(false, 0)
    m.hpane = gtk.NewHPaned()

    // Load program icon from base64.
    icon_bytes, _ := base64.StdEncoding.DecodeString(common.Logo)
    icon_pixbuf := gdkpixbuf.NewLoader()
    icon_pixbuf.Write(icon_bytes)
    logo = icon_pixbuf.GetPixbuf()
    m.window.SetIcon(logo)

    // Default window size.
    // ToDo: size and position restoring.
    m.window.SetSizeRequest(1000, 600)
    m.window.SetPosition(gtk.WIN_POS_CENTER)

    // Dialogs initialization.
    m.options_dialog = &OptionsDialog{}

    // Main menu.
    m.InitializeMainMenu()

    // Toolbar.
    m.InitializeToolbar()

    // Tabs initialization.
    m.InitializeTabs()

    // Sidebar initialization.
    m.initializeSidebar()

    m.vbox.PackStart(m.hpane, true, true, 5)

    // Temporary hack.
    var w, h int = 0, 0
    m.window.GetSize(&w, &h)
    m.hpane.SetPosition(w - 150)

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
    ctx.Eventer.LaunchEvent("loadAllServers")
    ctx.Eventer.LaunchEvent("loadFavoriteServers")

    gtk.Main()
}

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
}

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

    all_srv_name_column := gtk.NewTreeViewColumnWithAttributes("Name", gtk.NewCellRendererText(), "text", 1)
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

    fav_name_column := gtk.NewTreeViewColumnWithAttributes("Name", gtk.NewCellRendererText(), "text", 1)
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

    // Final separator.
    ctl_fav_sep := gtk.NewVSeparator()
    tab_fav_srv_ctl_vbox.PackStart(ctl_fav_sep, true, true, 5)

    // Add tab_widget widget to window.
    m.hpane.Add1(m.tab_widget)

    ctx.Eventer.AddEventHandler("loadAllServers", m.loadAllServers)
    ctx.Eventer.AddEventHandler("loadFavoriteServers", m.loadFavoriteServers)
}

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

    // Remove server from favorites button.
    fav_delete_button := gtk.NewToolButtonFromStock(gtk.STOCK_REMOVE)
    fav_delete_button.SetLabel("Remove from favorites")
    fav_delete_button.SetTooltipText("Remove selected server from favorites")
    fav_delete_button.OnClicked(m.deleteFromFavorites)
    m.toolbar.Insert(fav_delete_button, 3)
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
    m.launch_button.SetSensitive(false)
    // ToDo: handling server passwords.
    ctx.Launcher.Launch(&profile[0], srv_address, "", m.unlockInterface)

    return nil
}

func (m *MainWindow) loadAllServers() {
    fmt.Println("Loading servers lists into widgets...")
    servers := []datamodels.Server{}
    err := ctx.Database.Db.Select(&servers, "SELECT * FROM servers")
    if err != nil {
        fmt.Println(err.Error())
    }
    // ToDo: do it without clearing.
    m.all_servers_store.Clear()
    for _, srv := range servers {
        if m.all_servers_hide_offline.GetActive() && srv.Name == "" && srv.Players == "" {
            continue
        }
        var iter gtk.TreeIter
        m.all_servers_store.Append(&iter)
        if srv.Name == "" && srv.Players == "" {
            m.all_servers_store.Set(&iter, 0, gtk.NewImage().RenderIcon(gtk.STOCK_NO, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
        } else {
            m.all_servers_store.Set(&iter, 0, gtk.NewImage().RenderIcon(gtk.STOCK_OK, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
        }
        m.all_servers_store.Set(&iter, 1, srv.Name)
        m.all_servers_store.Set(&iter, 2, m.gamemodes[srv.Gamemode])
        m.all_servers_store.Set(&iter, 3, srv.Map)
        m.all_servers_store.Set(&iter, 4, srv.Players + "/" + srv.Maxplayers)
        m.all_servers_store.Set(&iter, 5, srv.Ping)
        m.all_servers_store.Set(&iter, 6, srv.Version)
        m.all_servers_store.Set(&iter, 7, srv.Ip + ":" + srv.Port)
    }
}

func (m *MainWindow) loadFavoriteServers() {
    fmt.Println("Loading favorite servers...")
    servers := []datamodels.Server{}
    err := ctx.Database.Db.Select(&servers, "SELECT * FROM servers WHERE favorite='1'")
    if err != nil {
        fmt.Println(err.Error())
    }
    // ToDo: do it without clearing.
    m.fav_servers_store.Clear()
    for _, srv := range servers {
        if srv.Favorite != "1" {
            continue
        }
        if m.fav_servers_hide_offline.GetActive() && srv.Name == "" && srv.Players == "" {
            continue
        }
        var iter gtk.TreeIter
        m.fav_servers_store.Append(&iter)
        if srv.Name == "" && srv.Players == "" {
            m.fav_servers_store.Set(&iter, 0, gtk.NewImage().RenderIcon(gtk.STOCK_NO, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
        } else {
            m.fav_servers_store.Set(&iter, 0, gtk.NewImage().RenderIcon(gtk.STOCK_OK, gtk.ICON_SIZE_SMALL_TOOLBAR, "").GPixbuf)
        }
        m.fav_servers_store.Set(&iter, 1, srv.Name)
        m.fav_servers_store.Set(&iter, 2, m.gamemodes[srv.Gamemode])
        m.fav_servers_store.Set(&iter, 3, srv.Map)
        m.fav_servers_store.Set(&iter, 4, srv.Players + "/" + srv.Maxplayers)
        m.fav_servers_store.Set(&iter, 5, srv.Ping)
        m.fav_servers_store.Set(&iter, 6, srv.Version)
        m.fav_servers_store.Set(&iter, 7, srv.Ip + ":" + srv.Port)
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

func (m *MainWindow) unlockInterface() {
    m.launch_button.SetSensitive(true)
    m.statusbar.Push(m.statusbar_context_id, "URTrator is ready.")
}

func (m *MainWindow) updateFavorites(done_chan chan map[string]*datamodels.Server, error_chan chan bool) {
    m.fav_servers_store.Clear()
    servers := []datamodels.Server{}
    err := ctx.Database.Db.Select(&servers, "SELECT * FROM servers WHERE favorite='1'")
    if err != nil {
        fmt.Println(err.Error())
    }

    var servers_from_db [][]string

    for s := range servers {
        servers_from_db = append(servers_from_db, []string{servers[s].Ip, servers[s].Port})
    }

    go ctx.Requester.UpdateFavoriteServers(servers_from_db, done_chan, error_chan)
}

func (m *MainWindow) UpdateServers() {
    m.statusbar.Push(m.statusbar_context_id, "Updating servers...")
    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    fmt.Println("Updating servers on tab '" + current_tab + "'...")
    done_chan := make(chan map[string]*datamodels.Server, 1)
    error_chan := make(chan bool, 1)
    if strings.Contains(current_tab, "Servers") {
        go ctx.Requester.UpdateAllServers(done_chan, error_chan)
    } else if strings.Contains(current_tab, "Favorites") {
        m.updateFavorites(done_chan, error_chan)
    }

    select {
    case data := <- done_chan:
        fmt.Println("Information about servers successfully gathered")
        ctx.Database.UpdateServers(data)
        if strings.Contains(current_tab, "Servers") {
            ctx.Eventer.LaunchEvent("loadAllServers")
        } else if strings.Contains(current_tab, "Favorites") {
            ctx.Eventer.LaunchEvent("loadFavoriteServers")
        }
    case <- error_chan:
        fmt.Println("Error occured")
    }

    m.statusbar.Push(m.statusbar_context_id, "Servers updated.")
}
