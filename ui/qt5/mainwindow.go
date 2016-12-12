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
    //"runtime"
    //"sort"
    //"strconv"
    //"strings"

    // Local
    //"github.com/pztrn/urtrator/datamodels"
    //"github.com/pztrn/urtrator/ioq3dataparser"

    // github
    "github.com/therecipe/qt/widgets"
)

type MainWindow struct {
    //////////////////////////////////////////////////
    // Main widgets and pointers.
    //////////////////////////////////////////////////
    // Application.
    app *widgets.QApplication
    // Main window.
    window *widgets.QMainWindow
    // Main menu.
    mainmenu *widgets.QMenuBar
    // Main vertical box.
    vbox *widgets.QVBoxLayout
    // Toolbar.
    toolbar *widgets.QToolBar
    // Splitter.
    splitter *widgets.QSplitter
    // Tabs widget.
    tabs *widgets.QTabWidget

    //////////////////////////////////////////////////
    // Servers lists and related.
    //////////////////////////////////////////////////
    // "Servers" tab list.
    all_servers *widgets.QTreeView
    // Hide offline servers checkbox.
    all_servers_hide_offline *widgets.QCheckBox
    // Hide private servers?
    all_servers_hide_private *widgets.QCheckBox
    // Server's version.
    all_servers_version *widgets.QComboBox
    // Game mode.
    all_servers_gamemode *widgets.QComboBox
    // Favorites tab list.
    fav_servers *widgets.QTreeView
    // Hide offline servers checkbox.
    fav_servers_hide_offline *widgets.QCheckBox
    // Hide private servers?
    fav_servers_hide_private *widgets.QCheckBox
    // Server's version.
    fav_servers_version *widgets.QComboBox
    // Game mode.
    fav_servers_gamemode *widgets.QComboBox
    // Sidebar's server's information widget.
    sidebar_server_info *widgets.QTreeView
    // Sidebar's server's players widget.
    sidebar_server_players *widgets.QTreeView

    //////////////////////////////////////////////////
    // Datas.
    //////////////////////////////////////////////////
    // Window size.
    window_width int
    window_height int
    // Window position.
    window_pos_x int
    window_pos_y int
    // Supported game modes.
    gamemodes map[string]string
    // Columns names for servers tabs.
    column_names map[string]string
    // Real columns positions on servers tabs.
    column_pos map[string]map[string]int
}

func (m *MainWindow) close(a bool) {
    fmt.Println("Closing URTrator...")
    m.app.Quit()
}

func (m *MainWindow) dropDatabasesData(bool) {
    fmt.Println("About to drop databases data...")
}

func (m *MainWindow) showAboutDialog(a bool) {
    fmt.Println("Showing about dialog...")
}

func (m *MainWindow) showAboutQtDialog(a bool) {
    fmt.Println("Showing about Qt dialog...")
    widgets.QMessageBox_AboutQt(m.window, "About Qt")
}


func (m *MainWindow) showOptionsDialog(a bool) {
    fmt.Println("Showing options dialog...")
}

func (m *MainWindow) splitterMoved(pos int, index int) {
    fmt.Println("Splitter moved!")
    fmt.Println(index, pos)
}
