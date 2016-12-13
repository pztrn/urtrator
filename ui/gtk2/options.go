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
    "runtime"

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
    // Servers updating tab.
    // Master server address.
    master_server_addr *gtk.Entry
    // Servers autoupdate.
    servers_autoupdate *gtk.CheckButton
    // Timeout for servers autoupdating.
    servers_autoupdate_timeout *gtk.Entry

    // Data stores.
    // Urban Terror profiles list.
    profiles_list_store *gtk.ListStore
}

func (o *OptionsDialog) addProfile() {
    fmt.Println("Adding profile...")

    op := OptionsProfile{}
    op.Initialize(false)
    ctx.Eventer.LaunchEvent("flushProfiles", map[string]string{})
    ctx.Eventer.LaunchEvent("loadProfilesIntoOptionsWindow", map[string]string{})
    ctx.Eventer.LaunchEvent("loadProfilesIntoMainWindow", map[string]string{})
}

func (o *OptionsDialog) closeOptionsDialogByCancel() {
    o.window.Destroy()
}

func (o *OptionsDialog) closeOptionsDialogWithDiscard() {
}

func (o *OptionsDialog) closeOptionsDialogWithSaving() {
    fmt.Println("Saving changes to options...")

    o.saveGeneral()

    // Temporary disable all these modals on Linux.
    // See https://github.com/mattn/go-gtk/issues/289.
    if runtime.GOOS != "linux" {
        mbox_string := "Some options require application restart to be applied."
        m := gtk.NewMessageDialog(o.window, gtk.DIALOG_MODAL, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, mbox_string)
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
    }

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
        ctx.Eventer.LaunchEvent("deleteProfile", map[string]string{"profile_name": profile_name})
        ctx.Eventer.LaunchEvent("flushProfiles", map[string]string{})
        ctx.Eventer.LaunchEvent("loadProfilesIntoMainWindow", map[string]string{})
        ctx.Eventer.LaunchEvent("loadProfilesIntoOptionsWindow", map[string]string{})
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
        op.InitializeUpdate(profile_name)
        ctx.Eventer.LaunchEvent("flushProfiles", map[string]string{})
        ctx.Eventer.LaunchEvent("loadProfilesIntoMainWindow", map[string]string{})
        ctx.Eventer.LaunchEvent("loadProfilesIntoOptionsWindow", map[string]string{})
    }
}

func (o *OptionsDialog) fill() {
    if ctx.Cfg.Cfg["/general/show_tray_icon"] == "1" {
        o.show_tray_icon.SetActive(true)
    }
    if ctx.Cfg.Cfg["/general/urtrator_autoupdate"] == "1" {
        o.autoupdate.SetActive(true)
    }

    // Servers updating tab.
    master_server_addr, ok := ctx.Cfg.Cfg["/servers_updating/master_server"]
    if !ok {
        o.master_server_addr.SetText("master.urbanterror.info:27900")
    } else {
        o.master_server_addr.SetText(master_server_addr)
    }

    servers_autoupdate, ok1 := ctx.Cfg.Cfg["/servers_updating/servers_autoupdate"]
    if ok1 {
        if servers_autoupdate == "1" {
            o.servers_autoupdate.SetActive(true)
        }
    }

    servers_update_timeout, ok2 := ctx.Cfg.Cfg["/servers_updating/servers_autoupdate_timeout"]
    if !ok2 {
        o.servers_autoupdate_timeout.SetText("10")
    } else {
        o.servers_autoupdate_timeout.SetText(servers_update_timeout)
    }

}

// Appearance tab initialization.
func (o *OptionsDialog) initializeAppearanceTab() {
    appearance_vbox := gtk.NewVBox(false, 0)

    l := gtk.NewLabel("There will be some appearance configuration options soon.")
    appearance_vbox.PackStart(l, false, true, 5)

    o.tab_widget.AppendPage(appearance_vbox, gtk.NewLabel("Appearance"))
}

func (o *OptionsDialog) initializeGeneralTab() {
    general_vbox := gtk.NewVBox(false, 0)

    general_table := gtk.NewTable(2, 2, false)

    // Tray icon checkbox.
    show_tray_icon_label := gtk.NewLabel("Show icon in tray")
    show_tray_icon_label.SetAlignment(0, 0)
    general_table.Attach(show_tray_icon_label, 0, 1, 0, 1, gtk.FILL, gtk.SHRINK, 5, 5)

    o.show_tray_icon = gtk.NewCheckButtonWithLabel("")
    o.show_tray_icon.SetTooltipText("Show icon in tray")
    general_table.Attach(o.show_tray_icon, 1, 2, 0, 1, gtk.FILL, gtk.FILL, 5, 5)

    // Autoupdate checkbox.
    autoupdate_tooltip := "Should URTrator check for updates and update itself? Not working now."
    autoupdate_label := gtk.NewLabel("Automatically update URTrator?")
    autoupdate_label.SetTooltipText(autoupdate_tooltip)
    autoupdate_label.SetAlignment(0, 0)
    general_table.Attach(autoupdate_label, 0, 1, 1, 2, gtk.FILL, gtk.SHRINK, 5, 5)

    o.autoupdate = gtk.NewCheckButtonWithLabel("")
    o.autoupdate.SetTooltipText(autoupdate_tooltip)
    general_table.Attach(o.autoupdate, 1, 2, 1, 2, gtk.FILL, gtk.FILL, 5, 5)

    // Vertical separator.
    sep := gtk.NewVBox(false, 0)

    general_vbox.PackStart(general_table, false, true, 0)
    general_vbox.PackStart(sep, false, true, 0)

    o.tab_widget.AppendPage(general_vbox, gtk.NewLabel("General"))
}

