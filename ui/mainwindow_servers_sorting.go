package ui

import (
    // stdlib
    "strconv"
    "strings"

    // other
    "github.com/mattn/go-gtk/glib"
    "github.com/mattn/go-gtk/gtk"
)

func (m *MainWindow) sortServersByName() {

}

func (m *MainWindow) sortServersByPlayers(model *gtk.TreeModel, a *gtk.TreeIter, b *gtk.TreeIter) int {
    var players1_raw glib.GValue
    var players2_raw glib.GValue
    model.GetValue(a, m.column_pos["Servers"]["Players"], &players1_raw)
    model.GetValue(b, m.column_pos["Servers"]["Players"], &players2_raw)

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
    model.GetValue(a, m.column_pos["Servers"]["Ping"], &ping1_raw)
    model.GetValue(b, m.column_pos["Servers"]["Ping"], &ping2_raw)

    ping1, _ := strconv.Atoi(ping1_raw.GetString())
    ping2, _ := strconv.Atoi(ping2_raw.GetString())

    if ping1 < ping2 {
        return 1
    } else {
        return -1
    }

    return -1
}
