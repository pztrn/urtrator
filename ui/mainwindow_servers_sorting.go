package ui

import (
    // stdlib
    "strconv"
    "strings"

    // other
    "github.com/mattn/go-gtk/glib"
    "github.com/mattn/go-gtk/gtk"
)

func (m *MainWindow) sortServersByName(model *gtk.TreeModel, a *gtk.TreeIter, b *gtk.TreeIter) int {
    var name1_raw glib.GValue
    var name2_raw glib.GValue

    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    if current_tab == "Servers" {
        model.GetValue(a, m.column_pos["Servers"]["Name"], &name1_raw)
        model.GetValue(b, m.column_pos["Servers"]["Name"], &name2_raw)
    } else if current_tab == "Favorites" {
        model.GetValue(a, m.column_pos["Favorites"]["Name"], &name1_raw)
        model.GetValue(b, m.column_pos["Favorites"]["Name"], &name2_raw)
    } else {
        return 0
    }

    name1 := ctx.Colorizer.ClearFromMarkup(name1_raw.GetString())
    name2 := ctx.Colorizer.ClearFromMarkup(name2_raw.GetString())

    if name1 < name2 {
        return -1
    } else {
        return 1
    }

    return 0
}

func (m *MainWindow) sortServersByPlayers(model *gtk.TreeModel, a *gtk.TreeIter, b *gtk.TreeIter) int {
    var players1_raw glib.GValue
    var players2_raw glib.GValue

    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    if current_tab == "Servers" {
        model.GetValue(a, m.column_pos["Servers"]["Players"], &players1_raw)
        model.GetValue(b, m.column_pos["Servers"]["Players"], &players2_raw)
    } else if current_tab == "Favorites" {
        model.GetValue(a, m.column_pos["Favorites"]["Players"], &players1_raw)
        model.GetValue(b, m.column_pos["Favorites"]["Players"], &players2_raw)
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

func (m *MainWindow) sortServersByPing(model *gtk.TreeModel, a *gtk.TreeIter, b *gtk.TreeIter) int {
    var ping1_raw glib.GValue
    var ping2_raw glib.GValue

    current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
    if current_tab == "Servers" {
        model.GetValue(a, m.column_pos["Servers"]["Ping"], &ping1_raw)
        model.GetValue(b, m.column_pos["Servers"]["Ping"], &ping2_raw)
    } else if current_tab == "Favorites" {
        model.GetValue(a, m.column_pos["Favorites"]["Ping"], &ping1_raw)
        model.GetValue(b, m.column_pos["Favorites"]["Ping"], &ping2_raw)
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
