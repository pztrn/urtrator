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
	"errors"
	"fmt"
	"runtime"
	"strings"

	// Local
	"gitlab.com/pztrn/urtrator/datamodels"

	// other
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func (m *MainWindow) launchGame() error {
	fmt.Println("Launching Urban Terror...")
	if len(m.qc_server_address.GetText()) != 0 {
		m.launchWithQuickConnect()
	} else {
		m.launchAsUsual()
	}

	return nil
}

// Triggers if we clicked "Launch" button without any text in quick connect
// widget.
func (m *MainWindow) launchAsUsual() error {
	fmt.Println("Connecting to selected server...")
	var srv_address string = ""

	// Getting server's name from list.
	current_tab := m.tab_widget.GetTabLabelText(m.tab_widget.GetNthPage(m.tab_widget.GetCurrentPage()))
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
	srv_address_gval := glib.ValueFromNative(srv_addr)
	if strings.Contains(current_tab, ctx.Translator.Translate("Servers", nil)) {
		model.GetValue(iter, m.column_pos["Servers"]["IP"], srv_address_gval)
	} else if strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
		model.GetValue(iter, m.column_pos["Favorites"]["IP"], srv_address_gval)
	}
	srv_address = srv_address_gval.GetString()
	if len(srv_address) == 0 {
		// Temporary disable all these modals on Linux.
		// See https://github.com/mattn/go-gtk/issues/289.
		if runtime.GOOS != "linux" {
			mbox_string := ctx.Translator.Translate("No server selected.\n\nPlease, select a server to continue connecting.", nil)
			messagebox := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
			messagebox.Response(func() {
				messagebox.Destroy()
			})
			messagebox.Run()
		} else {
			ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">" + ctx.Translator.Translate("Select a server we will connect to!", nil) + "</span></markup>"})
		}
		return errors.New(ctx.Translator.Translate("No server selected.", nil))
	}
	server_profile := ctx.Cache.Servers[srv_address].Server

	// Check for proper server name. If length == 0: server is offline,
	// we should show notification to user.
	if len(server_profile.Name) == 0 {
		var will_continue bool = false
		// Temporary disable all these modals on Linux.
		// See https://github.com/mattn/go-gtk/issues/289.
		if runtime.GOOS != "linux" {
			mbox_string := ctx.Translator.Translate("Selected server is offline.\n\nWould you still want to launch Urban Terror?\nIt will just launch a game, without connecting to\nany server.", nil)
			messagebox := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_YES_NO, mbox_string)
			messagebox.Connect("response", func(resp *glib.CallbackContext) {
				if resp.Args(0) == 4294967287 {
					will_continue = false
				} else {
					will_continue = true
				}
				messagebox.Destroy()
			})
			messagebox.Run()
		} else {
			// We're okay to connect to empty server, temporary.
			will_continue = true
		}

		if !will_continue {
			return errors.New(ctx.Translator.Translate("User declined to connect to offline server", nil))
		}
	}

	// Getting selected profile's name.
	profile_name := m.profiles.GetActiveText()
	user_profile := &datamodels.Profile{}
	if strings.Contains(current_tab, ctx.Translator.Translate("Servers", nil)) {
		// Checking profile name length. If 0 - then stop executing :)
		// This check only relevant to "Servers" tab, favorite servers
		// have profiles defined (see next).
		if len(profile_name) == 0 {
			// Temporary disable all these modals on Linux.
			// See https://github.com/mattn/go-gtk/issues/289.
			if runtime.GOOS != "linux" {
				mbox_string := ctx.Translator.Translate("Invalid game profile selected.\n\nPlease, select profile and retry.", nil)
				messagebox := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
				messagebox.Response(func() {
					messagebox.Destroy()
				})
				messagebox.Run()
			} else {
				ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">" + ctx.Translator.Translate("Invalid game profile selected.", nil) + "</span></markup>"})
			}
			return errors.New(ctx.Translator.Translate("User didn't select valid profile.", nil))
		}
		user_profile = ctx.Cache.Profiles[profile_name].Profile
	} else if strings.Contains(current_tab, ctx.Translator.Translate("Favorites", nil)) {
		// For favorite servers profile specified in favorite server
		// information have higher priority, so we just override it :)
		user_profile_cached, ok := ctx.Cache.Profiles[server_profile.ProfileToUse]
		if !ok {
			// Temporary disable all these modals on Linux.
			// See https://github.com/mattn/go-gtk/issues/289.
			if runtime.GOOS != "linux" {
				mbox_string := ctx.Translator.Translate("Invalid game profile specified for favorite server.\n\nPlease, edit your favorite server, select valid profile and retry.", nil)
				messagebox := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
				messagebox.Response(func() {
					messagebox.Destroy()
				})
				messagebox.Run()
			} else {
				ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">" + ctx.Translator.Translate("Invalid game profile specified in favorite entry.", nil) + "</span></markup>"})
			}
			return errors.New(ctx.Translator.Translate("User didn't select valid profile.", nil))
		}
		user_profile = user_profile_cached.Profile
	}

	m.launchActually(server_profile, user_profile, "", "")
	return nil
}

