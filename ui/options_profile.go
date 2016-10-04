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
    "os"
    "runtime"
    "strings"

    // Local
    "github.com/pztrn/urtrator/datamodels"

    // Other
    "github.com/mattn/go-gtk/gtk"
    //"github.com/mattn/go-gtk/glib"
)

type OptionsProfile struct {
    // Window.
    window *gtk.Window
    // Main Vertical Box.
    vbox *gtk.VBox
    // Profile name.
    profile_name *gtk.Entry
    // Binary path.
    binary_path *gtk.Entry
    // Urban Terror versions combobox
    urt_version_combo *gtk.ComboBoxText
    // Another X session?
    another_x_session *gtk.CheckButton
    // Additional parameters for game launching.
    additional_parameters *gtk.Entry

    // File chooser dialog for selecting binary.
    f *gtk.FileChooserDialog

    // Flags.
    // This is profile update?
    update bool

    // Callbacks.
    // This will be triggered after we change profile.
    loadProfiles func()

    // Others.
    // Old profile, needed for proper update.
    old_profile *datamodels.Profile
}

func (op *OptionsProfile) browseForBinary() {
    op.f = gtk.NewFileChooserDialog("URTrator - Select Urban Terro binary", op.window, gtk.FILE_CHOOSER_ACTION_OPEN, gtk.STOCK_OK, gtk.RESPONSE_ACCEPT)
    op.f.Response(op.browseForBinaryHelper)
    op.f.Run()
}

