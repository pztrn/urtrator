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
    "os/exec"
)

type Launcher struct {}

func (l *Launcher) Initialize() {
    fmt.Println("Initializing game launcher...")
}

func (l *Launcher) Launch(binary string, params string, server string, password string, second_x bool, callback func()) {
    // ToDo: only one instance of Urban Terror should be launched, so button
    // should be disabled.
    fmt.Println("Launching Urban Terror...")

    done := make(chan bool, 1)
    // Create launch string.
    var launch_string string = params + " +connect " + server
    if len(password) > 0 {
        launch_string += " +password " + password
    }
    fmt.Println("Final command: " + launch_string)
    go func() {
        go func() {
            cmd := exec.Command(binary, launch_string)
            err := cmd.Run()
            if err != nil {
                fmt.Println("Launch error: " + err.Error())
            }
            done <- true
        }()

        select {
        case <- done:
            callback()
        }
    }()
}
