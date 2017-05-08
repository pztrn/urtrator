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
    "strconv"

    // local
    "github.com/pztrn/urtrator/cachemodels"
    "github.com/pztrn/urtrator/datamodels"
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

    _, ok := c.Profiles[data["profile_name"]]
    if ok {
        c.ProfilesMutex.Lock()
        delete(c.Profiles, data["profile_name"])
        c.ProfilesMutex.Unlock()
    }

    _, ok1 := c.Profiles[data["profile_name"]]
    if !ok1 {
        fmt.Println("Profile deleted")
        Database.Db.MustExec(Database.Db.Rebind("DELETE FROM urt_profiles WHERE name=?"), data["profile_name"])
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
        cached_profiles[raw_profiles[i].Name] = &raw_profiles[i]
    }

    new_profiles := make(map[string]*datamodels.Profile)

    for _, profile := range c.Profiles {
        _, ok := cached_profiles[profile.Profile.Name]
        if !ok {
            fmt.Println("Flushing new profile " + profile.Profile.Name)
            new_profiles[profile.Profile.Name] = &datamodels.Profile{}
            new_profiles[profile.Profile.Name].Name = profile.Profile.Name
            new_profiles[profile.Profile.Name].Version = profile.Profile.Version
            new_profiles[profile.Profile.Name].Binary = profile.Profile.Binary
            new_profiles[profile.Profile.Name].Second_x_session = profile.Profile.Second_x_session
            new_profiles[profile.Profile.Name].Additional_params = profile.Profile.Additional_params
        }
    }

    tx := Database.Db.MustBegin()
    fmt.Println("Adding new profiles...")
    for _, profile := range new_profiles {
        tx.NamedExec("INSERT INTO urt_profiles (name, version, binary, second_x_session, additional_parameters) VALUES (:name, :version, :binary, :second_x_session, :additional_parameters)", &profile)
    }
    fmt.Println("Updating existing profiles...")
    for _, profile := range cached_profiles {
        tx.NamedExec("UPDATE urt_profiles SET name=:name, version=:version, binary=:binary, second_x_session=:second_x_session, additional_parameters=:additional_parameters WHERE name=:name", &profile)
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

    for _, profile := range raw_profiles {
        c.Profiles[profile.Name] = &cachemodels.Profile{}
        c.Profiles[profile.Name].Profile = &datamodels.Profile{}
        c.Profiles[profile.Name].Profile.Name = profile.Name
        c.Profiles[profile.Name].Profile.Version = profile.Version
        c.Profiles[profile.Name].Profile.Binary = profile.Binary
        c.Profiles[profile.Name].Profile.Second_x_session = profile.Second_x_session
        c.Profiles[profile.Name].Profile.Additional_params = profile.Additional_params
    }

    fmt.Println("Load completed. Loaded " + strconv.Itoa(len(c.Profiles)) + " profiles.")
}
