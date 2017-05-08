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
    // Appearance tab.
    // Language to use.
    language_combo *gtk.ComboBoxText
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
    o.saveAppearance()

    // Temporary disable all these modals on Linux.
    // See https://github.com/mattn/go-gtk/issues/289.
    if runtime.GOOS != "linux" {
        mbox_string := ctx.Translator.Translate("Some options require application restart to be applied.", nil)
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

    appearance_table := gtk.NewTable(1, 2, false)

    language_selection_tooltip := ctx.Translator.Translate("Language which URTrator will use.\n\nChanging this requires URTrator restart!", nil)

    language_selection_label := gtk.NewLabel(ctx.Translator.Translate("Language:", nil))
    language_selection_label.SetAlignment(0, 0)
    language_selection_label.SetTooltipText(language_selection_tooltip)
    appearance_table.Attach(language_selection_label, 0, 1, 0, 1, gtk.FILL, gtk.SHRINK, 5, 5)

    o.language_combo = gtk.NewComboBoxText()
    o.language_combo.SetTooltipText(language_selection_tooltip)
    // Get all available languages and fill combobox.
    lang_idx := 0
    var lang_active int = 0
    for lang, _ := range ctx.Translator.AcceptedLanguages {
        o.language_combo.AppendText(lang)
        if ctx.Translator.AcceptedLanguages[lang] == ctx.Cfg.Cfg["/general/language"] {
            lang_active = lang_idx
        }
        lang_idx += 1
    }
    o.language_combo.SetActive(lang_active)

    appearance_table.Attach(o.language_combo, 1, 2, 0, 1, gtk.FILL | gtk.EXPAND, gtk.FILL, 5, 5)

    appearance_vbox.PackStart(appearance_table, false, true, 0)
    o.tab_widget.AppendPage(appearance_vbox, gtk.NewLabel(ctx.Translator.Translate("Appearance", nil)))
}

func (o *OptionsDialog) initializeGeneralTab() {
    general_vbox := gtk.NewVBox(false, 0)

    general_table := gtk.NewTable(2, 2, false)

    // Tray icon checkbox.
    show_tray_icon_label := gtk.NewLabel(ctx.Translator.Translate("Show icon in tray", nil))
    show_tray_icon_label.SetAlignment(0, 0)
    general_table.Attach(show_tray_icon_label, 0, 1, 0, 1, gtk.FILL, gtk.SHRINK, 5, 5)

    o.show_tray_icon = gtk.NewCheckButtonWithLabel("")
    o.show_tray_icon.SetTooltipText(ctx.Translator.Translate("Show icon in tray", nil))
    general_table.Attach(o.show_tray_icon, 1, 2, 0, 1, gtk.FILL | gtk.EXPAND, gtk.FILL, 5, 5)

    // Autoupdate checkbox.
    autoupdate_tooltip := ctx.Translator.Translate("Should URTrator check for updates and update itself? Not working now.", nil)
    autoupdate_label := gtk.NewLabel(ctx.Translator.Translate("Automatically update URTrator?", nil))
    autoupdate_label.SetTooltipText(autoupdate_tooltip)
    autoupdate_label.SetAlignment(0, 0)
    general_table.Attach(autoupdate_label, 0, 1, 1, 2, gtk.FILL, gtk.SHRINK, 5, 5)

    o.autoupdate = gtk.NewCheckButtonWithLabel("")
    o.autoupdate.SetTooltipText(autoupdate_tooltip)
    general_table.Attach(o.autoupdate, 1, 2, 1, 2, gtk.FILL | gtk.EXPAND, gtk.FILL, 5, 5)

    // Vertical separator.
    sep := gtk.NewVBox(false, 0)

    general_vbox.PackStart(general_table, false, true, 0)
    general_vbox.PackStart(sep, false, true, 0)

    o.tab_widget.AppendPage(general_vbox, gtk.NewLabel(ctx.Translator.Translate("General", nil)))
}

