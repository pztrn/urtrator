package ui

import (
    // other
    "github.com/couchbase/goutils/platform"
)

func (m *MainWindow) initializeWin() {
    platform.HideConsole(true)

    dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

    // ToDo: theming support and theme seletion in settings.
    gtk.RCParse(dir + "/themes/MS-Windows/gtk-2.0/gtkrc")
}