func (o *OptionsDialog) initializeServersOptionsTab() {
    servers_options_vbox := gtk.NewVBox(false, 0)

    servers_updating_table := gtk.NewTable(3, 2, false)
    servers_updating_table.SetRowSpacings(2)

    // Master server address.
    master_server_addr_tooltip := "Address of master server. Specify in form: addr:port."
    master_server_addr_label := gtk.NewLabel("Master server address")
    master_server_addr_label.SetTooltipText(master_server_addr_tooltip)
    master_server_addr_label.SetAlignment(0, 0)
    servers_updating_table.Attach(master_server_addr_label, 0, 1, 0, 1, gtk.FILL, gtk.SHRINK, 5, 5)

    o.master_server_addr = gtk.NewEntry()
    o.master_server_addr.SetTooltipText(master_server_addr_tooltip)
    servers_updating_table.Attach(o.master_server_addr, 1, 2, 0, 1, gtk.FILL, gtk.FILL, 5, 5)

    // Servers autoupdate checkbox.
    servers_autoupdate_cb_tooptip := "Should servers be automatically updated?"
    servers_autoupdate_cb_label := gtk.NewLabel("Servers autoupdate")
    servers_autoupdate_cb_label.SetTooltipText(servers_autoupdate_cb_tooptip)
    servers_autoupdate_cb_label.SetAlignment(0, 0)
    servers_updating_table.Attach(servers_autoupdate_cb_label, 0, 1 ,1, 2, gtk.FILL, gtk.SHRINK, 5, 5)

    o.servers_autoupdate = gtk.NewCheckButtonWithLabel("")
    o.servers_autoupdate.SetTooltipText(servers_autoupdate_cb_tooptip)
    servers_updating_table.Attach(o.servers_autoupdate, 1, 2, 1, 2, gtk.FILL, gtk.FILL, 5, 5)

    // Servers update timeout.
    servers_autoupdate_timeout_tooltip := "Timeout which will trigger servers information update, in minutes."
    servers_autoupdate_label := gtk.NewLabel("Servers update timeout (minutes)")
    servers_autoupdate_label.SetTooltipText(servers_autoupdate_timeout_tooltip)
    servers_autoupdate_label.SetAlignment(0, 0)
    servers_updating_table.Attach(servers_autoupdate_label, 0, 1, 2, 3, gtk.FILL, gtk.SHRINK, 5, 5)

    o.servers_autoupdate_timeout = gtk.NewEntry()
    o.servers_autoupdate_timeout.SetTooltipText(servers_autoupdate_timeout_tooltip)
    servers_updating_table.Attach(o.servers_autoupdate_timeout, 1, 2, 2, 3, gtk.FILL, gtk.FILL, 5, 5)

    // Vertical separator.
    sep := gtk.NewVBox(false, 0)

    servers_options_vbox.PackStart(servers_updating_table, false, true, 0)
    servers_options_vbox.PackStart(sep, true, true, 0)

    o.tab_widget.AppendPage(servers_options_vbox, gtk.NewLabel("Servers updating"))
}

func (o *OptionsDialog) initializeStorages() {
    // Structure:
    // Name|Version|Second X session
    o.profiles_list_store = gtk.NewListStore(glib.G_TYPE_STRING, glib.G_TYPE_STRING, glib.G_TYPE_BOOL)
}

func (o *OptionsDialog) initializeTabs() {
    o.initializeStorages()
    o.tab_widget = gtk.NewNotebook()
    o.tab_widget.SetTabPos(gtk.POS_LEFT)

    o.initializeGeneralTab()
    o.initializeAppearanceTab()
    o.initializeUrtTab()
    o.initializeServersOptionsTab()

    // Buttons for saving and discarding changes.
    buttons_hbox := gtk.NewHBox(false, 0)
    sep := gtk.NewHBox(false, 0)

    cancel_button := gtk.NewButtonWithLabel("Cancel")
    cancel_button.Clicked(o.closeOptionsDialogByCancel)

    ok_button := gtk.NewButtonWithLabel("OK")
    ok_button.Clicked(o.closeOptionsDialogWithSaving)

    buttons_hbox.PackStart(sep, true, true, 5)
    buttons_hbox.PackStart(cancel_button, false, true, 5)
    buttons_hbox.PackStart(ok_button, false, true, 5)

    o.vbox.PackStart(o.tab_widget, true, true, 5)
    o.vbox.PackStart(buttons_hbox, false, true, 5)

    ctx.Eventer.AddEventHandler("loadProfilesIntoOptionsWindow", o.loadProfiles)
}

