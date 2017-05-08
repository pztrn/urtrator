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

    // local
    "github.com/pztrn/urtrator/cachemodels"
    "github.com/pztrn/urtrator/datamodels"
)

func (c *Cache) CreateServer(addr string) {
    _, ok := c.Servers[addr]
    if !ok {
        c.ServersMutex.Lock()
        c.Servers[addr] = &cachemodels.Server{}
        c.Servers[addr].Server = &datamodels.Server{}
        c.ServersMutex.Unlock()
    } else {
        fmt.Println("Server " + addr + " already exist.")
    }
}

// Flush servers to database.
func (c *Cache) FlushServers(data map[string]string) {
    fmt.Println("Updating servers information in database...")
    raw_cached := []datamodels.Server{}
    Database.Db.Select(&raw_cached, "SELECT * FROM servers")

    // Create map[string]*datamodels.Server once, so we won't iterate
    // over slice of datamodels.Server everytime.
    cached_servers := make(map[string]*datamodels.Server)
    for s := range raw_cached {
        mapping_item_name := raw_cached[s].Ip + ":" + raw_cached[s].Port
        cached_servers[mapping_item_name] = &raw_cached[s]
    }

    new_servers := make(map[string]*datamodels.Server)

    // Update our cached mapping.
    for _, s := range c.Servers {
        mapping_item_name := s.Server.Ip + ":" + s.Server.Port
        _, ok := cached_servers[mapping_item_name]
        if !ok {
            fmt.Println(mapping_item_name + " not found!")
            new_servers[mapping_item_name] = &datamodels.Server{}
            new_servers[mapping_item_name].Ip = s.Server.Ip
            new_servers[mapping_item_name].Port = s.Server.Port
            new_servers[mapping_item_name].Name = s.Server.Name
            new_servers[mapping_item_name].Players = s.Server.Players
            new_servers[mapping_item_name].Bots = s.Server.Bots
            new_servers[mapping_item_name].Maxplayers = s.Server.Maxplayers
            new_servers[mapping_item_name].Ping = s.Server.Ping
            new_servers[mapping_item_name].Map = s.Server.Map
            new_servers[mapping_item_name].Gamemode = s.Server.Gamemode
            new_servers[mapping_item_name].Version = s.Server.Version
            new_servers[mapping_item_name].ExtendedConfig = s.Server.ExtendedConfig
            new_servers[mapping_item_name].PlayersInfo = s.Server.PlayersInfo
            new_servers[mapping_item_name].IsPrivate = s.Server.IsPrivate
            new_servers[mapping_item_name].Favorite = s.Server.Favorite
            new_servers[mapping_item_name].ProfileToUse = s.Server.ProfileToUse
            new_servers[mapping_item_name].Password = s.Server.Password
        } else {
            cached_servers[mapping_item_name].Ip = s.Server.Ip
            cached_servers[mapping_item_name].Port = s.Server.Port
            cached_servers[mapping_item_name].Name = s.Server.Name
            cached_servers[mapping_item_name].Players = s.Server.Players
            cached_servers[mapping_item_name].Bots = s.Server.Bots
            cached_servers[mapping_item_name].Maxplayers = s.Server.Maxplayers
            cached_servers[mapping_item_name].Ping = s.Server.Ping
            cached_servers[mapping_item_name].Map = s.Server.Map
            cached_servers[mapping_item_name].Gamemode = s.Server.Gamemode
            cached_servers[mapping_item_name].Version = s.Server.Version
            cached_servers[mapping_item_name].ExtendedConfig = s.Server.ExtendedConfig
            cached_servers[mapping_item_name].PlayersInfo = s.Server.PlayersInfo
            cached_servers[mapping_item_name].IsPrivate = s.Server.IsPrivate
            cached_servers[mapping_item_name].Favorite = s.Server.Favorite
            cached_servers[mapping_item_name].ProfileToUse = s.Server.ProfileToUse
            cached_servers[mapping_item_name].Password = s.Server.Password
        }
    }

    tx := Database.Db.MustBegin()
    fmt.Println("Adding new servers...")
    if len(new_servers) > 0 {
        for _, srv := range new_servers {
            tx.NamedExec("INSERT INTO servers (ip, port, name, ping, players, maxplayers, gamemode, map, version, extended_config, players_info, is_private, favorite, profile_to_use, bots) VALUES (:ip, :port, :name, :ping, :players, :maxplayers, :gamemode, :map, :version, :extended_config, :players_info, :is_private, :favorite, :profile_to_use, :bots)", srv)
        }
    }
    fmt.Println("Updating cached servers...")
    for _, srv := range cached_servers {
        _, err := tx.NamedExec("UPDATE servers SET name=:name, players=:players, maxplayers=:maxplayers, gamemode=:gamemode, map=:map, ping=:ping, version=:version, extended_config=:extended_config, favorite=:favorite, password=:password, players_info=:players_info, is_private=:is_private, profile_to_use=:profile_to_use, bots=:bots WHERE ip=:ip AND port=:port", &srv)
        if err != nil {
            fmt.Println(err.Error())
        }
    }

    tx.Commit()
    fmt.Println("Done")
}

func (c *Cache) LoadServers(data map[string]string) {
    fmt.Println("Loading servers into cache...")
    c.Servers = make(map[string]*cachemodels.Server)
    // Getting servers from database.
    raw_servers := []datamodels.Server{}
    err := Database.Db.Select(&raw_servers, "SELECT * FROM servers")
    if err != nil {
        fmt.Println(err.Error())
    }

    // Due to nature of pointers and goroutines thing (?) this should
    // be done in this way.
    for _, server := range raw_servers {
        key := server.Ip + ":" + server.Port
        c.CreateServer(key)
        c.Servers[key].Server.Name = server.Name
        c.Servers[key].Server.Ip = server.Ip
        c.Servers[key].Server.Port = server.Port
        c.Servers[key].Server.Players = server.Players
        c.Servers[key].Server.Bots = server.Bots
        c.Servers[key].Server.Maxplayers = server.Maxplayers
        c.Servers[key].Server.Ping = server.Ping
        c.Servers[key].Server.Gamemode = server.Gamemode
        c.Servers[key].Server.Map = server.Map
        c.Servers[key].Server.Version = server.Version
        c.Servers[key].Server.Favorite = server.Favorite
        c.Servers[key].Server.Password = server.Password
        c.Servers[key].Server.ProfileToUse = server.ProfileToUse
        c.Servers[key].Server.ExtendedConfig = server.ExtendedConfig
        c.Servers[key].Server.PlayersInfo = server.PlayersInfo
        c.Servers[key].Server.IsPrivate = server.IsPrivate
    }
    fmt.Println("Load completed.")
}
