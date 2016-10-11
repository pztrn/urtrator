package ui

import (
    // stdlib
    "errors"
    "fmt"
    "strings"

    // Local
    "github.com/pztrn/urtrator/datamodels"

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
    if strings.Contains(current_tab, "Favorites") {
        sel = m.fav_servers.GetSelection()
        model = m.fav_servers.GetModel()
    }
    iter := new(gtk.TreeIter)
    _ = sel.GetSelected(iter)

    // Getting server address.
    var srv_addr string
    srv_address_gval := glib.ValueFromNative(srv_addr)
    if strings.Contains(current_tab, "Servers") {
        model.GetValue(iter, m.column_pos["Servers"]["IP"], srv_address_gval)
    } else if strings.Contains(current_tab, "Favorites") {
        model.GetValue(iter, m.column_pos["Favorites"]["IP"], srv_address_gval)
    }
    srv_address = srv_address_gval.GetString()
    server_profile := ctx.Cache.Servers[srv_address].Server

    // Check for proper server name. If length == 0: server is offline,
    // we should show notification to user.
    if len(server_profile.Name) == 0 {
        var will_continue bool = false
        mbox_string := "Selected server is offline.\n\nWould you still want to launch Urban Terror?\nIt will just launch a game, without connecting to\nany server."
        m := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_YES_NO, mbox_string)
        m.Connect("response", func(resp *glib.CallbackContext) {
            if resp.Args(0) == 4294967287 {
                will_continue = false
            } else {
                will_continue = true
            }
        })
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
        if !will_continue {
            return errors.New("User declined to connect to offline server")
        }
    }

    // Getting selected profile's name.
    profile_name := m.profiles.GetActiveText()
    user_profile := &datamodels.Profile{}
    if strings.Contains(current_tab, "Servers") {
        // Checking profile name length. If 0 - then stop executing :)
        // This check only relevant to "Servers" tab, favorite servers
        // have profiles defined (see next).
        if len(profile_name) == 0 {
            mbox_string := "Invalid game profile selected.\n\nPlease, select profile and retry."
            m := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
            m.Response(func() {
                m.Destroy()
            })
            m.Run()
            return errors.New("User didn't select valid profile.")
        }
        user_profile = ctx.Cache.Profiles[profile_name].Profile
    } else if strings.Contains(current_tab, "Favorites") {
        // For favorite servers profile specified in favorite server
        // information have higher priority, so we just override it :)
        user_profile_cached, ok := ctx.Cache.Profiles[server_profile.ProfileToUse]
        if !ok {
            mbox_string := "Invalid game profile specified for favorite server.\n\nPlease, edit your favorite server, select valid profile and retry."
            m := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
            m.Response(func() {
                m.Destroy()
            })
            m.Run()
            return errors.New("User didn't select valid profile.")
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
        mbox_string := "Selected server is offline.\n\nWould you still want to launch Urban Terror?\nIt will just launch a game, without connecting to\nany server."
        m := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_YES_NO, mbox_string)
        m.Connect("response", func(resp *glib.CallbackContext) {
            if resp.Args(0) == 4294967287 {
                will_continue = false
            } else {
                will_continue = true
            }
        })
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
        if !will_continue {
            return errors.New("User declined to connect to offline server")
        }
    }

    // Check if server is applicable for selected profile.
    if server_profile.Version != user_profile.Version {
        mbox_string := "Invalid game profile selected.\n\nSelected profile have different game version than server.\nPlease, select valid profile and retry."
        m := gtk.NewMessageDialog(m.window, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
        return errors.New("User didn't select valid profile, mismatch with server's version.")
    }

    server_password := password
    if len(server_password) == 0 {
        server_password = server_profile.Password
    }

    // Hey, we're ok here! :) Launch Urban Terror!
    // Clear server name from "<markup></markup>" things.
    srv_name_for_label := server_profile.Name
    if strings.Contains(server_profile.Name, "markup") {
        srv_name_for_label = string([]byte(server_profile.Name)[8:len(server_profile.Name)-9])
    }
    // Show great coloured label.
    ctx.Eventer.LaunchEvent("setToolbarLabelText", map[string]string{"text": "<markup><span foreground=\"red\" font_weight=\"bold\">Urban Terror is launched with profile </span><span foreground=\"blue\">" + user_profile.Name + "</span> <span foreground=\"red\" font_weight=\"bold\">and connected to </span><span foreground=\"orange\" font_weight=\"bold\">" + srv_name_for_label + "</span></markup>"})
    m.launch_button.SetSensitive(false)
    // ToDo: handling server passwords.
    ctx.Launcher.Launch(server_profile, user_profile, server_password, []string{"+name", nickname_to_use}, m.unlockInterface)

    return nil
}
