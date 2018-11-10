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
	"fmt"
	"sort"

	// Local
	"gitlab.com/pztrn/urtrator/ioq3dataparser"

	// Other
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

type ServerCVarsDialog struct {
	// Window.
	window *gtk.Window
	// Main Vertical Box.
	vbox *gtk.VBox
	// Treeview for CVars.
	treeview *gtk.TreeView
	// Store for treeview.
	treeview_store *gtk.ListStore
}

func (scd *ServerCVarsDialog) Close() {
	scd.window.Destroy()
}

func (scd *ServerCVarsDialog) fill(srv_address string) {
	server_info := ctx.Cache.Servers[srv_address].Server
	parsed_general_data := ioq3dataparser.ParseInfoToMap(server_info.ExtendedConfig)

	// Sort it!
	general_data_keys := make([]string, 0, len(parsed_general_data))
	for k := range parsed_general_data {
		general_data_keys = append(general_data_keys, k)
	}

	sort.Strings(general_data_keys)

	for k := range general_data_keys {
		iter := new(gtk.TreeIter)
		scd.treeview_store.Append(iter)
		scd.treeview_store.SetValue(iter, 0, general_data_keys[k])
		scd.treeview_store.SetValue(iter, 1, parsed_general_data[general_data_keys[k]])
	}
}

func (scd *ServerCVarsDialog) Initialize(w *gtk.Window, srv_address string) {
	fmt.Println("Showing server's CVars...")

	scd.initializeStorages()

	scd.window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	scd.window.SetTitle(ctx.Translator.Translate("URTrator - Server CVars", nil))
	scd.window.Connect("destroy", scd.Close)
	scd.window.SetTransientFor(w)
	scd.window.SetDefaultSize(300, 400)

	scd.vbox = gtk.NewVBox(false, 0)
	scd.window.Add(scd.vbox)

	// CVars scrolls.
	si := gtk.NewScrolledWindow(nil, nil)
	scd.vbox.PackStart(si, true, true, 5)

	// CVars list.
	scd.treeview = gtk.NewTreeView()
	scd.treeview.SetModel(scd.treeview_store)
	si.Add(scd.treeview)

	scd.treeview.AppendColumn(gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Key", nil), gtk.NewCellRendererText(), "markup", 0))
	scd.treeview.AppendColumn(gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Value", nil), gtk.NewCellRendererText(), "markup", 1))

	// Close button.
	hbox := gtk.NewHBox(false, 0)
	scd.vbox.PackStart(hbox, false, true, 5)

	sep := gtk.NewHBox(false, 0)
	hbox.PackStart(sep, true, true, 5)

	close_button := gtk.NewButtonWithLabel(ctx.Translator.Translate("Close", nil))
	close_button.Clicked(scd.Close)
	hbox.PackStart(close_button, false, true, 5)

	scd.fill(srv_address)

	scd.window.ShowAll()
}

func (scd *ServerCVarsDialog) initializeStorages() {
	scd.treeview_store = gtk.NewListStore(glib.G_TYPE_STRING, glib.G_TYPE_STRING)
}
