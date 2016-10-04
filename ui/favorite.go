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
    "encoding/base64"
    "errors"
    "fmt"
    "strings"

    // Local
    "github.com/pztrn/urtrator/common"
    "github.com/pztrn/urtrator/datamodels"

    // Other
    "github.com/mattn/go-gtk/gdkpixbuf"
    "github.com/mattn/go-gtk/gtk"
)

type FavoriteDialog struct {
    // Widgets.
    // Dialog's window.
    window *gtk.Window
    // Main vertical box.
    vbox *gtk.VBox
    // Server name.
    server_name *gtk.Entry
    // Server address.
    server_address *gtk.Entry
    // Server password
    server_password *gtk.Entry
    // Profile.
    profile *gtk.ComboBoxText

    // Flags.
    // Is known server update performed?
    update bool

    // Data.
    // If known server is used - here server's datamodel is.
    server *datamodels.Server
}

func (f *FavoriteDialog) Close() {}

func (f *FavoriteDialog) closeByCancel() {
    f.window.Destroy()
}

func (f *FavoriteDialog) fill() {
    f.server_name.SetText(f.server.Name)
    f.server_address.SetText(f.server.Ip + ":" + f.server.Port)
    f.server_password.SetText(f.server.Password)

    // Profiles.
    profiles := []datamodels.Profile{}
    err := ctx.Database.Db.Select(&profiles, "SELECT * FROM urt_profiles")
    if err != nil {
        fmt.Println(err.Error())
    }
    var idx_in_combobox int = 0
    var idx_should_be_active int = 0
    for p := range profiles {
        if profiles[p].Version == f.server.Version {
            f.profile.AppendText(profiles[p].Name)
            idx_should_be_active = idx_in_combobox
            idx_in_combobox += 1
        }
    }

    f.profile.SetActive(idx_should_be_active)
}

func (f *FavoriteDialog) InitializeNew() {
    f.update = false
    f.initializeWindow()
}

func (f *FavoriteDialog) InitializeUpdate(server *datamodels.Server) {
    fmt.Println("Favorites updating...")
    f.update = true
    f.server = server
    f.initializeWindow()
    f.fill()
}

