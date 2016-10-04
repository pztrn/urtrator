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
    "fmt"
    "os"
    "os/exec"
    "strconv"
    "strings"

    // local
    "github.com/pztrn/urtrator/datamodels"
)

type Launcher struct {}

func (l *Launcher) findFreeDisplay() string {
    current_display_raw := os.Getenv("DISPLAY")
    current_display, _ := strconv.Atoi(strings.Split(current_display_raw, ":")[1])
    current_display += 1
    return strconv.Itoa(current_display)
}

func (l *Launcher) Initialize() {
    fmt.Println("Initializing game launcher...")
}

func (l *Launcher) Launch(profile *datamodels.Profile, server string, password string, callback func()) {
    // ToDo: only one instance of Urban Terror should be launched, so button
    // should be disabled.
    fmt.Println("Launching Urban Terror...")

    done := make(chan bool, 1)

    // Create launch string.
    var launch_bin string = ""
    launch_bin, err := exec.LookPath(profile.Binary)
    if err != nil {
        fmt.Println(err.Error())
    }

    var launch_params []string
    if len(server) > 0 {
        launch_params = append(launch_params, "+connect", server)
    }
    if len(password) > 0 {
        launch_params = append(launch_params, "+password", password)
    }
    if len(profile.Additional_params) > 0 {
        additional_params := strings.Split(profile.Additional_params, " ")
        launch_params = append(launch_params, additional_params...)
    }
    if profile.Second_x_session == "1" {
        fmt.Println(launch_params)
        launch_params = append([]string{launch_bin}, launch_params...)
        display := l.findFreeDisplay()
        launch_bin, err = exec.LookPath("xinit")
        if err != nil {
            fmt.Println(err.Error())
        }
        launch_params = append(launch_params, "--", ":" + display)
    }
    fmt.Println(launch_params)
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
