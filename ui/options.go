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
    "github.com/pztrn/urtrator/datamodels"

    // Other
    "github.com/mattn/go-gtk/gtk"
    "github.com/mattn/go-gtk/glib"
)

type OptionsDialog struct {
    // Window.
    window *gtk.Window
    // Options main VBox.
    vbox *gtk.VBox
    // Tabs widget.
    tab_widget *gtk.Notebook

    // Widgets.
    // General tab.
    // Show tray icon checkbutton.
    show_tray_icon *gtk.CheckButton
    // Enable autoupdate checkbutton.
    autoupdate *gtk.CheckButton
    // Urban Terror tab.
    // Profiles list.
    profiles_list *gtk.TreeView

    // Data stores.
    // Urban Terror profiles list.
    profiles_list_store *gtk.ListStore
}

func (o *OptionsDialog) addProfile() {
    fmt.Println("Adding profile...")

    op := OptionsProfile{}
    op.Initialize(false, o.loadProfiles)
    ctx.Eventer.LaunchEvent("loadProfiles")
}

func (o *OptionsDialog) closeOptionsDialogByCancel() {
    o.window.Destroy()
}

func (o *OptionsDialog) closeOptionsDialogWithDiscard() {
}

func (o *OptionsDialog) closeOptionsDialogWithSaving() {
    fmt.Println("Saving changes to options...")

    ctx.Eventer.LaunchEvent("loadProfiles")

    o.window.Destroy()
}

func (o *OptionsDialog) deleteProfile() {
    // Oh... dat... GTK...
    sel := o.profiles_list.GetSelection()
    model := o.profiles_list.GetModel()
    iter := new(gtk.TreeIter)
    _ = sel.GetSelected(iter)
    var p string
    gval := glib.ValueFromNative(p)
    model.GetValue(iter, 0, gval)
    profile_name := gval.GetString()

    if len(profile_name) > 0 {
        fmt.Println("Deleting profile '" + profile_name + "'")

        profile := datamodels.Profile{}
        profile.Name = profile_name
        ctx.Database.Db.NamedExec("DELETE FROM urt_profiles WHERE name=:name", &profile)
        ctx.Eventer.LaunchEvent("loadProfiles")
    }
}

func (o *OptionsDialog) editProfile() {
    // Oh... dat... GTK...
    sel := o.profiles_list.GetSelection()
    model := o.profiles_list.GetModel()
    iter := new(gtk.TreeIter)
    _ = sel.GetSelected(iter)
    var p string
    gval := glib.ValueFromNative(p)
    model.GetValue(iter, 0, gval)
    profile_name := gval.GetString()

    if len(profile_name) > 0 {
        op := OptionsProfile{}
        op.InitializeUpdate(profile_name, o.loadProfiles)
        ctx.Eventer.LaunchEvent("loadProfiles")
    }
}

func (o *OptionsDialog) initializeGeneralTab() {
    general_vbox := gtk.NewVBox(false, 0)

    // Tray icon checkbox.
    o.show_tray_icon = gtk.NewCheckButtonWithLabel("Show tray icon?")
    o.show_tray_icon.SetTooltipText("Show icon in tray")
    general_vbox.PackStart(o.show_tray_icon, false, true, 5)

    // Autoupdate checkbox.
    o.autoupdate = gtk.NewCheckButtonWithLabel("Automatically update URTrator?")
    o.autoupdate.SetTooltipText("Should URTrator check for updates and update itself? Not working now.")
    general_vbox.PackStart(o.autoupdate, false, true, 5)

    o.tab_widget.AppendPage(general_vbox, gtk.NewLabel("General"))
}

func (o *OptionsDialog) initializeStorages() {
    // Structure:
    // Name|Version|Second X session
    o.profiles_list_store = gtk.NewListStore(glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_BOOL)
}

