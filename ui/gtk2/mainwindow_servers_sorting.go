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
	"strconv"
	"strings"

	// other
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func (m *MainWindow) sortServersByName(model *gtk.TreeModel, a *gtk.TreeIter, b *gtk.TreeIter, userData interface{}) int {
	var name1_raw glib.GValue
	var name2_raw glib.GValue

	current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
	if strings.Contains(current_tab, ctx.Translator.Translate("Servers", nil)) {
		model.GetValue(a, m.column_pos["Servers"][ctx.Translator.Translate("Name", nil)], &name1_raw)
		model.GetValue(b, m.column_pos["Servers"][ctx.Translator.Translate("Name", nil)], &name2_raw)
	} else if strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
		model.GetValue(a, m.column_pos["Favorites"][ctx.Translator.Translate("Name", nil)], &name1_raw)
		model.GetValue(b, m.column_pos["Favorites"][ctx.Translator.Translate("Name", nil)], &name2_raw)
	} else {
		return 0
	}

	name1 := strings.ToLower(ctx.Colorizer.ClearFromMarkup(name1_raw.GetString()))
	name2 := strings.ToLower(ctx.Colorizer.ClearFromMarkup(name2_raw.GetString()))

	if name1 < name2 {
		return -1
	} else {
		return 1
	}

	return 0
}

func (m *MainWindow) sortServersByPlayers(model *gtk.TreeModel, a *gtk.TreeIter, b *gtk.TreeIter, userData interface{}) int {
	var players1_raw glib.GValue
	var players2_raw glib.GValue

	current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
	if strings.Contains(current_tab, ctx.Translator.Translate("Servers", nil)) {
		model.GetValue(a, m.column_pos["Servers"][ctx.Translator.Translate("Players", nil)], &players1_raw)
		model.GetValue(b, m.column_pos["Servers"][ctx.Translator.Translate("Players", nil)], &players2_raw)
	} else if strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
		model.GetValue(a, m.column_pos["Favorites"][ctx.Translator.Translate("Players", nil)], &players1_raw)
		model.GetValue(b, m.column_pos["Favorites"][ctx.Translator.Translate("Players", nil)], &players2_raw)
	} else {
		return 0
	}

	players1_online := strings.Split(players1_raw.GetString(), "/")[0]
	players2_online := strings.Split(players2_raw.GetString(), "/")[0]

	if len(players1_online) > 0 && len(players2_online) > 0 {
		players1, _ := strconv.Atoi(players1_online)
		players2, _ := strconv.Atoi(players2_online)
		if players1 > players2 {
			return -1
		} else {
			return 1
		}
	}

	return -1
}

func (m *MainWindow) sortServersByPing(model *gtk.TreeModel, a *gtk.TreeIter, b *gtk.TreeIter, userData interface{}) int {
	var ping1_raw glib.GValue
	var ping2_raw glib.GValue

	current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
	if strings.Contains(current_tab, ctx.Translator.Translate("Servers", nil)) {
		model.GetValue(a, m.column_pos["Servers"][ctx.Translator.Translate("Ping", nil)], &ping1_raw)
		model.GetValue(b, m.column_pos["Servers"][ctx.Translator.Translate("Ping", nil)], &ping2_raw)
	} else if strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
		model.GetValue(a, m.column_pos["Favorites"][ctx.Translator.Translate("Ping", nil)], &ping1_raw)
		model.GetValue(b, m.column_pos["Favorites"][ctx.Translator.Translate("Ping", nil)], &ping2_raw)
	} else {
		return 0
	}

	ping1, _ := strconv.Atoi(ping1_raw.GetString())
	ping2, _ := strconv.Atoi(ping2_raw.GetString())

	if ping1 < ping2 {
		return 1
	} else {
		return -1
	}

	return -1
}
