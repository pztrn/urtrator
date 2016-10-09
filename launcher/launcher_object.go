// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package launcher

import (
    // stdlib
    "errors"
    "fmt"
    "os"
    "os/exec"
    "runtime"
    "strconv"
    "strings"

    // local
    "github.com/pztrn/urtrator/datamodels"

    // Github
    "github.com/mattn/go-gtk/gtk"
)

type Launcher struct {
    // Flags.
    // Is Urban Terror launched ATM?
    launched bool
}

func (l *Launcher) CheckForLaunchedUrbanTerror() error {
    if l.launched {
        mbox_string := "Game is launched.\n\nCannot quit, because game is launched.\nQuit Urban Terror to exit URTrator!"
        m := gtk.NewMessageDialog(nil, gtk.DIALOG_MODAL, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, mbox_string)
        m.Response(func() {
            m.Destroy()
        })
        m.Run()
        return errors.New("User didn't select valid profile, mismatch with server's version.")
    }

    return nil
}

func (l *Launcher) findFreeDisplay() string {
    current_display_raw := os.Getenv("DISPLAY")
    current_display, _ := strconv.Atoi(strings.Split(current_display_raw, ":")[1])
    current_display += 1
    return strconv.Itoa(current_display)
}

func (l *Launcher) Initialize() {
    fmt.Println("Initializing game launcher...")
}

func (l *Launcher) Launch(server_profile *datamodels.Server, user_profile *datamodels.Profile, password string, additional_parameters []string, callback func()) {
    // ToDo: only one instance of Urban Terror should be launched, so button
    // should be disabled.
    fmt.Println("Launching Urban Terror...")

    done := make(chan bool, 1)

    // Create launch string.
    var launch_bin string = ""
    launch_bin, err := exec.LookPath(user_profile.Binary)
    if err != nil {
        fmt.Println(err.Error())
    }

    server_address := server_profile.Ip + ":" + server_profile.Port

    var launch_params []string
    if len(server_address) > 0 {
        launch_params = append(launch_params, "+connect", server_address)
    }
    if len(password) > 0 {
        launch_params = append(launch_params, "+password", password)
    }
    if len(user_profile.Additional_params) > 0 {
        additional_params := strings.Split(user_profile.Additional_params, " ")
        launch_params = append(launch_params, additional_params...)
    }
    if len(additional_parameters) > 0 {
        for i := range additional_parameters {
            launch_params = append(launch_params, additional_parameters[i])
        }
    }
    if runtime.GOOS == "linux" && user_profile.Second_x_session == "1" {
        launch_params = append([]string{launch_bin}, launch_params...)
        display := l.findFreeDisplay()
        launch_bin, err = exec.LookPath("xinit")
        if err != nil {
            fmt.Println(err.Error())
        }
        launch_params = append(launch_params, "--", ":" + display)
    }
    if runtime.GOOS == "darwin" {
        // On macOS we should not start binary, but application bundle.
        // So we will obtain app bundle path.
        bundle_path := strings.Split(launch_bin, "/Contents")[0]
        // and create special launch string, which involves open.
        launch_bin = "/usr/bin/open"
        launch_params = append([]string{launch_bin, "-W", bundle_path, "--args"}, launch_params...)
    }
    fmt.Println(launch_bin, launch_params)
    go func() {
        go func() {
            cmd := exec.Command(launch_bin, launch_params...)
            out, err1 := cmd.Output()
            if err1 != nil {
                fmt.Println("Launch error: " + err1.Error())
            }
            fmt.Println(string(out))
            done <- true
        }()

        select {
        case <- done:
            callback()
        }
    }()
}
