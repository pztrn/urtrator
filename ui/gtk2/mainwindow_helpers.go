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
	"strings"

	// other
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func (m *MainWindow) getGameModeName(name string) string {
	val, ok := m.gamemodes[name]

	if !ok {
		return "Unknown or custom"
	}

	return val
}

func (m *MainWindow) getIpFromServersList(current_tab string) string {
	// Getting server's address from list.
	// Assuming that we're on "Servers" tab by default.
	sel := m.all_servers.GetSelection()
	model := m.all_servers.GetModel()
	if strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
		sel = m.fav_servers.GetSelection()
		model = m.fav_servers.GetModel()
	}

	iter := new(gtk.TreeIter)
	_ = sel.GetSelected(iter)

	// Getting server address.
	var srv_addr string
	srv_addr_gval := glib.ValueFromNative(srv_addr)

	if strings.Contains(current_tab, ctx.Translator.Translate("Servers", nil)) {
		model.GetValue(iter, m.column_pos["Servers"]["IP"], srv_addr_gval)
	} else if strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
		model.GetValue(iter, m.column_pos["Favorites"]["IP"], srv_addr_gval)
	}
	server_address := srv_addr_gval.GetString()

	return server_address
}