func (o *OptionsDialog) initializeTabs() {
    o.initializeStorages()
    o.tab_widget = gtk.NewNotebook()

    o.initializeGeneralTab()
    o.initializeUrtTab()

    // Buttons for saving and discarding changes.
    buttons_hbox := gtk.NewHBox(false, 0)
    sep := gtk.NewHSeparator()

    cancel_button := gtk.NewButtonWithLabel("Cancel")
    cancel_button.Clicked(o.closeOptionsDialogByCancel)

    ok_button := gtk.NewButtonWithLabel("OK")
    ok_button.Clicked(o.closeOptionsDialogWithSaving)

    buttons_hbox.PackStart(sep, true, true, 5)
    buttons_hbox.PackStart(cancel_button, false, true, 5)
    buttons_hbox.PackStart(ok_button, false, true, 5)

    o.vbox.PackStart(o.tab_widget, true, true, 5)
    o.vbox.PackStart(buttons_hbox, false, true, 5)
}

func (o *OptionsDialog) initializeUrtTab() {
    urt_hbox := gtk.NewHBox(false, 0)

    // Profiles list.
    o.profiles_list = gtk.NewTreeView()
    o.profiles_list.SetTooltipText("All available profiles")
    urt_hbox.Add(o.profiles_list)
    o.profiles_list.SetModel(o.profiles_list_store)
    o.profiles_list.AppendColumn(gtk.NewTreeViewColumnWithAttributes("Profile name", gtk.NewCellRendererText(), "text", 0))
    o.profiles_list.AppendColumn(gtk.NewTreeViewColumnWithAttributes("Urban Terror version", gtk.NewCellRendererText(), "text", 1))
    crt := gtk.NewCellRendererToggle()
    o.profiles_list.AppendColumn(gtk.NewTreeViewColumnWithAttributes("Second X session", crt, "bool", 2))

    // Profiles list buttons.
    urt_profiles_buttons_vbox := gtk.NewVBox(false, 0)

    button_add := gtk.NewButtonWithLabel("Add")
    button_add.SetTooltipText("Add new profile")
    button_add.Clicked(o.addProfile)
    urt_profiles_buttons_vbox.PackStart(button_add, false, true, 5)

    button_edit := gtk.NewButtonWithLabel("Edit")
    button_edit.SetTooltipText("Edit selected profile. Do nothing if no profile was selected.")
    button_edit.Clicked(o.editProfile)
    urt_profiles_buttons_vbox.PackStart(button_edit, false, true, 5)

    // Spacer for profiles list buttons.
    sep := gtk.NewVSeparator()
    urt_profiles_buttons_vbox.PackStart(sep, true, true, 5)

    button_delete := gtk.NewButtonWithLabel("Delete")
    button_delete.SetTooltipText("Delete selected profile. Do nothing if no profile was selected.")
    button_delete.Clicked(o.deleteProfile)
    urt_profiles_buttons_vbox.PackStart(button_delete, false, true, 5)

    urt_hbox.Add(urt_profiles_buttons_vbox)

    o.tab_widget.AppendPage(urt_hbox, gtk.NewLabel("Urban Terror"))

    // Load Profiles.
    o.loadProfiles()
}

func (o *OptionsDialog) loadProfiles() {
    fmt.Println("Loading profiles...")
    o.profiles_list_store.Clear()
    // Get profiles from database.
    profiles := []datamodels.Profile{}
    err := ctx.Database.Db.Select(&profiles, "SELECT * FROM urt_profiles")
    if err != nil {
        fmt.Println(err.Error())
    }

    for p := range profiles {
        var iter gtk.TreeIter
        o.profiles_list_store.Append(&iter)
        o.profiles_list_store.Set(&iter, 0, profiles[p].Name)
        o.profiles_list_store.Set(&iter, 1, profiles[p].Version)
    }
}

func (o *OptionsDialog) ShowOptionsDialog() {
    o.window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
    o.window.SetTitle("URTrator - Options")
    o.window.Connect("destroy", o.closeOptionsDialogWithDiscard)
    o.window.SetModal(true)
    o.window.SetSizeRequest(550, 400)
    o.window.SetPosition(gtk.WIN_POS_CENTER)
    o.window.SetIcon(logo)

    o.vbox = gtk.NewVBox(false, 0)

    o.initializeTabs()

    o.window.Add(o.vbox)
    o.window.ShowAll()
}
