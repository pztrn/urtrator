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
