// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package configuration

import (
    // stdlib
    "fmt"
    "os"
    "path"
    "runtime"
)

type Config struct {
    // Temporary (or runtime) configuration things.
    TEMP map[string]string
}

func (c *Config) initializePathsNix() {
    fmt.Println("Initializing configuration paths...")

    // Get storage path. By default we will use ~/.config/urtrator
    // directory.
    home_path := os.Getenv("HOME")
    data_path := path.Join(home_path, ".config", "urtrator")
    fmt.Println("Will use data path: " + data_path)
    c.TEMP["DATA"] = data_path

    if _, err := os.Stat(data_path); os.IsNotExist(err) {
        os.MkdirAll(data_path, 0755)
    }
}

func (c *Config) initializeStorages() {
    c.TEMP = make(map[string]string)
}

func (c *Config) Initialize() {
    fmt.Println("Initializing configuration storage...")
    c.initializeStorages()

    if runtime.GOOS == "linux" {
        c.initializePathsNix()
    } else {
        panic("We're not ready for other OSes yet!")
    }
}