func (o *OptionsDialog) initializeServersOptionsTab() {
    servers_options_vbox := gtk.NewVBox(false, 0)

    servers_updating_table := gtk.NewTable(3, 2, false)
    servers_updating_table.SetRowSpacings(2)

    // Master server address.
    master_server_addr_tooltip := ctx.Translator.Translate("Address of master server. Specify in form: addr:port.", nil)
    master_server_addr_label := gtk.NewLabel(ctx.Translator.Translate("Master server address", nil))
    master_server_addr_label.SetTooltipText(master_server_addr_tooltip)
    master_server_addr_label.SetAlignment(0, 0)
    servers_updating_table.Attach(master_server_addr_label, 0, 1, 0, 1, gtk.FILL, gtk.SHRINK, 5, 5)

    o.master_server_addr = gtk.NewEntry()
    o.master_server_addr.SetTooltipText(master_server_addr_tooltip)
    servers_updating_table.Attach(o.master_server_addr, 1, 2, 0, 1, gtk.FILL, gtk.FILL, 5, 5)

    // Servers autoupdate checkbox.
    servers_autoupdate_cb_tooptip := ctx.Translator.Translate("Should servers be automatically updated?", nil)
    servers_autoupdate_cb_label := gtk.NewLabel(ctx.Translator.Translate("Servers autoupdate", nil))
    servers_autoupdate_cb_label.SetTooltipText(servers_autoupdate_cb_tooptip)
    servers_autoupdate_cb_label.SetAlignment(0, 0)
    servers_updating_table.Attach(servers_autoupdate_cb_label, 0, 1 ,1, 2, gtk.FILL, gtk.SHRINK, 5, 5)

    o.servers_autoupdate = gtk.NewCheckButtonWithLabel("")
    o.servers_autoupdate.SetTooltipText(servers_autoupdate_cb_tooptip)
    servers_updating_table.Attach(o.servers_autoupdate, 1, 2, 1, 2, gtk.EXPAND | gtk.FILL, gtk.FILL, 5, 5)

    // Servers update timeout.
    servers_autoupdate_timeout_tooltip := ctx.Translator.Translate("Timeout which will trigger servers information update, in minutes.", nil)
    servers_autoupdate_label := gtk.NewLabel(ctx.Translator.Translate("Servers update timeout (minutes)", nil))
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

    o.tab_widget.AppendPage(servers_options_vbox, gtk.NewLabel(ctx.Translator.Translate("Servers updating", nil)))
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

    cancel_button := gtk.NewButtonWithLabel(ctx.Translator.Translate("Cancel", nil))
    cancel_button.Clicked(o.closeOptionsDialogByCancel)

    ok_button := gtk.NewButtonWithLabel(ctx.Translator.Translate("OK", nil))
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
    o.profiles_list.SetTooltipText(ctx.Translator.Translate("All available profiles", nil))
    urt_hbox.Add(o.profiles_list)
    o.profiles_list.SetModel(o.profiles_list_store)
    o.profiles_list.AppendColumn(gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Profile name", nil), gtk.NewCellRendererText(), "text", 0))
    o.profiles_list.AppendColumn(gtk.NewTreeViewColumnWithAttributes(ctx.Translator.Translate("Urban Terror version", nil), gtk.NewCellRendererText(), "text", 1))

    // Profiles list buttons.
    urt_profiles_buttons_vbox := gtk.NewVBox(false, 0)

    button_add := gtk.NewButtonWithLabel(ctx.Translator.Translate("Add", nil))
    button_add.SetTooltipText(ctx.Translator.Translate("Add new profile", nil))
    button_add.Clicked(o.addProfile)
    urt_profiles_buttons_vbox.PackStart(button_add, false, true, 0)

    button_edit := gtk.NewButtonWithLabel(ctx.Translator.Translate("Edit", nil))
    button_edit.SetTooltipText(ctx.Translator.Translate("Edit selected profile. Do nothing if no profile was selected.", nil))
    button_edit.Clicked(o.editProfile)
    urt_profiles_buttons_vbox.PackStart(button_edit, false, true, 5)

    // Spacer for profiles list buttons.
    sep := gtk.NewVBox(false, 0)
    urt_profiles_buttons_vbox.PackStart(sep, true, true, 5)

    button_delete := gtk.NewButtonWithLabel(ctx.Translator.Translate("Delete", nil))
    button_delete.SetTooltipText(ctx.Translator.Translate("Delete selected profile. Do nothing if no profile was selected.", nil))
    button_delete.Clicked(o.deleteProfile)
    urt_profiles_buttons_vbox.PackStart(button_delete, false, true, 0)

    urt_hbox.Add(urt_profiles_buttons_vbox)

    o.tab_widget.AppendPage(urt_hbox, gtk.NewLabel(ctx.Translator.Translate("Urban Terror", nil)))

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

func (o *OptionsDialog) saveAppearance() {
    ctx.Cfg.Cfg["/general/language"] = ctx.Translator.AcceptedLanguages[o.language_combo.GetActiveText()]
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
    o.window.SetTitle(ctx.Translator.Translate("URTrator - Options", nil))
    o.window.Connect("destroy", o.closeOptionsDialogWithDiscard)
    o.window.SetModal(true)
    o.window.SetSizeRequest(750, 600)
    o.window.SetPosition(gtk.WIN_POS_CENTER)
    o.window.SetIcon(logo)

    o.vbox = gtk.NewVBox(false, 0)

    o.initializeTabs()
    o.fill()

    o.window.Add(o.vbox)

    ctx.Eventer.LaunchEvent("loadProfilesIntoOptionsWindow", map[string]string{})

    o.window.ShowAll()
}