func (f *FavoriteDialog) initializeWindow() {
    f.window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
    if f.update {
        f.window.SetTitle("URTrator - Updating favorite server")
    } else {
        f.window.SetTitle("URTrator - New favorite server")
    }
    f.window.Connect("destroy", f.Close)
    f.window.SetPosition(gtk.WIN_POS_CENTER)
    f.window.SetModal(true)
    f.window.SetSizeRequest(400, 200)
    f.window.SetResizable(false)
    f.vbox = gtk.NewVBox(false, 0)

    // Load program icon from base64.
    icon_bytes, _ := base64.StdEncoding.DecodeString(common.Logo)
    icon_pixbuf := gdkpixbuf.NewLoader()
    icon_pixbuf.Write(icon_bytes)
    logo = icon_pixbuf.GetPixbuf()
    f.window.SetIcon(logo)

    // Server name.
    srv_name_hbox := gtk.NewHBox(false, 0)
    f.vbox.PackStart(srv_name_hbox, false, true, 5)
    srv_name_label := gtk.NewLabel("Server name:")
    srv_name_hbox.PackStart(srv_name_label, false, true, 5)
    srv_name_sep := gtk.NewHSeparator()
    srv_name_hbox.PackStart(srv_name_sep, true, true, 5)
    f.server_name = gtk.NewEntry()
    srv_name_hbox.PackStart(f.server_name, true, true, 5)

    // Server address.
    srv_addr_hbox := gtk.NewHBox(false, 0)
    f.vbox.PackStart(srv_addr_hbox, false, true, 5)
    srv_addr_label := gtk.NewLabel("Server address:")
    srv_addr_hbox.PackStart(srv_addr_label, false, true, 5)
    srv_addr_sep := gtk.NewHSeparator()
    srv_addr_hbox.PackStart(srv_addr_sep, true, true, 5)
    f.server_address = gtk.NewEntry()
    srv_addr_hbox.PackStart(f.server_address, true, true, 5)
    if f.update {
        f.server_address.SetSensitive(false)
    }

    // Server password.
    srv_pass_hbox := gtk.NewHBox(false, 0)
    f.vbox.PackStart(srv_pass_hbox, false, true, 5)
    srv_pass_label := gtk.NewLabel("Password:")
    srv_pass_hbox.PackStart(srv_pass_label, false, true, 5)
    srv_pass_sep := gtk.NewHSeparator()
    srv_pass_hbox.PackStart(srv_pass_sep, true, true, 5)
    f.server_password = gtk.NewEntry()
    srv_pass_hbox.PackStart(f.server_password, true, true, 5)

    // Profile to use.
    profile_hbox := gtk.NewHBox(false, 0)
    f.vbox.PackStart(profile_hbox, false, true, 5)
    profile_label := gtk.NewLabel("Profile:")
    profile_hbox.PackStart(profile_label, false, true, 5)
    profile_sep := gtk.NewHSeparator()
    profile_hbox.PackStart(profile_sep, true, true, 5)
    f.profile = gtk.NewComboBoxText()
    profile_hbox.PackStart(f.profile, false, true, 5)

    // Buttons hbox.
    buttons_hbox := gtk.NewHBox(false, 0)
    sep := gtk.NewHSeparator()
    buttons_hbox.PackStart(sep, true, true, 5)
    // OK-Cancel buttons.
    cancel_button := gtk.NewButtonWithLabel("Cancel")
    cancel_button.Clicked(f.closeByCancel)
    buttons_hbox.PackStart(cancel_button, false, true, 5)

    ok_button := gtk.NewButtonWithLabel("OK")
    ok_button.Clicked(f.saveFavorite)
    buttons_hbox.PackStart(ok_button, false, true, 5)

    f.vbox.PackStart(buttons_hbox, false, true, 5)

    f.window.Add(f.vbox)
    f.window.ShowAll()
}

func (f *FavoriteDialog) saveFavorite() error {
    if len(f.server_address.GetText()) == 0 {
        mbox_string := "Server address is empty.\n\nServers without address cannot be added."
        m := gtk.NewMessageDialog(f.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
        return errors.New("No server address specified")
    }

    if len(f.profile.GetActiveText()) == 0 {
        mbox_string := "Profile wasn't selected.\n\nPlease, select valid profile for this server.\nIf you haven't add profiles yet - you can do it\nin options on \"Urban Terror\" tab."
        m := gtk.NewMessageDialog(f.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
        return errors.New("No game profile specified")
    }

    var port string = ""
    if strings.Contains(f.server_address.GetText(), ":") {
        port = strings.Split(f.server_address.GetText(), ":")[1]
    } else {
        port = "27960"
    }

    fmt.Println("Saving favorite server...")

    server := datamodels.Server{}
    server.Ip = strings.Split(f.server_address.GetText(), ":")[0]
    server.Port = port
    server.Name = f.server_name.GetText()
    server.Password = f.server_password.GetText()
    server.ProfileToUse = f.profile.GetActiveText()
    server.Favorite = "1"

    if f.update {
        q := "UPDATE servers SET name=:name, ip=:ip, port=:port, password=:password, favorite=:favorite, profile_to_use=:profile_to_use WHERE ip='" + f.server.Ip + "' AND port='" + f.server.Port + "'"
        fmt.Println("Query: " + q)
        ctx.Database.Db.NamedExec(q, &server)
    } else {
        q := "INSERT INTO servers (name, ip, port, password, favorite, profile_to_use) VALUES (:name, :ip, :port, :password, \"1\", :profile_to_use)"
        fmt.Println(q)
        ctx.Database.Db.NamedExec(q, &server)
    }

    ctx.Eventer.LaunchEvent("loadFavoriteServers")
    f.window.Destroy()

    return nil
}
