// URTrator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016-2018, Stanslav N. a.k.a pztrn (or p0z1tr0n) and
// URTrator contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package ui

import (
	// stdlib
	"os"
	"path/filepath"
	"runtime"

	// other
	"github.com/mattn/go-gtk/gtk"
)

func (m *MainWindow) initializeMac() {
	if runtime.GOOS == "darwin" {
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

		gtk.RCParse(dir + "/../Resources/themes/gtkrc-keybindings")

		// ToDo: theming support and theme seletion in settings.
		gtk.RCParse(dir + "/../Resources/themes/ClearlooksBrave/gtk-2.0/gtkrc")
	}
}

func (m *MainWindow) initializeMacAfter() {
	m.toolbar.SetStyle(gtk.TOOLBAR_ICONS)
}

func (m *MainWindow) initializeMacMenu() {
	// This is a placeholder, in future we will use native mac menu.
	// For now it launches default menu initialization.
	m.InitializeMainMenu()
}