func (o *OptionsDialog) initializeUrtTab() {
    urt_hbox := gtk.NewHBox(false, 5)

    // Profiles list.
    o.profiles_list = gtk.NewTreeView()
    o.profiles_list.SetTooltipText("All available profiles")
    urt_hbox.Add(o.profiles_list)
    o.profiles_list.SetModel(o.profiles_list_store)
    o.profiles_list.AppendColumn(gtk.NewTreeViewColumnWithAttributes("Profile name", gtk.NewCellRendererText(), "text", 0))
    o.profiles_list.AppendColumn(gtk.NewTreeViewColumnWithAttributes("Urban Terror version", gtk.NewCellRendererText(), "text", 1))

    // Profiles list buttons.
    urt_profiles_buttons_vbox := gtk.NewVBox(false, 0)

    button_add := gtk.NewButtonWithLabel("Add")
    button_add.SetTooltipText("Add new profile")
    button_add.Clicked(o.addProfile)
    urt_profiles_buttons_vbox.PackStart(button_add, false, true, 0)

    button_edit := gtk.NewButtonWithLabel("Edit")
    button_edit.SetTooltipText("Edit selected profile. Do nothing if no profile was selected.")
    button_edit.Clicked(o.editProfile)
    urt_profiles_buttons_vbox.PackStart(button_edit, false, true, 5)

    // Spacer for profiles list buttons.
    sep := gtk.NewVBox(false, 0)
    urt_profiles_buttons_vbox.PackStart(sep, true, true, 5)

    button_delete := gtk.NewButtonWithLabel("Delete")
    button_delete.SetTooltipText("Delete selected profile. Do nothing if no profile was selected.")
    button_delete.Clicked(o.deleteProfile)
    urt_profiles_buttons_vbox.PackStart(button_delete, false, true, 0)

    urt_hbox.Add(urt_profiles_buttons_vbox)

    o.tab_widget.AppendPage(urt_hbox, gtk.NewLabel("Urban Terror"))

    // Load Profiles.
    ctx.Eventer.LaunchEvent("loadProfilesIntoOptionsWindow", map[string]string{})
}

func (o *OptionsDialog) loadProfiles(data map[string]string) {
    fmt.Println("Loading profiles...")
    o.profiles_list_store.Clear()

    for _, p := range ctx.Cache.Profiles {
        var iter gtk.TreeIter
        o.profiles_list_store.Append(&iter)
        o.profiles_list_store.Set(&iter, 0, p.Profile.Name)
        o.profiles_list_store.Set(&iter, 1, p.Profile.Version)
    }
}

func (o *OptionsDialog) saveGeneral() {
    if o.show_tray_icon.GetActive() {
        ctx.Cfg.Cfg["/general/show_tray_icon"] = "1"
    } else {
        ctx.Cfg.Cfg["/general/show_tray_icon"] = "0"
    }

    if o.autoupdate.GetActive() {
        ctx.Cfg.Cfg["/general/urtrator_autoupdate"] = "1"
    } else {
        ctx.Cfg.Cfg["/general/urtrator_autoupdate"] = "0"
    }

    // Servers updating tab.
    master_server_addr := o.master_server_addr.GetText()
    if len(master_server_addr) < 1 {
        ctx.Cfg.Cfg["/servers_updating/master_server"] = "master.urbanterror.info:27900"
    } else {
        ctx.Cfg.Cfg["/servers_updating/master_server"] = master_server_addr
    }

    if o.servers_autoupdate.GetActive() {
        ctx.Cfg.Cfg["/servers_updating/servers_autoupdate"] = "1"
    } else {
        ctx.Cfg.Cfg["/servers_updating/servers_autoupdate"] = "0"
    }

    update_timeout := o.servers_autoupdate_timeout.GetText()
    if len(update_timeout) < 1 {
        ctx.Cfg.Cfg["/servers_updating/servers_autoupdate_timeout"] = "10"
    } else {
        ctx.Cfg.Cfg["/servers_updating/servers_autoupdate_timeout"] = update_timeout
    }

    ctx.Eventer.LaunchEvent("initializeTasksForMainWindow", map[string]string{})
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
    o.fill()

    o.window.Add(o.vbox)

    ctx.Eventer.LaunchEvent("loadProfilesIntoOptionsWindow", map[string]string{})

    o.window.ShowAll()
}