// Triggers when Launch button was clicked with some text in quick connect
// widget.
func (m *MainWindow) launchWithQuickConnect() error {
	fmt.Println("Launching game with data from quick connect...")

	srv_address := m.qc_server_address.GetText()
	srv_password := m.qc_password.GetText()
	srv_nickname := m.qc_nickname.GetText()
	current_profile_name := m.profiles.GetActiveText()

	// As we're launching without any profile defined - we should
	// check server version and globally selected profile.
	// Checking if we have server defined in cache.
	var ip string = ""
	var port string = ""
	if strings.Contains(srv_address, ":") {
		ip = strings.Split(srv_address, ":")[0]
		port = strings.Split(srv_address, ":")[1]
	} else {
		ip = strings.Split(srv_address, ":")[0]
		port = "27960"
	}

	key := ip + ":" + port

	_, ok := ctx.Cache.Servers[key]
	if !ok {
		ctx.Cache.CreateServer(key)
		fmt.Println("Server not found in cache, requesting information...")
		ctx.Requester.UpdateOneServer(key)
	}

	server_profile := ctx.Cache.Servers[key]
	user_profile := ctx.Cache.Profiles[current_profile_name]

	m.launchActually(server_profile.Server, user_profile.Profile, srv_password, srv_nickname)
	return nil
}

func (m *MainWindow) launchActually(server_profile *datamodels.Server, user_profile *datamodels.Profile, password string, nickname_to_use string) error {
	if server_profile.Name == "" {
		var will_continue bool = false
		// Temporary disable all these modals on Linux.
		// See https://github.com/mattn/go-gtk/issues/289.
		if runtime.GOOS != "linux" {
			mbox_string := ctx.Translator.Translate("Selected server is offline.\n\nWould you still want to launch Urban Terror?\nIt will just launch a game, without connecting to\nany server.", nil)
			messagebox := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_YES_NO, mbox_string)
			messagebox.Connect("response", func(resp *glib.CallbackContext) {
				if resp.Args(0) == 4294967287 {
					will_continue = false
				} else {
					will_continue = true
				}
				messagebox.Destroy()
			})
			messagebox.Run()
		} else {
			// We're ok here, temporary.
			will_continue = true
		}

		if !will_continue {
			return errors.New(ctx.Translator.Translate("User declined to connect to offline server", nil))
		}
	}

	// Check if server is applicable for selected profile.
	if server_profile.Version != user_profile.Version {
		// Temporary disable all these modals on Linux.
		// See https://github.com/mattn/go-gtk/issues/289.
		if runtime.GOOS != "linux" {
			mbox_string := ctx.Translator.Translate("Invalid game profile selected.\n\nSelected profile have different game version than server.\nPlease, select valid profile and retry.", nil)
			messagebox := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
			messagebox.Response(func() {
				messagebox.Destroy()
			})
			messagebox.Run()
		} else {
			ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">" + ctx.Translator.Translate("Invalid game profile selected.", nil) + "</span></markup>"})
		}
		return errors.New(ctx.Translator.Translate("User didn't select valid profile, mismatch with server's version.", nil))
	}

	server_password := password
	if len(server_password) == 0 {
		server_password = server_profile.Password
	}

	// Hey, we're ok here! :) Launch Urban Terror!
	// Clear server name from "<markup></markup>" things.
	srv_name_for_label := server_profile.Name
	if strings.Contains(server_profile.Name, "markup") {
		srv_name_for_label = string([]byte(server_profile.Name)[8 : len(server_profile.Name)-9])
	} else {
		srv_name := ctx.Colorizer.Fix(server_profile.Name)
		srv_name_for_label = string([]byte(srv_name)[8 : len(srv_name)-9])
	}
	// Show great coloured label.
	ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">" + ctx.Translator.Translate("Urban Terror is launched with profile", nil) + " </span><span foreground=\"blue\">" + user_profile.Name + "</span> <span foreground=\"red\" font_weight=\"bold\">" + ctx.Translator.Translate("and connected to", nil) + " </span><span foreground=\"orange\" font_weight=\"bold\">" + srv_name_for_label + "</span></markup>"})
	m.launch_button.SetSensitive(false)
	// ToDo: handling server passwords.
	ctx.Launcher.Launch(server_profile, user_profile, server_password, []string{"+name", nickname_to_use}, m.unlockInterface)

	return nil
}
