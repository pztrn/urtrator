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
    "runtime"
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
    // Server's we're working with.
    server *datamodels.Server
    // Profiles count that was added to profiles combobox.
    profiles int
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
    // Remove old profiles.
    if f.profiles > 0 {
        for i := 0; i <= f.profiles; i++ {
            f.profile.RemoveText(0)
        }
    }

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
            f.profiles += 1
        }
    }

    f.profile.SetActive(idx_should_be_active)
}

func (f *FavoriteDialog) InitializeNew() {
    f.update = false
    f.server = &datamodels.Server{}
    f.profiles = 0
    f.initializeWindow()
}

func (f *FavoriteDialog) InitializeUpdate(server *datamodels.Server) {
    fmt.Println("Favorites updating...")
    f.update = true
    f.server = server
    f.profiles = 0
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

    // Set some GTK options for this window.
    gtk_opts_raw := gtk.SettingsGetDefault()
    gtk_opts := gtk_opts_raw.ToGObject()
    gtk_opts.Set("gtk-button-images", true)

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
    srv_addr_update_btn := gtk.NewButton()
    srv_addr_update_btn.SetTooltipText("Update server information")
    srv_addr_update_btn_image := gtk.NewImageFromStock(gtk.STOCK_REDO, gtk.ICON_SIZE_SMALL_TOOLBAR)
    srv_addr_update_btn.SetImage(srv_addr_update_btn_image)
    srv_addr_update_btn.Clicked(f.updateServerInfo)
    srv_addr_hbox.PackStart(srv_addr_update_btn, false, true, 5)
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
    // Update server's information.
    f.server.Name = f.server_name.GetText()
    //ctx.Requester.Pooler.UpdateSpecificServer(f.server)

    if len(f.server_address.GetText()) == 0 {
        // Temporary disable all these modals on Linux.
        // See https://github.com/mattn/go-gtk/issues/289.
        if runtime.GOOS != "linux" {
            mbox_string := "Server address is empty.\n\nServers without address cannot be added."
            m := gtk.NewMessageDialog(f.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
            m.Response(func() {
                m.Destroy()
            })
            m.Run()
        }
        return errors.New("No server address specified")
    }

    var port string = ""
    if strings.Contains(f.server_address.GetText(), ":") {
        port = strings.Split(f.server_address.GetText(), ":")[1]
    } else {
        port = "27960"
    }
    f.server.Ip = strings.Split(f.server_address.GetText(), ":")[0]
    f.server.Port = port

    if len(f.profile.GetActiveText()) == 0 {
        // Temporary disable all these modals on Linux.
        // See https://github.com/mattn/go-gtk/issues/289.
        if runtime.GOOS != "linux" {
            mbox_string := "Profile wasn't selected.\n\nPlease, select valid profile for this server.\nIf you haven't add profiles yet - you can do it\nin options on \"Urban Terror\" tab."
            m := gtk.NewMessageDialog(f.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
            m.Response(func() {
                m.Destroy()
            })
            m.Run()
        }
        return errors.New("No game profile specified")
    }

    fmt.Println("Saving favorite server...")

    key := strings.Split(f.server_address.GetText(), ":")[0] + ":" + port
    ctx.Cache.Servers[key].Server.Ip = f.server.Ip
    ctx.Cache.Servers[key].Server.Port = f.server.Port
    ctx.Cache.Servers[key].Server.Name = f.server.Name
    ctx.Cache.Servers[key].Server.Password = f.server_password.GetText()
    ctx.Cache.Servers[key].Server.ProfileToUse = f.profile.GetActiveText()
    ctx.Cache.Servers[key].Server.Favorite = "1"
    ctx.Cache.Servers[key].Server.ExtendedConfig = f.server.ExtendedConfig
    ctx.Cache.Servers[key].Server.PlayersInfo = f.server.PlayersInfo

    ctx.Eventer.LaunchEvent("flushServers", map[string]string{})
    ctx.Eventer.LaunchEvent("loadFavoriteServers", map[string]string{})
    f.window.Destroy()

    return nil
}

func (f *FavoriteDialog) updateServerInfo() {
    fmt.Println("Updating server information...")
    var port string = ""
    if strings.Contains(f.server_address.GetText(), ":") {
        port = strings.Split(f.server_address.GetText(), ":")[1]
    } else {
        port = "27960"
    }
    f.server.Ip = strings.Split(f.server_address.GetText(), ":")[0]
    f.server.Port = port

    ctx.Requester.Pooler.UpdateSpecificServer(f.server)

    f.fill()
}
