// URTrator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016-2020, Stanslav N. a.k.a pztrn (or p0z1tr0n) and
// URTrator contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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

	profile_path := path.Join(home_path, "Library", "Application Support", "Quake3", "q3ut4")
	c.TEMP["DEFAULT_PROFILE_PATH"] = profile_path

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

	profile_path := path.Join(home_path, ".q3a", "q3ut4")
	c.TEMP["DEFAULT_PROFILE_PATH"] = profile_path

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

	// Verify it!
	profile_path := path.Join(homedrive, homepath_without_drive, "AppData", "UrbanTerror43", "q3ut4")
	c.TEMP["DEFAULT_PROFILE_PATH"] = profile_path

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
