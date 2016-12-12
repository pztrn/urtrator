// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package cache

import (
    // stdlib
    "fmt"
    "sync"

    // local
    "github.com/pztrn/urtrator/cachemodels"
)

type Cache struct {
    // Profiles cache.
    Profiles map[string]*cachemodels.Profile
    // Profiles cache mutex.
    ProfilesMutex sync.Mutex
    // Servers cache.
    Servers map[string]*cachemodels.Server
    // Servers cache mutex.
    ServersMutex sync.Mutex
}

func (c *Cache) Initialize() {
    fmt.Println("Initializing cache...")
    c.initializeStorages()
    c.LoadServers(map[string]string{})

    Eventer.AddEventHandler("deleteProfile", c.deleteProfile)
    Eventer.AddEventHandler("flushProfiles", c.FlushProfiles)
    Eventer.AddEventHandler("loadProfiles", c.LoadProfiles)

    Eventer.AddEventHandler("flushServers", c.FlushServers)
    Eventer.AddEventHandler("loadServersIntoCache", c.LoadServers)
}

func (c *Cache) initializeStorages() {
    // Profiles cache.
    c.Profiles = make(map[string]*cachemodels.Profile)
    // Servers cache.
    c.Servers = make(map[string]*cachemodels.Server)
}
