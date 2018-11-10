// URTrator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016-2018, Stanslav N. a.k.a pztrn (or p0z1tr0n) and
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

package cache

import (
	// stdlib
	"fmt"
	"sync"

	// local
	"gitlab.com/pztrn/urtrator/cachemodels"
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
