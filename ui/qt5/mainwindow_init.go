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
	"os"
	//"runtime"
	"sort"
	"strconv"

	// local
	"github.com/pztrn/urtrator/common"

	// Qt5
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func (m *MainWindow) Initialize() {
	fmt.Println("Initializing main window...")

	m.app = widgets.NewQApplication(len(os.Args), os.Args)

	m.initializeStorages()

	m.window = widgets.NewQMainWindow(nil, 0)
	m.window.SetWindowTitle("URTrator")

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
	m.window.SetGeometry2(win_pos_x, win_pos_y, m.window_width, m.window_height)

	m.initializeMenu()

	// Central widget.
	cv := widgets.NewQWidget(nil, core.Qt__Widget)
	//cv_policy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__DefaultType)
	//cv.SetSizePolicy(cv_policy)
	m.window.SetCentralWidget(cv)

	// Main vertical box.
	m.vbox = widgets.NewQVBoxLayout()
	m.vbox.SetContentsMargins(4, 4, 4, 4)
	cv.SetLayout(m.vbox)

	m.initializeToolbar()
	m.initializeTabs()
	m.initializeSidebar()

	m.window.Show()

	// Restore splitter position.
	// We will restore saved thing, or will use "window_width - 150".
	saved_pane_pos, ok := ctx.Cfg.Cfg["/mainwindow/pane_negative_position"]
	if ok {
		pane_negative_pos, _ := strconv.Atoi(saved_pane_pos)
		new_splitter_pos := m.window_width - pane_negative_pos
		fmt.Println(new_splitter_pos)
		m.splitter.MoveSplitter(new_splitter_pos, 1)
		fmt.Println(m.splitter.ClosestLegalPosition(1, new_splitter_pos))
	} else {
		g := m.window.Geometry()
		w := g.Width()
		m.splitter.MoveSplitter(w-150, 1)
	}

	m.splitter.ConnectSplitterMoved(m.splitterMoved)

	widgets.QApplication_Exec()
}

