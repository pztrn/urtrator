// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.

// build darwin
package ui

import (
	// stdlib
	//"fmt"
	//"os"
	//"runtime"

	// Qt5
	"github.com/therecipe/qt/widgets"
)

func (m *MainWindow) initializeMenu() {
	m.mainmenu = widgets.NewQMenuBar(nil)

	//////////////////////////////////////////////////
	// File menu.
	//////////////////////////////////////////////////
	filemenu := widgets.NewQMenu2("&File", nil)

	// Options action.
	file_options := filemenu.AddAction("&Options")
	file_options.SetMenuRole(widgets.QAction__PreferencesRole)
	file_options.ConnectTriggered(m.showOptionsDialog)

	// Separator :)
	filemenu.AddSeparator()

	// Exit URTrator.
	file_exit := filemenu.AddAction("&Exit")
	file_exit.SetMenuRole(widgets.QAction__QuitRole)
	file_exit.ConnectTriggered(m.close)

	m.mainmenu.AddMenu(filemenu)
	//////////////////////////////////////////////////
	// About menu
	//////////////////////////////////////////////////
	aboutmenu := widgets.NewQMenu2("&Help", nil)

	// About URTrator.
	about_about := aboutmenu.AddAction("&About URTrator...")
	about_about.SetMenuRole(widgets.QAction__AboutRole)
	about_about.ConnectTriggered(m.showAboutDialog)

	// About Qt.
	about_about_qt := aboutmenu.AddAction("About &Qt...")
	about_about_qt.SetMenuRole(widgets.QAction__AboutQtRole)
	about_about_qt.ConnectTriggered(m.showAboutQtDialog)

	// Separator :)
	aboutmenu.AddSeparator()

	// Drop database data.
	about_drop_database := aboutmenu.AddAction("&Drop database...")
	//about_drop_database.SetMenuRole(widgets.QAction__ApplicationSpecificRole)
	about_drop_database.ConnectTriggered(m.dropDatabasesData)

	m.mainmenu.AddMenu(aboutmenu)

	m.window.SetMenuBar(m.mainmenu)
}
