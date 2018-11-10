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
