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
    // Configuration from database.
    Cfg map[string]string
    // Temporary (or runtime) configuration things.
    TEMP map[string]string
}

func (c *Config) initializePathsMac() {
    fmt.Println("Initializing configuration paths...")
    home_path := os.Getenv("HOME")
    data_path := path.Join(home_path, "Library", "Application Support", "URTrator")
    fmt.Println("Will use data path: " + data_path)
    c.TEMP["DATA"] = data_path

    if _, err := os.Stat(data_path); os.IsNotExist(err) {
        os.MkdirAll(data_path, 0755)
    }
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

func (c *Config) initializePathsWin() {
    fmt.Println("Initializing configuration paths...")
    homepath_without_drive := os.Getenv("HOMEPATH")
    homedrive := os.Getenv("HOMEDRIVE")
    data_path := path.Join(homedrive, homepath_without_drive, "AppData", "Roaming", "URTrator")
    c.TEMP["DATA"] = data_path

    if _, err := os.Stat(data_path); os.IsNotExist(err) {
        os.MkdirAll(data_path, 0755)
    }
}

func (c *Config) initializeStorages() {
    c.TEMP = make(map[string]string)
    c.Cfg = make(map[string]string)
}

func (c *Config) Initialize() {
    fmt.Println("Initializing configuration storage...")
    c.initializeStorages()

    if runtime.GOOS == "linux" {
        c.initializePathsNix()
    } else if runtime.GOOS == "darwin" {
        c.initializePathsMac()
    } else if runtime.GOOS == "windows" {
        c.initializePathsWin()
    } else {
        panic("We're not ready for other OSes yet!")
    }
}
