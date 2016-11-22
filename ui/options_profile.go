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

    // Others.
    // Old profile, needed for proper update.
    old_profile *datamodels.Profile
}

func (op *OptionsProfile) browseForBinary() {
    op.f = gtk.NewFileChooserDialog("URTrator - Select Urban Terror binary", op.window, gtk.FILE_CHOOSER_ACTION_OPEN, gtk.STOCK_OK, gtk.RESPONSE_ACCEPT)
    op.f.Response(op.browseForBinaryHelper)
    op.f.Run()
}

func (op *OptionsProfile) browseForBinaryHelper() {
    filename := op.f.GetFilename()
    op.binary_path.SetText(filename)
    op.f.Destroy()
    fmt.Println(filename)

    // Check for valid filename.
    // ToDo: add more OSes.
    if runtime.GOOS == "linux" {
        // Filename should end with approriate arch.
        if runtime.GOARCH == "amd64" {
            if len(filename) > 0 && strings.Split(filename, ".")[1] != "x86_64" && strings.Split(filename, ".")[0] != "Quake3-UrT" {
                fmt.Println("Invalid binary selected!")
                // Temporary disable all these modals on Linux.
                // See https://github.com/mattn/go-gtk/issues/289.
                if runtime.GOOS != "linux" {
                    mbox_string := "Invalid binary selected!\nAccording to your OS, it should be Quake3-UrT.x86_64."
                    m := gtk.NewMessageDialog(op.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
                    m.Response(func() {
                        m.Destroy()
                    })
                    m.Run()
                } else {
                    //
                }
                op.binary_path.SetText("")
            }
        }
    } else if runtime.GOOS == "darwin" {
        // Official application: Quake3-UrT.app. Split by it and get second
        // part of string.
        if strings.Contains(filename, "Quake3-UrT.app") {
            filename = strings.Split(filename, "Quake3-UrT.app")[1]
            if len(filename) > 0 && !strings.Contains(strings.Split(filename, ".")[1], "x86_64") && !strings.Contains(strings.Split(filename, ".")[0], "Quake3-UrT") {
                fmt.Println("Invalid binary selected!")
                // Temporary disable all these modals on Linux.
                // See https://github.com/mattn/go-gtk/issues/289.
                if runtime.GOOS != "linux" {
                    mbox_string := "Invalid binary selected!\nAccording to your OS, it should be Quake3-UrT.app/Contents/MacOS/Quake3-UrT.x86_64."
                    m := gtk.NewMessageDialog(op.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
                    m.Response(func() {
                        m.Destroy()
                    })
                    m.Run()
                } else {
                    //
                }
                op.binary_path.SetText("")
            }
        } else {
            // Temporary disable all these modals on Linux.
            // See https://github.com/mattn/go-gtk/issues/289.
            if runtime.GOOS != "linux" {
                mbox_string := "Invalid binary selected!\nAccording to your OS, it should be Quake3-UrT.app/Contents/MacOS/Quake3-UrT.x86_64.\n\nNote, that currently URTrator supports only official binary."
                m := gtk.NewMessageDialog(op.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
                m.Response(func() {
                    m.Destroy()
                })
                m.Run()
            } else {
                //
            }
        }
    }
}

func (op *OptionsProfile) closeByCancel() {
    op.window.Destroy()
}

func (op *OptionsProfile) closeWithDiscard() {
}

func (op *OptionsProfile) Initialize(update bool) {
    if update {
        op.update = true
    }

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
    profile_name_tooltip := "This how you will see profile on profiles lists."
    pn_hbox := gtk.NewHBox(false, 0)
    pn_label := gtk.NewLabel("Profile name:")
    pn_label.SetTooltipText(profile_name_tooltip)
    profile_name_sep := gtk.NewHSeparator()
    profile_name_sep.SetTooltipText(profile_name_tooltip)
    op.profile_name = gtk.NewEntry()
    op.profile_name.SetTooltipText(profile_name_tooltip)
    pn_hbox.PackStart(pn_label, false, true, 5)
    pn_hbox.PackStart(profile_name_sep, true, true, 5)
    pn_hbox.PackStart(op.profile_name, true, true, 5)
    op.vbox.PackStart(pn_hbox, false, true, 5)

    // Urban Terror version.
    urt_version_tooltip := "Urban Terror version for which this profile applies."
    urt_version_hbox := gtk.NewHBox(false, 0)
    urt_version_label := gtk.NewLabel("Urban Terror version:")
    urt_version_label.SetTooltipText(urt_version_tooltip)
    urt_version_sep := gtk.NewHSeparator()
    urt_version_sep.SetTooltipText(urt_version_tooltip)
    op.urt_version_combo = gtk.NewComboBoxText()
    op.urt_version_combo.SetTooltipText(urt_version_tooltip)
    op.urt_version_combo.AppendText("4.2.023")
    op.urt_version_combo.AppendText("4.3.0")
    op.urt_version_combo.AppendText("4.3.1")
    op.urt_version_combo.SetActive(2)
    urt_version_hbox.PackStart(urt_version_label, false, true, 5)
    urt_version_hbox.PackStart(urt_version_sep, true, true, 5)
    urt_version_hbox.PackStart(op.urt_version_combo, true, true, 5)
    op.vbox.PackStart(urt_version_hbox, false, true, 5)

    // Urban Terror binary path.
    select_binary_tooltip := "Urban Terror binary. Some checks will be executed, so make sure you have selected right binary:\n\nQuake3-UrT.i386 for linux-x86\nQuake3-UrT.x86_64 for linux-amd64\nQuake3-UrT.app for macOS"
    binpath_hbox := gtk.NewHBox(false, 0)
    binpath_label := gtk.NewLabel("Urban Terror binary:")
    binpath_label.SetTooltipText(select_binary_tooltip)
    binpath_sep := gtk.NewHSeparator()
    binpath_sep.SetTooltipText(select_binary_tooltip)
    op.binary_path = gtk.NewEntry()
    op.binary_path.SetTooltipText(select_binary_tooltip)
    button_select_binary := gtk.NewButtonWithLabel("Browse")
    button_select_binary.SetTooltipText(select_binary_tooltip)
    button_select_binary.Clicked(op.browseForBinary)
    binpath_hbox.PackStart(binpath_label, false, true, 5)
    binpath_hbox.PackStart(binpath_sep, true, true, 5)
    binpath_hbox.PackStart(op.binary_path, true, true, 5)
    binpath_hbox.PackStart(button_select_binary, false, true, 5)
    op.vbox.PackStart(binpath_hbox, false, true, 5)

    // Should we use additional X session?
    another_x_tooltip := "If this is checked, Urban Terror will be launched in another X session.\n\nThis could help if you're experiencing visual lag, glitches and FPS drops under compositing WMs, like Mutter and KWin."
    op.another_x_session = gtk.NewCheckButtonWithLabel("Start Urban Terror in another X session?")
    op.another_x_session.SetTooltipText(another_x_tooltip)
    op.vbox.PackStart(op.another_x_session, false, true, 5)
    // macOS can't do that :).
    if runtime.GOOS != "linux" {
        op.another_x_session.SetSensitive(false)
    }

    // Additional game parameters.
    params_tooltip := "Additional parameters that will be passed to Urban Terror executable."
    params_hbox := gtk.NewHBox(false, 0)
    params_label := gtk.NewLabel("Additional parameters:")
    params_label.SetTooltipText(params_tooltip)
    params_sep := gtk.NewHSeparator()
    params_sep.SetTooltipText(params_tooltip)
    op.additional_parameters = gtk.NewEntry()
    op.additional_parameters.SetTooltipText(params_tooltip)
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
    cancel_button.SetTooltipText("Close without saving")
    cancel_button.Clicked(op.closeByCancel)
    buttons_box.PackStart(cancel_button, false, true, 5)

    buttons_box.PackStart(buttons_sep, true, true, 5)

    add_button := gtk.NewButton()
    if op.update {
        add_button.SetLabel("Update")
        add_button.SetTooltipText("Update profile")
    } else {
        add_button.SetLabel("Add")
        add_button.SetTooltipText("Add profile")
    }
    add_button.Clicked(op.saveProfile)
    buttons_box.PackStart(add_button, false, true, 5)

    op.vbox.PackStart(buttons_box, false, true, 5)

    op.window.Add(op.vbox)
    op.window.ShowAll()
}

func (op *OptionsProfile) InitializeUpdate(profile_name string) {
    fmt.Println("Updating profile '" + profile_name + "'")
    op.Initialize(true)

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

    if profile[0].Version == "4.3.0" {
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
    _, err := os.Stat(op.binary_path.GetText())
    if err != nil {
        mbox_string := "Invalid path to binary!\n\nError was:\n" + err.Error() + "\n\nCheck binary path and try again."
        m := gtk.NewMessageDialog(op.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
    } else {
        // ToDo: executable flag checking.
        //fmt.Println(filestat.Mode())
        profile_name := op.profile_name.GetText()

        _, ok := ctx.Cache.Profiles[profile_name]
        if ok && !op.update {
            mbox_string := "Game profile with same name already exist.\nRename profile for saving."
            m := gtk.NewMessageDialog(op.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
            m.Response(func() {
                m.Destroy()
            })
            m.Run()
        } else {
            ctx.Cache.CreateProfile(profile_name)
            ctx.Cache.Profiles[profile_name].Profile.Name = profile_name
            ctx.Cache.Profiles[profile_name].Profile.Version = op.urt_version_combo.GetActiveText()
            ctx.Cache.Profiles[profile_name].Profile.Binary = op.binary_path.GetText()
            ctx.Cache.Profiles[profile_name].Profile.Additional_params = op.additional_parameters.GetText()

            if op.another_x_session.GetActive() {
                ctx.Cache.Profiles[profile_name].Profile.Second_x_session = "1"
            } else {
                ctx.Cache.Profiles[profile_name].Profile.Second_x_session = "0"
            }
        }
    }
    ctx.Eventer.LaunchEvent("loadProfilesIntoOptionsWindow", map[string]string{})
    ctx.Eventer.LaunchEvent("loadProfilesIntoMainWindow", map[string]string{})
    op.window.Destroy()
}
