// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package ui

import (
    // stdlib
    "fmt"

    // Local
    //"github.com/pztrn/urtrator/datamodels"

    // Other
    "github.com/mattn/go-gtk/gtk"
    //"github.com/mattn/go-gtk/glib"
)

type ServerInfoDialog struct {
    // Window.
    window *gtk.Window
    // Main Vertical Box.
    vbox *gtk.VBox
}

func (sid *ServerInfoDialog) Initialize() {
    fmt.Println("Showing server's information...")
}
