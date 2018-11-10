package ui

import (
	// stdlib
	"encoding/base64"
	"fmt"
	"runtime"
	"sort"
	"strconv"

	// local
	"gitlab.com/pztrn/urtrator/common"
	"gitlab.com/pztrn/urtrator/timer"

	// Other
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

// Main window initialization.
func (m *MainWindow) Initialize() {

	gtk.Init(nil)

	m.initializeStorages()
	ctx.InitializeClipboardWatcher()

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

	// Additional OS-specific initialization.
	if runtime.GOOS == "windows" {
		m.initializeWin()
	}
	if runtime.GOOS == "darwin" {
		m.initializeMac()
	}

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

	// Set some GTK options for this window.
	gtk_opts_raw := gtk.SettingsGetDefault()
	gtk_opts := gtk_opts_raw.ToGObject()
	gtk_opts.Set("gtk-button-images", true)

	// Dialogs initialization.
	m.options_dialog = &OptionsDialog{}
	m.server_cvars_dialog = &ServerCVarsDialog{}

	// Main menu.
	if runtime.GOOS == "darwin" {
		m.initializeMacMenu()
	} else {
		m.InitializeMainMenu()
	}

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
		w, _ := m.window.GetSize()
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
	sep := gtk.NewHBox(false, 0)
	profile_and_launch_hbox.PackStart(sep, true, true, 5)

	// Profile selection.
	profiles_label := gtk.NewLabel(ctx.Translator.Translate("Game profile:", nil))
	m.profiles = gtk.NewComboBoxText()
	m.profiles.SetTooltipText(ctx.Translator.Translate("Profile which will be used for launching", nil))

	profile_and_launch_hbox.PackStart(profiles_label, false, true, 5)
	profile_and_launch_hbox.PackStart(m.profiles, false, true, 5)

	// One more separator.
	sepp := gtk.NewVSeparator()
	profile_and_launch_hbox.PackStart(sepp, false, true, 5)

	// Game launching button.
	m.launch_button = gtk.NewButtonWithLabel(ctx.Translator.Translate("Launch!", nil))
	m.launch_button.SetTooltipText(ctx.Translator.Translate("Launch Urban Terror", nil))
	m.launch_button.Clicked(m.launchGame)
	launch_button_image := gtk.NewImageFromPixbuf(logo.ScaleSimple(24, 24, gdkpixbuf.INTERP_HYPER))
	m.launch_button.SetImage(launch_button_image)
	profile_and_launch_hbox.PackStart(m.launch_button, false, true, 5)

	m.window.Add(m.vbox)

	if runtime.GOOS == "darwin" {
		m.initializeMacAfter()
	}

	m.window.ShowAll()

	// Launch events.
	ctx.Eventer.LaunchEvent("loadProfiles", map[string]string{})
	ctx.Eventer.LaunchEvent("loadProfilesIntoMainWindow", map[string]string{})
	ctx.Eventer.LaunchEvent("loadServersIntoCache", map[string]string{})
	ctx.Eventer.LaunchEvent("loadAllServers", map[string]string{})
	ctx.Eventer.LaunchEvent("loadFavoriteServers", map[string]string{})
	ctx.Eventer.LaunchEvent("initializeTasksForMainWindow", map[string]string{})
	ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": ctx.Translator.Translate("URTrator is ready.", nil)})

	// Set flag that shows to other parts that we're initialized.
	m.initialized = true

	gtk.Main()
}

// Events initialization.
func (m *MainWindow) initializeEvents() {
	fmt.Println("Initializing events...")
	ctx.Eventer.AddEventHandler("initializeTasksForMainWindow", m.initializeTasks)
	ctx.Eventer.AddEventHandler("loadAllServers", m.loadAllServers)
	ctx.Eventer.AddEventHandler("loadFavoriteServers", m.loadFavoriteServers)
	ctx.Eventer.AddEventHandler("loadProfilesIntoMainWindow", m.loadProfiles)
	ctx.Eventer.AddEventHandler("serversUpdateCompleted", m.serversUpdateCompleted)
	ctx.Eventer.AddEventHandler("setQuickConnectDetails", m.setQuickConnectDetails)
	ctx.Eventer.AddEventHandler("setToolbarLabelText", m.setToolbarLabelText)
	ctx.Eventer.AddEventHandler("updateAllServers", m.UpdateServersEventHandler)
}

// Main menu initialization.
func (m *MainWindow) InitializeMainMenu() {
	m.menubar = gtk.NewMenuBar()
	m.vbox.PackStart(m.menubar, false, false, 0)

	// File menu.
	fm := gtk.NewMenuItemWithMnemonic(ctx.Translator.Translate("File", nil))
	m.menubar.Append(fm)
	file_menu := gtk.NewMenu()
	fm.SetSubmenu(file_menu)

	// Options.
	options_menu_item := gtk.NewMenuItemWithMnemonic(ctx.Translator.Translate("_Options", nil))
	file_menu.Append(options_menu_item)
	options_menu_item.Connect("activate", m.options_dialog.ShowOptionsDialog)

	// Separator.
	file_menu_sep1 := gtk.NewSeparatorMenuItem()
	file_menu.Append(file_menu_sep1)

	// Exit.
	exit_menu_item := gtk.NewMenuItemWithMnemonic(ctx.Translator.Translate("E_xit", nil))
	file_menu.Append(exit_menu_item)
	exit_menu_item.Connect("activate", m.Close)

	// About menu.
	am := gtk.NewMenuItemWithMnemonic(ctx.Translator.Translate("_?", nil))
	m.menubar.Append(am)
	about_menu := gtk.NewMenu()
	am.SetSubmenu(about_menu)

	// About app item.
	about_app_item := gtk.NewMenuItemWithMnemonic(ctx.Translator.Translate("About _URTrator...", nil))
	about_menu.Append(about_app_item)
	about_app_item.Connect("activate", ShowAboutDialog)

	// Separator.
	about_menu_sep1 := gtk.NewSeparatorMenuItem()
	about_menu.Append(about_menu_sep1)

	// Drop databases thing.
	about_menu_drop_database_data_item := gtk.NewMenuItemWithMnemonic(ctx.Translator.Translate("Drop local caches and settings", nil))
	about_menu.Append(about_menu_drop_database_data_item)
	about_menu_drop_database_data_item.Connect("activate", m.dropDatabasesData)
}

// Sidebar (with quick connect and server's information) initialization.
func (m *MainWindow) initializeSidebar() {
	sidebar_vbox := gtk.NewVBox(false, 0)

	server_info_frame := gtk.NewFrame(ctx.Translator.Translate("Server information", nil))
	sidebar_vbox.PackStart(server_info_frame, true, true, 5)
	si_vbox := gtk.NewVBox(false, 0)
	server_info_frame.Add(si_vbox)

	// Scrolled thing.
	si_scroll := gtk.NewScrolledWindow(nil, nil)
	si_scroll.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	si_vbox.PackStart(si_scroll, true, true, 5)

	// Server's information.
	m.server_info = gtk.NewTreeView()
	m.server_info.SetModel(m.server_info_store)

	key_column := gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Key", nil), gtk.NewCellRendererText(), "markup", 0)
	m.server_info.AppendColumn(key_column)

	value_column := gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Value", nil), gtk.NewCellRendererText(), "markup", 1)
	m.server_info.AppendColumn(value_column)

	si_scroll.Add(m.server_info)

	// Players information.
	players_info_frame := gtk.NewFrame(ctx.Translator.Translate("Players", nil))
	sidebar_vbox.PackStart(players_info_frame, true, true, 5)

	pi_scroll := gtk.NewScrolledWindow(nil, nil)
	pi_scroll.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	players_info_frame.Add(pi_scroll)

	m.players_info = gtk.NewTreeView()
	m.players_info.SetModel(m.players_info_store)
	pi_scroll.Add(m.players_info)

	name_column := gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Player name", nil), gtk.NewCellRendererText(), "markup", 0)
	m.players_info.AppendColumn(name_column)

	frags_column := gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Frags", nil), gtk.NewCellRendererText(), "markup", 1)
	m.players_info.AppendColumn(frags_column)

	ping_column := gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Ping", nil), gtk.NewCellRendererText(), "markup", 2)
	m.players_info.AppendColumn(ping_column)

	// Show CVars button.
	show_cvars_button := gtk.NewButtonWithLabel(ctx.Translator.Translate("Show CVars", nil))
	show_cvars_button.SetTooltipText(ctx.Translator.Translate("Show server's CVars", nil))
	show_cvars_button.Clicked(m.showServerCVars)
	sidebar_vbox.PackStart(show_cvars_button, false, true, 5)

	// Quick connect frame.
	quick_connect_frame := gtk.NewFrame(ctx.Translator.Translate("Quick connect", nil))
	sidebar_vbox.PackStart(quick_connect_frame, false, true, 5)
	qc_vbox := gtk.NewVBox(false, 0)
	quick_connect_frame.Add(qc_vbox)

	// Server address.
	srv_tooltip := ctx.Translator.Translate("Server address we will connect to", nil)
	srv_label := gtk.NewLabel(ctx.Translator.Translate("Server address:", nil))
	srv_label.SetTooltipText(srv_tooltip)
	qc_vbox.PackStart(srv_label, false, true, 5)

	m.qc_server_address = gtk.NewEntry()
	m.qc_server_address.SetTooltipText(srv_tooltip)
	qc_vbox.PackStart(m.qc_server_address, false, true, 5)

	// Password.
	pass_tooltip := ctx.Translator.Translate("Password we will use for server", nil)
	pass_label := gtk.NewLabel(ctx.Translator.Translate("Password:", nil))
	pass_label.SetTooltipText(pass_tooltip)
	qc_vbox.PackStart(pass_label, false, true, 5)

	m.qc_password = gtk.NewEntry()
	m.qc_password.SetTooltipText(pass_tooltip)
	qc_vbox.PackStart(m.qc_password, false, true, 5)

	// Nickname
	nick_tooltip := ctx.Translator.Translate("Nickname we will use", nil)
	nick_label := gtk.NewLabel(ctx.Translator.Translate("Nickname:", nil))
	nick_label.SetTooltipText(nick_tooltip)
	qc_vbox.PackStart(nick_label, false, true, 5)

	m.qc_nickname = gtk.NewEntry()
	m.qc_nickname.SetTooltipText(nick_tooltip)
	qc_vbox.PackStart(m.qc_nickname, false, true, 5)

	m.hpane.Add2(sidebar_vbox)
}

// Initializes internal storages.
func (m *MainWindow) initializeStorages() {
	// Application isn't initialized.
	m.initialized = false
	m.use_other_servers_tab = false
	m.servers_already_updating = false
	// Gamemodes.
	m.gamemodes = make(map[string]string)
	m.gamemodes = map[string]string{
		"1":  "Last Man Standing",
		"2":  "Free For All",
		"3":  "Team DM",
		"4":  "Team Survivor",
		"5":  "Follow The Leader",
		"6":  "Cap'n'Hold",
		"7":  "Capture The Flag",
		"8":  "Bomb",
		"9":  "Jump",
		"10": "Freeze Tag",
		"11": "Gun Game",
		"12": "Instagib",
	}

	// Columns names.
	// Key - default position in lists.
	m.column_names = map[string]string{
		"2": ctx.Translator.Translate("Name", nil),
		"3": ctx.Translator.Translate("Mode", nil),
		"4": ctx.Translator.Translate("Map", nil),
		"5": ctx.Translator.Translate("Players", nil),
		"6": ctx.Translator.Translate("Ping", nil),
		"7": ctx.Translator.Translate("Version", nil),
		"8": ctx.Translator.Translate("IP", nil),
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

	// Players information store. Used in sidebar for players list for
	// currently selected server.
	m.players_info_store = gtk.NewListStore(glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_STRING)

	// Profiles count after filling combobox. Defaulting to 0.
	m.old_profiles_count = 0

	// Window hidden flag.
	m.hidden = false

	// Pixbufs.
	// Offline server.
	srv_offline_bytes, _ := base64.StdEncoding.DecodeString(common.SERVER_OFFLINE)
	srv_offline_pixbuf, _ := gdkpixbuf.NewLoaderWithType("png")
	srv_offline_pixbuf.SetSize(24, 24)
	srv_offline_pixbuf.Write(srv_offline_bytes)
	m.server_offline_pic = srv_offline_pixbuf.GetPixbuf()
	// Online server.
	srv_online_bytes, _ := base64.StdEncoding.DecodeString(common.SERVER_ONLINE)
	srv_online_pixbuf, _ := gdkpixbuf.NewLoaderWithType("png")
	srv_online_pixbuf.SetSize(24, 24)
	srv_online_pixbuf.Write(srv_online_bytes)
	m.server_online_pic = srv_online_pixbuf.GetPixbuf()
	// Private server.
	srv_private_bytes, _ := base64.StdEncoding.DecodeString(common.SERVER_PRIVATE)
	srv_private_pixbuf, _ := gdkpixbuf.NewLoaderWithType("png")
	srv_private_pixbuf.SetSize(24, 24)
	srv_private_pixbuf.Write(srv_private_bytes)
	m.server_private_pic = srv_private_pixbuf.GetPixbuf()
	// Public server.
	srv_public_bytes, _ := base64.StdEncoding.DecodeString(common.SERVER_PUBLIC)
	srv_public_pixbuf, _ := gdkpixbuf.NewLoaderWithType("png")
	srv_public_pixbuf.SetSize(24, 24)
	srv_public_pixbuf.Write(srv_public_bytes)
	m.server_public_pic = srv_public_pixbuf.GetPixbuf()
}

// Tabs widget initialization, including all child widgets.
func (m *MainWindow) InitializeTabs() {
	// Create tabs widget.
	m.tab_widget = gtk.NewNotebook()
	m.tab_widget.Connect("switch-page", m.tabChanged)

	tab_allsrv_hbox := gtk.NewHBox(false, 0)
	swin1 := gtk.NewScrolledWindow(nil, nil)

	m.all_servers = gtk.NewTreeView()
	swin1.Add(m.all_servers)
	tab_allsrv_hbox.PackStart(swin1, true, true, 5)
	m.tab_widget.AppendPage(tab_allsrv_hbox, gtk.NewLabel(ctx.Translator.Translate("Servers", nil)))

	m.all_servers.SetModel(m.all_servers_store)
	// These columns are static.
	m.all_servers.AppendColumn(gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Status", nil), gtk.NewCellRendererPixbuf(), "pixbuf", 0))
	m.all_servers.AppendColumn(gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Public", nil), gtk.NewCellRendererPixbuf(), "pixbuf", 1))

	// ...aand lets do dynamic generation :)
	// +2 because we have 2 static columns.
	all_servers_columns_to_append := make([]*gtk.TreeViewColumn, len(m.column_names)+2)
	for pos, name := range m.column_names {
		// Check if we have column position saved. If so - use it.
		// Otherwise use default position.
		// Should be actual only for first launch.
		position := ctx.Cfg.Cfg["/mainwindow/all_servers/"+ctx.Translator.Translate(name, nil)+"_position"]
		if len(position) == 0 {
			position = pos
		}
		position_int, _ := strconv.Atoi(position)
		// Same for width.
		width := ctx.Cfg.Cfg["/mainwindow/all_servers/"+ctx.Translator.Translate(name, nil)+"_width"]
		if len(width) == 0 {
			width = "-1"
		}
		width_int, _ := strconv.Atoi(width)

		col := gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate(name, nil), gtk.NewCellRendererText(), "markup", position_int)
		col.SetSortColumnId(position_int)
		col.SetReorderable(true)
		col.SetResizable(true)
		// GtkTreeViewColumn.SetFixedWidth() accepts only positive integers.
		if width_int > 1 {
			col.SetSizing(gtk.TREE_VIEW_COLUMN_FIXED)
			col.SetFixedWidth(width_int)
		}
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
	m.all_servers_store_sortable.SetSortColumnId(m.column_pos["Servers"][ctx.Translator.Translate("Name", nil)], gtk.SORT_ASCENDING)

	// Sorting functions.
	m.all_servers_store_sortable.SetSortFunc(m.column_pos["Servers"][ctx.Translator.Translate("Name", nil)], m.sortServersByName, nil)
	m.all_servers_store_sortable.SetSortFunc(m.column_pos["Servers"][ctx.Translator.Translate("Players", nil)], m.sortServersByPlayers, nil)
	m.all_servers_store_sortable.SetSortFunc(m.column_pos["Servers"][ctx.Translator.Translate("Ping", nil)], m.sortServersByPing, nil)

	// Selection changed signal, which will update server's short info pane.
	m.all_servers.Connect("cursor-changed", m.showShortServerInformation)

	// VBox for some servers list controllers.
	tab_all_srv_ctl_vbox := gtk.NewVBox(false, 0)
	tab_allsrv_hbox.PackStart(tab_all_srv_ctl_vbox, false, true, 5)

	// Checkbox for hiding offline servers.
	m.all_servers_hide_offline = gtk.NewCheckButtonWithLabel(ctx.Translator.Translate("Hide offline servers", nil))
	m.all_servers_hide_offline.SetTooltipText(ctx.Translator.Translate("Hide offline servers on Servers tab", nil))
	tab_all_srv_ctl_vbox.PackStart(m.all_servers_hide_offline, false, true, 5)
	m.all_servers_hide_offline.Clicked(m.hideOfflineAllServers)
	// Restore value of hide offline servers checkbox.
	// Set to checked for new installations.
	all_servers_hide_offline_cb_val, ok := ctx.Cfg.Cfg["/serverslist/all_servers/hide_offline"]
	if !ok {
		m.all_servers_hide_offline.SetActive(true)
	} else {
		if all_servers_hide_offline_cb_val == "1" {
			m.all_servers_hide_offline.SetActive(true)
		}
	}

	// Checkbox for hiding passworded servers.
	m.all_servers_hide_private = gtk.NewCheckButtonWithLabel(ctx.Translator.Translate("Hide private servers", nil))
	m.all_servers_hide_private.SetTooltipText(ctx.Translator.Translate("Hide servers which requires password to enter", nil))
	tab_all_srv_ctl_vbox.PackStart(m.all_servers_hide_private, false, true, 5)
	m.all_servers_hide_private.Clicked(m.hidePrivateAllServers)
	// Restore checkbox value.
	all_servers_hide_private_cb_val, ok := ctx.Cfg.Cfg["/serverslist/all_servers/hide_private"]
	if !ok {
		m.all_servers_hide_private.SetActive(true)
	} else {
		if all_servers_hide_private_cb_val == "1" {
			m.all_servers_hide_private.SetActive(true)
		}
	}

	// Filtering by version.
	m.all_servers_version = gtk.NewComboBoxText()
	m.all_servers_version.SetTooltipText(ctx.Translator.Translate("Show only servers which uses selected version of Urban Terror", nil))
	m.all_servers_version.AppendText(ctx.Translator.Translate("All versions", nil))
	for i := range common.SUPPORTED_URT_VERSIONS {
		m.all_servers_version.AppendText(common.SUPPORTED_URT_VERSIONS[i])
	}
	all_servers_version_val, ok := ctx.Cfg.Cfg["/serverslist/all_servers/version"]
	if ok {
		all_servers_version_int, _ := strconv.Atoi(all_servers_version_val)
		m.all_servers_version.SetActive(all_servers_version_int)
	} else {
		m.all_servers_version.SetActive(0)
	}
	m.all_servers_version.Connect("changed", m.allServersVersionFilterChanged)
	tab_all_srv_ctl_vbox.PackStart(m.all_servers_version, false, true, 5)

	// Filtering by gamemode
	m.all_servers_gamemode = gtk.NewComboBoxText()
	m.all_servers_gamemode.SetTooltipText(ctx.Translator.Translate("Show only servers which uses selected game mode", nil))
	m.all_servers_gamemode.AppendText(ctx.Translator.Translate("All gamemodes", nil))
	// Get sorted gamemodes keys.
	gm_keys := make([]int, 0, len(m.gamemodes))
	for i := range m.gamemodes {
		key, _ := strconv.Atoi(i)
		gm_keys = append(gm_keys, key)
	}
	sort.Ints(gm_keys)
	for i := range gm_keys {
		m.all_servers_gamemode.AppendText(m.gamemodes[strconv.Itoa(gm_keys[i])])
	}
	all_servers_gamemode_val, ok := ctx.Cfg.Cfg["/serverslist/all_servers/gamemode"]
	if ok {
		all_servers_gamemode_int, _ := strconv.Atoi(all_servers_gamemode_val)
		m.all_servers_gamemode.SetActive(all_servers_gamemode_int)
	} else {
		m.all_servers_gamemode.SetActive(0)
	}
	m.all_servers_gamemode.Connect("changed", m.allServersGamemodeFilterChanged)
	tab_all_srv_ctl_vbox.PackStart(m.all_servers_gamemode, false, true, 5)

	// Final separator.
	ctl_sep := gtk.NewVBox(false, 0)
	tab_all_srv_ctl_vbox.PackStart(ctl_sep, true, true, 5)

	// Favorites servers
	// ToDo: sorting as in all servers list.
	tab_fav_srv_hbox := gtk.NewHBox(false, 0)
	m.fav_servers = gtk.NewTreeView()
	swin2 := gtk.NewScrolledWindow(nil, nil)
	swin2.Add(m.fav_servers)
	tab_fav_srv_hbox.PackStart(swin2, true, true, 5)
	m.tab_widget.AppendPage(tab_fav_srv_hbox, gtk.NewLabel(ctx.Translator.Translate("Favorites", nil)))
	m.fav_servers.SetModel(m.fav_servers_store)
	m.fav_servers.AppendColumn(gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Status", nil), gtk.NewCellRendererPixbuf(), "pixbuf", 0))
	m.fav_servers.AppendColumn(gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Public", nil), gtk.NewCellRendererPixbuf(), "pixbuf", 1))

	// +2 because we have 2 static columns.
	fav_servers_columns_to_append := make([]*gtk.TreeViewColumn, len(m.column_names)+2)
	for pos, name := range m.column_names {
		// Check if we have column position saved. If so - use it.
		// Otherwise use default position.
		// Should be actual only for first launch.
		position := ctx.Cfg.Cfg["/mainwindow/fav_servers/"+ctx.Translator.Translate(name, nil)+"_position"]
		if len(position) == 0 {
			position = pos
		}
		position_int, _ := strconv.Atoi(position)
		// Same for width.
		width := ctx.Cfg.Cfg["/mainwindow/fav_servers/"+ctx.Translator.Translate(name, nil)+"_width"]
		if len(width) == 0 {
			width = "-1"
		}
		width_int, _ := strconv.Atoi(width)

		col := gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate(name, nil), gtk.NewCellRendererText(), "markup", position_int)
		// For some reason this cause panic on Windows, so disabling
		// default sorting here.
		if runtime.GOOS != "windows" {
			col.SetSortColumnId(position_int)
		}
		col.SetReorderable(true)
		col.SetResizable(true)
		// GtkTreeViewColumn.SetFixedWidth() accepts only positive integers.
		if width_int > 1 {
			col.SetSizing(gtk.TREE_VIEW_COLUMN_FIXED)
			col.SetFixedWidth(width_int)
		}
		m.column_pos["Favorites"][name] = position_int
		fav_servers_columns_to_append[position_int] = col
	}

	for i := range fav_servers_columns_to_append {
		if i < 2 {
			continue
		}
		m.fav_servers.AppendColumn(fav_servers_columns_to_append[i])
	}

	// Sorting functions.
	m.all_servers_store_sortable.SetSortFunc(m.column_pos["Favorites"][ctx.Translator.Translate("Name", nil)], m.sortServersByName, nil)
	m.all_servers_store_sortable.SetSortFunc(m.column_pos["Favorites"][ctx.Translator.Translate("Players", nil)], m.sortServersByPlayers, nil)
	m.all_servers_store_sortable.SetSortFunc(m.column_pos["Favorites"][ctx.Translator.Translate("Ping", nil)], m.sortServersByPing, nil)

	// Selection changed signal, which will update server's short info pane.
	m.fav_servers.Connect("cursor-changed", m.showShortServerInformation)

	// VBox for some servers list controllers.
	tab_fav_srv_ctl_vbox := gtk.NewVBox(false, 0)
	tab_fav_srv_hbox.PackStart(tab_fav_srv_ctl_vbox, false, true, 5)

	// Checkbox for hiding offline servers.
	m.fav_servers_hide_offline = gtk.NewCheckButtonWithLabel(ctx.Translator.Translate("Hide offline servers", nil))
	m.fav_servers_hide_offline.SetTooltipText(ctx.Translator.Translate("Hide offline servers on Favorites tab", nil))
	tab_fav_srv_ctl_vbox.PackStart(m.fav_servers_hide_offline, false, true, 5)
	m.fav_servers_hide_offline.Clicked(m.hideOfflineFavoriteServers)
	// Restore value of hide offline servers checkbox.
	// Set to checked for new installations.
	favorite_servers_hide_offline_cb_val, ok := ctx.Cfg.Cfg["/serverslist/favorite/hide_offline"]
	if !ok {
		m.fav_servers_hide_offline.SetActive(true)
	} else {
		if favorite_servers_hide_offline_cb_val == "1" {
			m.fav_servers_hide_offline.SetActive(true)
		}
	}

	// Checkbox for hiding passworded servers.
	m.fav_servers_hide_private = gtk.NewCheckButtonWithLabel(ctx.Translator.Translate("Hide private servers", nil))
	m.fav_servers_hide_private.SetTooltipText(ctx.Translator.Translate("Hide servers which requires password to enter", nil))
	tab_fav_srv_ctl_vbox.PackStart(m.fav_servers_hide_private, false, true, 5)
	m.fav_servers_hide_private.Clicked(m.hidePrivateFavoriteServers)
	// Restore checkbox value.
	fav_servers_hide_private_cb_val, ok := ctx.Cfg.Cfg["/serverslist/favorite/hide_private"]
	if !ok {
		m.fav_servers_hide_private.SetActive(true)
	} else {
		if fav_servers_hide_private_cb_val == "1" {
			m.fav_servers_hide_private.SetActive(true)
		}
	}

	m.fav_servers_version = gtk.NewComboBoxText()
	m.fav_servers_version.SetTooltipText(ctx.Translator.Translate("Show only servers which uses selected version of Urban Terror", nil))
	m.fav_servers_version.AppendText(ctx.Translator.Translate("All versions", nil))
	for i := range common.SUPPORTED_URT_VERSIONS {
		m.fav_servers_version.AppendText(common.SUPPORTED_URT_VERSIONS[i])
	}
	fav_servers_version_val, ok := ctx.Cfg.Cfg["/serverslist/favorite/version"]
	if ok {
		fav_servers_version_int, _ := strconv.Atoi(fav_servers_version_val)
		m.fav_servers_version.SetActive(fav_servers_version_int)
	} else {
		m.fav_servers_version.SetActive(0)
	}
	m.fav_servers_version.Connect("changed", m.favServersVersionFilterChanged)
	tab_fav_srv_ctl_vbox.PackStart(m.fav_servers_version, false, true, 5)

	// Filtering by gamemode
	m.fav_servers_gamemode = gtk.NewComboBoxText()
	m.fav_servers_gamemode.SetTooltipText(ctx.Translator.Translate("Show only servers which uses selected game mode", nil))
	m.fav_servers_gamemode.AppendText(ctx.Translator.Translate("All gamemodes", nil))
	// Gamemode keys already sorted while adding same filter to "Servers"
	// tab, so just re-use them.
	for i := range gm_keys {
		m.fav_servers_gamemode.AppendText(m.gamemodes[strconv.Itoa(gm_keys[i])])
	}
	fav_servers_gamemode_val, ok := ctx.Cfg.Cfg["/serverslist/favorite/gamemode"]
	if ok {
		fav_servers_gamemode_int, _ := strconv.Atoi(fav_servers_gamemode_val)
		m.fav_servers_gamemode.SetActive(fav_servers_gamemode_int)
	} else {
		m.fav_servers_gamemode.SetActive(0)
	}
	m.fav_servers_gamemode.Connect("changed", m.favServersGamemodeFilterChanged)
	tab_fav_srv_ctl_vbox.PackStart(m.fav_servers_gamemode, false, true, 5)

	// Final separator.
	ctl_fav_sep := gtk.NewVBox(false, 0)
	tab_fav_srv_ctl_vbox.PackStart(ctl_fav_sep, true, true, 5)

	// Add tab_widget widget to window.
	m.hpane.Add1(m.tab_widget)
}

// Tasks.
func (m *MainWindow) initializeTasks(data map[string]string) {
	// Get task status, if it already running.
	task_status := ctx.Timer.GetTaskStatus("Server's autoupdating")
	// Remove tasks if they exist.
	ctx.Timer.RemoveTask("Server's autoupdating")

	// Add servers autoupdate task.
	if ctx.Cfg.Cfg["/servers_updating/servers_autoupdate"] == "1" {
		task := timer.TimerTask{
			Name:       "Server's autoupdating",
			Callee:     "updateAllServers",
			InProgress: task_status,
		}

		timeout, ok := ctx.Cfg.Cfg["/servers_updating/servers_autoupdate_timeout"]
		if ok {
			timeout_int, err := strconv.Atoi(timeout)
			if err != nil {
				task.Timeout = 10 * 60
			} else {
				task.Timeout = timeout_int * 60
			}
		} else {
			task.Timeout = 10 * 60
		}

		ctx.Timer.AddTask(&task)
	}
}

// Toolbar initialization.
func (m *MainWindow) InitializeToolbar() {
	m.toolbar = gtk.NewToolbar()
	m.vbox.PackStart(m.toolbar, false, false, 5)

	// Update servers button.
	button_update_all_servers_icon_bytes, _ := base64.StdEncoding.DecodeString(common.REFRESH_ALL_SERVERS)
	button_update_all_servers_icon_pixbuf, _ := gdkpixbuf.NewLoaderWithType("png")
	button_update_all_servers_icon_pixbuf.SetSize(24, 24)
	button_update_all_servers_icon_pixbuf.Write(button_update_all_servers_icon_bytes)
	button_update_all_servers_icon := gtk.NewImageFromPixbuf(button_update_all_servers_icon_pixbuf.GetPixbuf())
	button_update_all_servers := gtk.NewToolButton(button_update_all_servers_icon, ctx.Translator.Translate("Update all servers", nil))
	button_update_all_servers.SetTooltipText(ctx.Translator.Translate("Update all servers in currently selected tab", nil))
	button_update_all_servers.OnClicked(m.UpdateServers)
	m.toolbar.Insert(button_update_all_servers, 0)

	button_update_one_server_icon_bytes, _ := base64.StdEncoding.DecodeString(common.REFRESH_ONE_SERVER)
	button_update_one_server_icon_pixbuf, _ := gdkpixbuf.NewLoaderWithType("png")
	button_update_one_server_icon_pixbuf.SetSize(24, 24)
	button_update_one_server_icon_pixbuf.Write(button_update_one_server_icon_bytes)
	button_update_one_server_icon := gtk.NewImageFromPixbuf(button_update_one_server_icon_pixbuf.GetPixbuf())
	button_update_one_server := gtk.NewToolButton(button_update_one_server_icon, ctx.Translator.Translate("Update selected server", nil))
	button_update_one_server.SetTooltipText(ctx.Translator.Translate("Update only selected server", nil))
	button_update_one_server.OnClicked(m.updateOneServer)
	m.toolbar.Insert(button_update_one_server, 1)

	// Separator.
	separator := gtk.NewSeparatorToolItem()
	m.toolbar.Insert(separator, 2)

	// Add server to favorites button.
	fav_button_icon_bytes, _ := base64.StdEncoding.DecodeString(common.ADD_TO_FAVORITES)
	fav_button_icon_pixbuf, _ := gdkpixbuf.NewLoaderWithType("png")
	fav_button_icon_pixbuf.SetSize(24, 24)
	fav_button_icon_pixbuf.Write(fav_button_icon_bytes)
	fav_button_icon := gtk.NewImageFromPixbuf(fav_button_icon_pixbuf.GetPixbuf())
	fav_button := gtk.NewToolButton(fav_button_icon, ctx.Translator.Translate("Add to favorites", nil))
	fav_button.SetTooltipText(ctx.Translator.Translate("Add selected server to favorites", nil))
	fav_button.OnClicked(m.addToFavorites)
	m.toolbar.Insert(fav_button, 3)

	fav_edit_button_icon_bytes, _ := base64.StdEncoding.DecodeString(common.EDIT_FAVORITE)
	fav_edit_button_icon_pixbuf, _ := gdkpixbuf.NewLoaderWithType("png")
	fav_edit_button_icon_pixbuf.SetSize(24, 24)
	fav_edit_button_icon_pixbuf.Write(fav_edit_button_icon_bytes)
	fav_edit_button_icon := gtk.NewImageFromPixbuf(fav_edit_button_icon_pixbuf.GetPixbuf())
	fav_edit_button := gtk.NewToolButton(fav_edit_button_icon, ctx.Translator.Translate("Edit favorite", nil))
	fav_edit_button.SetTooltipText(ctx.Translator.Translate("Edit selected favorite server", nil))
	fav_edit_button.OnClicked(m.editFavorite)
	m.toolbar.Insert(fav_edit_button, 4)

	// Remove server from favorites button.
	fav_delete_button_icon_bytes, _ := base64.StdEncoding.DecodeString(common.REMOVE_FAVORITE)
	fav_delete_button_icon_pixbuf, _ := gdkpixbuf.NewLoaderWithType("png")
	fav_delete_button_icon_pixbuf.SetSize(24, 24)
	fav_delete_button_icon_pixbuf.Write(fav_delete_button_icon_bytes)
	fav_delete_button_icon := gtk.NewImageFromPixbuf(fav_delete_button_icon_pixbuf.GetPixbuf())
	fav_delete_button := gtk.NewToolButton(fav_delete_button_icon, ctx.Translator.Translate("Remove from favorites", nil))
	fav_delete_button.SetTooltipText(ctx.Translator.Translate("Remove selected server from favorites", nil))
	fav_delete_button.OnClicked(m.deleteFromFavorites)
	m.toolbar.Insert(fav_delete_button, 5)

	// Copy server address button.
	copy_srv_addr_button_icon_bytes, _ := base64.StdEncoding.DecodeString(common.COPY_CREDENTIALS)
	copy_srv_addr_button_icon_pixbuf, _ := gdkpixbuf.NewLoaderWithType("png")
	copy_srv_addr_button_icon_pixbuf.SetSize(24, 24)
	copy_srv_addr_button_icon_pixbuf.Write(copy_srv_addr_button_icon_bytes)
	copy_srv_addr_button_icon := gtk.NewImageFromPixbuf(copy_srv_addr_button_icon_pixbuf.GetPixbuf())
	copy_srv_addr_button := gtk.NewToolButton(copy_srv_addr_button_icon, ctx.Translator.Translate("Copy server's creds", nil))
	copy_srv_addr_button.SetTooltipText(ctx.Translator.Translate("Copy server's credentials to clipboard for sharing", nil))
	copy_srv_addr_button.OnClicked(m.copyServerCredentialsToClipboard)
	m.toolbar.Insert(copy_srv_addr_button, 6)

	// Separator for toolbar's label and buttons.
	toolbar_separator_toolitem := gtk.NewToolItem()
	toolbar_separator_toolitem.SetExpand(true)
	m.toolbar.Insert(toolbar_separator_toolitem, 7)
	// Toolbar's label.
	m.toolbar_label = gtk.NewLabel(ctx.Translator.Translate("URTrator is ready", nil))
	toolbar_label_toolitem := gtk.NewToolItem()
	toolbar_label_toolitem.Add(m.toolbar_label)
	m.toolbar.Insert(toolbar_label_toolitem, 8)
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
	m.tray_icon.SetTooltipText(ctx.Translator.Translate("URTrator is ready", nil))

	// Tray menu is still buggy on windows, so skipping initialization,
	// if OS is Windows.
	if runtime.GOOS != "windows" {
		m.tray_menu = gtk.NewMenu()

		// Open/Close URTrator menu item.
		open_close_item := gtk.NewMenuItemWithLabel(ctx.Translator.Translate("Show / Hide URTrator", nil))
		open_close_item.Connect("activate", m.showHide)
		m.tray_menu.Append(open_close_item)

		// Separator
		sep1 := gtk.NewSeparatorMenuItem()
		m.tray_menu.Append(sep1)

		// Exit menu item.
		exit_item := gtk.NewMenuItemWithLabel(ctx.Translator.Translate("Exit", nil))
		exit_item.Connect("activate", m.window.Destroy)
		m.tray_menu.Append(exit_item)

		// Connect things.
		m.tray_icon.Connect("activate", m.showHide)
		m.tray_icon.Connect("popup-menu", m.showTrayMenu)
		m.tray_menu.ShowAll()
	}
}
