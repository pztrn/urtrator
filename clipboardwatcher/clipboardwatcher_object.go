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

package clipboardwatcher

import (
	// stdlib
	"errors"
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

	// Flags.
	// We have just copy connect string to clipboard.
	// Used to ignore clipboard data in check*Input()
	just_set bool
}

func (cw *ClipboardWatcher) checkInput() {
	if !cw.just_set {
		text := cw.clipboard.WaitForText()
		cw.parseData(text)
	} else {
		cw.just_set = false
	}
}

func (cw *ClipboardWatcher) checkPrimaryInput() {
	if !cw.just_set {
		text := cw.prim_clipboard.WaitForText()
		cw.parseData(text)
	} else {
		cw.just_set = false
	}
}

func (cw *ClipboardWatcher) CopyServerData(server_address string) error {
	server, ok := Cache.Servers[server_address]
	if !ok {
		// ToDo: show message box?
		return errors.New("Server wasn't selected")
	}

	// Composing connection string.
	var connect_string string = ""
	connect_string += "/connect " + server.Server.Ip + ":" + server.Server.Port
	if len(server.Server.Password) >= 1 {
		connect_string += ";password " + server.Server.Password
	}
	fmt.Println("Connect string: ", connect_string)
	cw.just_set = true
	cw.clipboard.SetText(connect_string)

	return nil
}

func (cw *ClipboardWatcher) Initialize() {
	fmt.Println("Initializing clipboard watcher...")

	cw.just_set = false

	cw.clipboard = gtk.NewClipboardGetForDisplay(gdk.DisplayGetDefault(), gdk.SELECTION_CLIPBOARD)
	cw.clipboard.Connect("owner-change", cw.checkInput)

	cw.prim_clipboard = gtk.NewClipboardGetForDisplay(gdk.DisplayGetDefault(), gdk.SELECTION_PRIMARY)
	cw.prim_clipboard.Connect("owner-change", cw.checkPrimaryInput)
}

func (cw *ClipboardWatcher) parseData(data string) {
	// We should check only first string.
	data = strings.Split(data, "\n")[0]
	// Checking if we have connection string here.
	if strings.Contains(data, "ct ") {
		fmt.Println("Connection string detected!")
		var server string = ""
		var password string = ""
		conn_string := strings.Split(data, ";")
		if len(conn_string) > 0 {
			srv_string := strings.Split(data, ";")[0]
			srv_splitted := strings.Split(srv_string, "ct ")
			if len(srv_splitted) > 1 {
				server_raw := strings.Split(srv_splitted[1], " ")[0]
				// Get rid of spaces.
				server = strings.TrimSpace(server_raw)
			}
		}
		if len(conn_string) > 1 && strings.Contains(data, "password") {
			pw_string := strings.Split(data, ";")[1]
			pw_splitted := strings.Split(pw_string, "password ")
			if len(pw_splitted) > 1 {
				password_raw := strings.Split(pw_splitted[1], " ")[0]
				// Get rid of spaces.
				password = strings.TrimSpace(password_raw)
			}
		}
		Eventer.LaunchEvent("setQuickConnectDetails", map[string]string{"server": server, "password": password})
	}
}