func (op *OptionsProfile) browseForBinaryHelper() {
    filename := op.f.GetFilename()
    op.binary_path.SetText(filename)
    op.f.Destroy()

    // Check for valid filename.
    // ToDo: add more OSes.
    if runtime.GOOS == "linux" {
        // Filename should end with approriate arch.
        if runtime.GOARCH == "amd64" {
            if len(filename) > 0 && strings.Split(filename, ".")[1] != "x86_64" && strings.Split(filename, ".")[0] != "Quake3-UrT" {
                fmt.Println("Invalid binary selected!")
                mbox_string := "Invalid binary selected!\nAccording to your OS, it should be Quake3-UrT.x86_64."
                m := gtk.NewMessageDialog(op.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
                m.Response(func() {
                    m.Destroy()
                })
                m.Run()
                op.binary_path.SetText("")
            }
        }
    }
}

func (op *OptionsProfile) closeByCancel() {
    op.window.Destroy()
}

func (op *OptionsProfile) closeWithDiscard() {
}

func (op *OptionsProfile) Initialize(update bool, lp func()) {
    if update {
        op.update = true
    }

    op.loadProfiles = lp

    op.window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
    if update {
        op.window.SetTitle("URTrator - Update Urban Terror profile")
    } else {
        op.window.SetTitle("URTrator - Add Urban Terror profile")
    }
    op.window.Connect("destroy", op.closeWithDiscard)
    op.window.SetModal(true)
    op.window.SetSizeRequest(550, 400)
    op.window.SetPosition(gtk.WIN_POS_CENTER)
    op.window.SetIcon(logo)

    op.vbox = gtk.NewVBox(false, 0)


    // Profile name.
    pn_hbox := gtk.NewHBox(false, 0)
    pn_label := gtk.NewLabel("Profile name:")
    sep1 := gtk.NewHSeparator()
    op.profile_name = gtk.NewEntry()
    pn_hbox.PackStart(pn_label, false, true, 5)
    pn_hbox.PackStart(sep1, true, true, 5)
    pn_hbox.PackStart(op.profile_name, true, true, 5)
    op.vbox.PackStart(pn_hbox, false, true, 5)

    // Urban Terror version.
    urt_version_hbox := gtk.NewHBox(false, 0)
    urt_version_label := gtk.NewLabel("Urban Terror version:")
    urt_version_sep := gtk.NewHSeparator()
    op.urt_version_combo = gtk.NewComboBoxText()
    op.urt_version_combo.AppendText("4.2.023")
    op.urt_version_combo.AppendText("4.3.000")
    op.urt_version_combo.SetActive(1)
    urt_version_hbox.PackStart(urt_version_label, false, true, 5)
    urt_version_hbox.PackStart(urt_version_sep, true, true, 5)
    urt_version_hbox.PackStart(op.urt_version_combo, true, true, 5)
    op.vbox.PackStart(urt_version_hbox, false, true, 5)

    // Urban Terror binary path.
    binpath_hbox := gtk.NewHBox(false, 0)
    binpath_label := gtk.NewLabel("Urban Terror binary:")
    sep2 := gtk.NewHSeparator()
    op.binary_path = gtk.NewEntry()
    button_select_binary := gtk.NewButtonWithLabel("Browse")
    button_select_binary.Clicked(op.browseForBinary)
    binpath_hbox.PackStart(binpath_label, false, true, 5)
    binpath_hbox.PackStart(sep2, true, true, 5)
    binpath_hbox.PackStart(op.binary_path, true, true, 5)
    binpath_hbox.PackStart(button_select_binary, false, true, 5)
    op.vbox.PackStart(binpath_hbox, false, true, 5)

    // Should we use additional X session?
    op.another_x_session = gtk.NewCheckButtonWithLabel("Start Urban Terror in another X session?")
    op.vbox.PackStart(op.another_x_session, false, true, 5)

    // Additional game parameters.
    params_hbox := gtk.NewHBox(false, 0)
    params_label := gtk.NewLabel("Additional parameters:")
    params_sep := gtk.NewHSeparator()
    op.additional_parameters = gtk.NewEntry()
    params_hbox.PackStart(params_label, false, true, 5)
    params_hbox.PackStart(params_sep, true, true, 5)
    params_hbox.PackStart(op.additional_parameters, true, true, 5)
    op.vbox.PackStart(params_hbox, false, true, 5)

    // Vertical separator.
    vert_sep := gtk.NewVSeparator()
    op.vbox.PackStart(vert_sep, true, true, 5)

    // The buttons.
    buttons_box := gtk.NewHBox(false, 0)
    buttons_sep := gtk.NewHSeparator()

    cancel_button := gtk.NewButtonWithLabel("Cancel")
    cancel_button.Clicked(op.closeByCancel)
    buttons_box.PackStart(cancel_button, false, true, 5)

    buttons_box.PackStart(buttons_sep, true, true, 5)

    add_button := gtk.NewButton()
    if op.update {
        add_button.SetLabel("Update")
    } else {
        add_button.SetLabel("Add")
    }
    add_button.Clicked(op.saveProfile)
    buttons_box.PackStart(add_button, false, true, 5)

    op.vbox.PackStart(buttons_box, false, true, 5)

    op.window.Add(op.vbox)
    op.window.ShowAll()
}

func (op *OptionsProfile) InitializeUpdate(profile_name string, lp func()) {
    fmt.Println("Updating profile '" + profile_name + "'")
    op.Initialize(true, lp)

    // Get profile data.
    profile := []datamodels.Profile{}
    err := ctx.Database.Db.Select(&profile, ctx.Database.Db.Rebind("SELECT * FROM urt_profiles WHERE name=?"), profile_name)
    if err != nil {
        fmt.Println(err.Error())
    }

    op.profile_name.SetText(profile[0].Name)
    op.binary_path.SetText(profile[0].Binary)
    op.additional_parameters.SetText(profile[0].Additional_params)
    if profile[0].Second_x_session == "1" {
        op.another_x_session.SetActive(true)
    }

    if profile[0].Version == "4.3.000" {
        op.urt_version_combo.SetActive(1)
    } else {
        op.urt_version_combo.SetActive(0)
    }

    op.old_profile = &profile[0]

}

func (op *OptionsProfile) saveProfile() {
    fmt.Println("Saving profile...")

    // Validating fields.
    // Profile name must not be empty.
    if len(op.profile_name.GetText()) < 1 {
        mbox_string := "Empty profile name!\nProfile must be named somehow."
        m := gtk.NewMessageDialog(op.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
    }
    // Binary path must also be filled.
    if len(op.binary_path.GetText()) < 1 {
        mbox_string := "Empty path to binary!\nThis profile will be unusable if you\nwill not provide path to binary!"
        m := gtk.NewMessageDialog(op.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
    }
    // ...and must be executable! :)
    filestat, err := os.Stat(op.binary_path.GetText())
    if err != nil {
        mbox_string := "Invalid path to binary!\n\nError was:\n" + err.Error() + "\n\nCheck binary path and try again."
        m := gtk.NewMessageDialog(op.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
    } else {
        // ToDo: executable flag checking.
        fmt.Println(filestat.Mode())

        // If we here - we can try to save game profile :).
        profile := datamodels.Profile{
            Name:               op.profile_name.GetText(),
            Version:            op.urt_version_combo.GetActiveText(),
            Binary:             op.binary_path.GetText(),
            Additional_params:  op.additional_parameters.GetText(),
        }

        if op.another_x_session.GetActive() {
            profile.Second_x_session = "1"
        } else {
            profile.Second_x_session = "0"
        }

        // Check if we already have profile with such name.
        profiles := []datamodels.Profile{}
        err1 := ctx.Database.Db.Select(&profiles, "SELECT * FROM urt_profiles")
        if err1 != nil {
            fmt.Println(err1.Error())
        }

        var found bool = false
        for p := range profiles {
            if profiles[p].Name == profile.Name {
                found = true
            }
        }

        if found {
            mbox_string := "Game profile with same name already exist.\nRename profile for saving."
            m := gtk.NewMessageDialog(op.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
            m.Response(func() {
                m.Destroy()
            })
            m.Run()
        } else {
            if op.update {
                ctx.Database.Db.NamedExec("UPDATE urt_profiles SET name=:name, version=:version, binary=:binary, second_x_session=:second_x_session, additional_parameters=:additional_parameters WHERE name='" + op.old_profile.Name + "'", &profile)
            } else {
                ctx.Database.Db.NamedExec("INSERT INTO urt_profiles (name, version, binary, second_x_session, additional_parameters) VALUES (:name, :version, :binary, :second_x_session, :additional_parameters)", &profile)
            }
        }
    }
    op.loadProfiles()
    op.window.Destroy()
}