func (m *MainWindow) initializeSidebar() {
	sidebar_widget := widgets.NewQWidget(nil, core.Qt__Widget)
	m.splitter.AddWidget(sidebar_widget)

	sidebar_layout := widgets.NewQVBoxLayout()
	sidebar_layout.SetContentsMargins(4, 4, 4, 4)
	sidebar_widget.SetLayout(sidebar_layout)

	// Server's information list.
	m.sidebar_server_info = widgets.NewQTreeView(nil)
	sidebar_layout.AddWidget(m.sidebar_server_info, 0, core.Qt__AlignHCenter&core.Qt__AlignTop)

	// Server's players widget.
	m.sidebar_server_players = widgets.NewQTreeView(nil)
	sidebar_layout.AddWidget(m.sidebar_server_players, 0, core.Qt__AlignHCenter&core.Qt__AlignTop)

	// Add spacer.
	spacer := widgets.NewQSpacerItem(6, 6, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	sidebar_layout.AddSpacerItem(spacer)
}

func (m *MainWindow) initializeStorages() {
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
}

func (m *MainWindow) initializeTabs() {
	m.splitter = widgets.NewQSplitter(nil)
	m.splitter.SetOrientation(core.Qt__Horizontal)
	m.vbox.AddWidget(m.splitter, 0, core.Qt__AlignHCenter&core.Qt__AlignTop)

	m.tabs = widgets.NewQTabWidget(nil)
	m.splitter.AddWidget(m.tabs)

	// Default size policy for filters widget.
	filters_size_policy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__DefaultType)

	//////////////////////////////////////////////////
	// Servers page.
	//////////////////////////////////////////////////
	serverspagewidget := widgets.NewQWidget(nil, core.Qt__Widget)
	serverspagewidget_policy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__DefaultType)
	serverspagewidget.SetSizePolicy(serverspagewidget_policy)
	m.tabs.AddTab(serverspagewidget, ctx.Translator.Translate("Servers", nil))
	serverspagewidget_layout := widgets.NewQHBoxLayout()
	serverspagewidget_layout.SetContentsMargins(4, 4, 4, 4)
	serverspagewidget.SetLayout(serverspagewidget_layout)

	// Servers list.
	m.all_servers = widgets.NewQTreeView(nil)
	serverspagewidget_layout.AddWidget(m.all_servers, 0, core.Qt__AlignLeft&core.Qt__AlignTop)

	// Servers list filters widget.
	serverspagewidget_filters_widget := widgets.NewQWidget(nil, core.Qt__Widget)
	serverspagewidget_filters_widget_policy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__DefaultType)
	serverspagewidget_filters_widget.SetSizePolicy(serverspagewidget_filters_widget_policy)
	serverspagewidget_layout.AddWidget(serverspagewidget_filters_widget, 0, core.Qt__AlignRight&core.Qt__AlignTop)

	// Servers list filters layout.
	serverspagewidget_filters_layout := widgets.NewQVBoxLayout()
	serverspagewidget_filters_widget.SetLayout(serverspagewidget_filters_layout)
	serverspagewidget_filters_layout.SetContentsMargins(4, 4, 4, 4)

	// Filters itself.

	// Hide offline servers checkbox.
	m.all_servers_hide_offline = widgets.NewQCheckBox(nil)
	m.all_servers_hide_offline.SetText(ctx.Translator.Translate("Hide offline servers", nil))
	m.all_servers_hide_offline.SetSizePolicy(filters_size_policy)
	serverspagewidget_filters_layout.AddWidget(m.all_servers_hide_offline, 0, core.Qt__AlignTop)
	// Restore value of hide offline servers checkbox.
	// Set to checked for new installations.
	all_servers_hide_offline_cb_val, ok := ctx.Cfg.Cfg["/serverslist/all_servers/hide_offline"]
	if !ok {
		m.all_servers_hide_offline.SetCheckState(2)
	} else {
		if all_servers_hide_offline_cb_val == "1" {
			m.all_servers_hide_offline.SetCheckState(2)
		}
	}

	// Hide private servers.
	m.all_servers_hide_private = widgets.NewQCheckBox(nil)
	m.all_servers_hide_private.SetText(ctx.Translator.Translate("Hide private servers", nil))
	m.all_servers_hide_private.SetSizePolicy(filters_size_policy)
	serverspagewidget_filters_layout.AddWidget(m.all_servers_hide_private, 0, core.Qt__AlignTop)
	// Restore checkbox value.
	all_servers_hide_private_cb_val, ok := ctx.Cfg.Cfg["/serverslist/all_servers/hide_private"]
	if !ok {
		m.all_servers_hide_private.SetCheckState(2)
	} else {
		if all_servers_hide_private_cb_val == "1" {
			m.all_servers_hide_private.SetCheckState(2)
		}
	}

	// Game version.
	m.all_servers_version = widgets.NewQComboBox(nil)
	m.all_servers_version.SetSizePolicy(filters_size_policy)
	serverspagewidget_filters_layout.AddWidget(m.all_servers_version, 0, core.Qt__AlignTop)
	// Fill game version combobox with supported versions.
	m.all_servers_version.AddItems(common.SUPPORTED_URT_VERSIONS)

	// Game mode.
	m.all_servers_gamemode = widgets.NewQComboBox(nil)
	m.all_servers_gamemode.SetSizePolicy(filters_size_policy)
	serverspagewidget_filters_layout.AddWidget(m.all_servers_gamemode, 0, core.Qt__AlignTop)
	// Fill game mode with supported gamemodes.
	// First - create sorted gamemodes keys slice.
	gm_keys := make([]int, 0, len(m.gamemodes))
	for k, _ := range m.gamemodes {
		ki, _ := strconv.Atoi(k)
		gm_keys = append(gm_keys, ki)
	}
	sort.Ints(gm_keys)
	// Create a strings slice with gamemodes, using sorted keys.
	gmodes := make([]string, 0, len(m.gamemodes))
	// Add "All gamemodes" as first gamemode :)
	gmodes = append(gmodes, ctx.Translator.Translate("All gamemodes", nil))
	for i := range gm_keys {
		ks := strconv.Itoa(gm_keys[i])
		gmodes = append(gmodes, m.gamemodes[ks])
	}
	m.all_servers_gamemode.AddItems(gmodes)

	// After creating filters - add spacer to move them on top of widget.
	all_servers_filters_spacer := widgets.NewQSpacerItem(6, 6, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	serverspagewidget_filters_layout.AddSpacerItem(all_servers_filters_spacer)

	//////////////////////////////////////////////////
	// Favorites page.
	//////////////////////////////////////////////////
	favoritespagewidget := widgets.NewQWidget(nil, core.Qt__Widget)
	favoritespagewidget_policy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__DefaultType)
	favoritespagewidget.SetSizePolicy(favoritespagewidget_policy)
	m.tabs.AddTab(favoritespagewidget, ctx.Translator.Translate("Favorites", nil))
	favoritespagewidget_layout := widgets.NewQHBoxLayout()
	favoritespagewidget_layout.SetContentsMargins(4, 4, 4, 4)
	favoritespagewidget.SetLayout(favoritespagewidget_layout)

	// Favorites list.
	m.fav_servers = widgets.NewQTreeView(nil)
	favoritespagewidget_layout.AddWidget(m.fav_servers, 0, core.Qt__AlignHCenter&core.Qt__AlignTop)

	// Favorites list filters widget.
	favoritespagewidget_filters_widget := widgets.NewQWidget(nil, core.Qt__Widget)
	favoritespagewidget_filters_widget_policy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__DefaultType)
	favoritespagewidget_filters_widget.SetSizePolicy(favoritespagewidget_filters_widget_policy)
	favoritespagewidget_layout.AddWidget(favoritespagewidget_filters_widget, 0, core.Qt__AlignRight&core.Qt__AlignTop)

	// Favorites list filters layout.
	favoritespagewidget_filters_layout := widgets.NewQVBoxLayout()
	favoritespagewidget_filters_widget.SetLayout(favoritespagewidget_filters_layout)
	favoritespagewidget_filters_layout.SetContentsMargins(4, 4, 4, 4)

	// Filters itself.

	// Hide offline servers checkbox.
	m.fav_servers_hide_offline = widgets.NewQCheckBox(nil)
	m.fav_servers_hide_offline.SetText(ctx.Translator.Translate("Hide offline servers", nil))
	m.fav_servers_hide_offline.SetSizePolicy(filters_size_policy)
	favoritespagewidget_filters_layout.AddWidget(m.fav_servers_hide_offline, 0, core.Qt__AlignTop)
	// Restore it's value.
	favorite_servers_hide_offline_cb_val, ok := ctx.Cfg.Cfg["/serverslist/favorite/hide_offline"]
	if !ok {
		m.fav_servers_hide_offline.SetCheckState(2)
	} else {
		if favorite_servers_hide_offline_cb_val == "1" {
			m.fav_servers_hide_offline.SetCheckState(2)
		}
	}

	// Hide private servers.
	m.fav_servers_hide_private = widgets.NewQCheckBox(nil)
	m.fav_servers_hide_private.SetText(ctx.Translator.Translate("Hide private servers", nil))
	m.fav_servers_hide_private.SetSizePolicy(filters_size_policy)
	favoritespagewidget_filters_layout.AddWidget(m.fav_servers_hide_private, 0, core.Qt__AlignTop)
	fav_servers_hide_private_cb_val, ok := ctx.Cfg.Cfg["/serverslist/favorite/hide_private"]
	if !ok {
		m.fav_servers_hide_private.SetCheckState(2)
	} else {
		if fav_servers_hide_private_cb_val == "1" {
			m.fav_servers_hide_private.SetCheckState(2)
		}
	}

	// Game version.
	m.fav_servers_version = widgets.NewQComboBox(nil)
	m.fav_servers_version.SetSizePolicy(filters_size_policy)
	favoritespagewidget_filters_layout.AddWidget(m.fav_servers_version, 0, core.Qt__AlignTop)
	// Fill game version combobox with supported versions.
	m.fav_servers_version.AddItems(common.SUPPORTED_URT_VERSIONS)

	// Game mode.
	m.fav_servers_gamemode = widgets.NewQComboBox(nil)
	m.fav_servers_gamemode.SetSizePolicy(filters_size_policy)
	favoritespagewidget_filters_layout.AddWidget(m.fav_servers_gamemode, 0, core.Qt__AlignTop)
	// Fill game mode with supported gamemodes.
	// As we have previously created this list - reuse it.
	m.fav_servers_gamemode.AddItems(gmodes)

	// After creating filters - add spacer to move them on top of widget.
	fav_servers_filters_spacer := widgets.NewQSpacerItem(6, 6, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	favoritespagewidget_filters_layout.AddSpacerItem(fav_servers_filters_spacer)
}

func (m *MainWindow) initializeToolbar() {
	m.toolbar = widgets.NewQToolBar("Main Toolbar", m.window)
	m.window.AddToolBar(core.Qt__TopToolBarArea, m.toolbar)
}
