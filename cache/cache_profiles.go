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

package cache

import (
	// stdlib
	"fmt"
	"strconv"

	// local
	"go.dev.pztrn.name/urtrator/cachemodels"
	"go.dev.pztrn.name/urtrator/datamodels"
)

func (c *Cache) CreateProfile(name string) {
	fmt.Println("Creating profile " + name)
	_, ok := c.Profiles[name]

	if !ok {
		c.ProfilesMutex.Lock()
		c.Profiles[name] = &cachemodels.Profile{}
		c.Profiles[name].Profile = &datamodels.Profile{}
		c.ProfilesMutex.Unlock()
	}
}

func (c *Cache) deleteProfile(data map[string]string) {
	fmt.Println("Deleting profile " + data["profile_name"])

	c.ProfilesMutex.Lock()
	_, ok := c.Profiles[data["profile_name"]]
	c.ProfilesMutex.Unlock()
	if ok {
		c.ProfilesMutex.Lock()
		delete(c.Profiles, data["profile_name"])
		c.ProfilesMutex.Unlock()
	}

	c.ProfilesMutex.Lock()
	_, ok1 := c.Profiles[data["profile_name"]]
	c.ProfilesMutex.Unlock()
	if !ok1 {
		Database.Db.MustExec(Database.Db.Rebind("DELETE FROM urt_profiles WHERE name=?"), data["profile_name"])
		fmt.Println("Profile deleted")
	} else {
		fmt.Println("Something goes wrong! Profile is still here!")
	}
}

func (c *Cache) FlushProfiles(data map[string]string) {
	fmt.Println("Flushing profiles to database...")

	raw_profiles := []datamodels.Profile{}
	err := Database.Db.Select(&raw_profiles, "SELECT * FROM urt_profiles")
	if err != nil {
		fmt.Println(err.Error())
	}

	cached_profiles := make(map[string]*datamodels.Profile)
	for i := range raw_profiles {
		cached_profiles[raw_profiles[i].Name] = c.Profiles[raw_profiles[i].Name].Profile
	}

	new_profiles := make(map[string]*datamodels.Profile)

	c.ProfilesMutex.Lock()
	for _, profile := range c.Profiles {
		_, ok := cached_profiles[profile.Profile.Name]
		if !ok {
			fmt.Println("Flushing new profile " + profile.Profile.Name)
			new_profiles[profile.Profile.Name] = profile.Profile
		}
	}
	c.ProfilesMutex.Unlock()

	tx := Database.Db.MustBegin()
	fmt.Println("Adding new profiles...")
	for _, profile := range new_profiles {
		tx.NamedExec("INSERT INTO urt_profiles (name, version, binary, second_x_session, additional_parameters, profile_path) VALUES (:name, :version, :binary, :second_x_session, :additional_parameters, :profile_path)", &profile)
	}
	fmt.Println("Updating existing profiles...")
	for _, profile := range cached_profiles {
		fmt.Println(fmt.Sprintf("%+v", profile))
		tx.NamedExec("UPDATE urt_profiles SET name=:name, version=:version, binary=:binary, second_x_session=:second_x_session, additional_parameters=:additional_parameters, profile_path=:profile_path WHERE name=:name", &profile)
	}
	tx.Commit()
	fmt.Println("Done")
}

func (c *Cache) LoadProfiles(data map[string]string) {
	fmt.Println("Loading profiles to cache...")

	raw_profiles := []datamodels.Profile{}
	err := Database.Db.Select(&raw_profiles, "SELECT * FROM urt_profiles")
	if err != nil {
		fmt.Println(err.Error())
	}

	c.ProfilesMutex.Lock()
	for _, profile := range raw_profiles {
		c.Profiles[profile.Name] = &cachemodels.Profile{}
		c.Profiles[profile.Name].Profile = &profile
	}
	c.ProfilesMutex.Unlock()

	fmt.Println("Load completed. Loaded " + strconv.Itoa(len(c.Profiles)) + " profiles.")
}
