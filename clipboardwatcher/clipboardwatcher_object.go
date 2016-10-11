// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package clipboardwatcher

import (
    // stdlib
    "fmt"
    "strings"

    // other
    "github.com/mattn/go-gtk/gdk"
    "github.com/mattn/go-gtk/gtk"
)

type ClipboardWatcher struct {
    // Clipboard.
    clipboard *gtk.Clipboard
    // PRIMARY clipboard.
    prim_clipboard *gtk.Clipboard
}

func (cw *ClipboardWatcher) checkInput() {
    text := cw.clipboard.WaitForText()
    cw.parseData(text)
}

func (cw *ClipboardWatcher) checkPrimaryInput() {
    text := cw.prim_clipboard.WaitForText()
    cw.parseData(text)
}

func (cw *ClipboardWatcher) Initialize() {
    fmt.Println("Initializing clipboard watcher...")

    cw.clipboard = gtk.NewClipboardGetForDisplay(gdk.DisplayGetDefault(), gdk.SELECTION_CLIPBOARD)
    cw.clipboard.Connect("owner-change", cw.checkInput)

    cw.prim_clipboard = gtk.NewClipboardGetForDisplay(gdk.DisplayGetDefault(), gdk.SELECTION_PRIMARY)
    cw.prim_clipboard.Connect("owner-change", cw.checkPrimaryInput)
}

func (cw *ClipboardWatcher) parseData(data string) {
    fmt.Println(data)
    // Checking if we have connection string here.
    if strings.Contains(data, "connect") && strings.Contains(data, "password") {
        fmt.Println("Connection string detected!")
        var server string = ""
        var password string = ""
        conn_string := strings.Split(data, ";")
        if len(conn_string) > 0 {
            srv_string := strings.Split(data, ";")[0]
            srv_splitted := strings.Split(srv_string, "connect ")
            if len(srv_splitted) > 1 {
                server_raw := srv_splitted[1]
                // Get rid of spaces.
                server_raw_splitted := strings.Split(server_raw, " ")
                for i := range server_raw_splitted {
                    if server_raw_splitted[i] == "" {
                        continue
                    }
                    server = server_raw_splitted[i]
                }
            }
        }
        if len(conn_string) > 1 {
            pw_string := strings.Split(data, ";")[1]
            pw_splitted := strings.Split(pw_string, "password ")
            if len(pw_splitted) > 1 {
                password_raw := pw_splitted[1]
                // Get rid of spaces.
                password_raw_splitted := strings.Split(password_raw, " ")
                for i := range password_raw_splitted {
                    if password_raw_splitted[i] == "" {
                        continue
                    }
                    password = password_raw_splitted[i]
                }
            }
        }
        Eventer.LaunchEvent("setQuickConnectDetails", map[string]string{"server": server, "password": password})
    }
}
