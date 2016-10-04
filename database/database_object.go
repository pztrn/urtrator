// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package database

import (
    // stdlib
    //"database/sql"
    "fmt"
    "path"
    "strconv"

    // local
    "github.com/pztrn/urtrator/configuration"
    "github.com/pztrn/urtrator/datamodels"

    // Other
    "github.com/jmoiron/sqlx"
    _ "github.com/mattn/go-sqlite3"
)

type Database struct {
    // Configuration.
    cfg *configuration.Config
    // Pointer to initialized database connection.
    Db *sqlx.DB
}

func (d *Database) Close() {
    fmt.Println("Closing database...")

    // Save configuration.
    // Delete previous configuration.
    d.Db.MustExec("DELETE FROM configuration")
    tx := d.Db.MustBegin()
    for k, v := range cfg.Cfg {
        cfg_item := datamodels.Configuration{}
        cfg_item.Key = k
        cfg_item.Value = v
        tx.NamedExec("INSERT INTO configuration (key, value) VALUES (:key, :value)", &cfg_item)
    }
    tx.Commit()

    d.Db.Close()
}

func (d *Database) Initialize(cfg *configuration.Config) {
    fmt.Println("Initializing database...")

    // Connect to database.
    db_path := path.Join(cfg.TEMP["DATA"], "database.sqlite3")
    fmt.Println("Database path: " + db_path)
    db, err := sqlx.Connect("sqlite3", db_path)
    if err != nil {
        fmt.Println(err.Error())
    }
    d.Db = db

    // Load configuration.
    cfgs := []datamodels.Configuration{}
    d.Db.Select(&cfgs, "SELECT * FROM configuration")
    if len(cfgs) > 0 {
        for i := range cfgs {
            cfg.Cfg[cfgs[i].Key] = cfgs[i].Value
        }
    }
}

func (d *Database) Migrate() {
    // Getting current database version.
    dbver := 0
    database := []datamodels.Database{}
    d.Db.Select(&database, "SELECT * FROM database")
    if len(database) > 0 {
        fmt.Println("Current database version: " + database[0].Version)
        dbver, _ = strconv.Atoi(database[0].Version)
    } else {
        fmt.Println("No database found, will create new one")
    }


    migrate_full(d, dbver)
}

func (d *Database) UpdateServers(data map[string]*datamodels.Server) {
    fmt.Println("Updating servers information in database...")
    raw_cached := []datamodels.Server{}
    d.Db.Select(&raw_cached, "SELECT * FROM servers")

    // Create map[string]*datamodels.Server once, so we won't iterate
    // over slice of datamodels.Server everytime.
    cached_servers := make(map[string]*datamodels.Server)
    for s := range raw_cached {
        mapping_item_name := raw_cached[s].Ip + ":" + raw_cached[s].Port
        cached_servers[mapping_item_name] = &raw_cached[s]
    }

    new_servers := make(map[string]*datamodels.Server)

    // Update our cached mapping.
    for _, s := range data {
        mapping_item_name := s.Ip + ":" + s.Port
        _, ok := cached_servers[mapping_item_name]
        if !ok {
            fmt.Println(mapping_item_name + " not found!")
            new_servers[mapping_item_name] = s
        } else {
            cached_servers[mapping_item_name].Ip = s.Ip
            cached_servers[mapping_item_name].Port = s.Port
            cached_servers[mapping_item_name].Name = s.Name
            cached_servers[mapping_item_name].Players = s.Players
            cached_servers[mapping_item_name].Maxplayers = s.Maxplayers
            cached_servers[mapping_item_name].Ping = s.Ping
            cached_servers[mapping_item_name].Map = s.Map
            cached_servers[mapping_item_name].Gamemode = s.Gamemode
            cached_servers[mapping_item_name].Version = s.Version
            cached_servers[mapping_item_name].ExtendedConfig = s.ExtendedConfig
            cached_servers[mapping_item_name].PlayersInfo = s.PlayersInfo
        }
    }

    tx := d.Db.MustBegin()
    fmt.Println("Adding new servers...")
    for _, srv := range new_servers {
        tx.NamedExec("INSERT INTO servers (ip, port, name, ping, players, maxplayers, gamemode, map, version) VALUES (:ip, :port, :name, :ping, :players, :maxplayers, :gamemode, :map, :version)", srv)
    }
    fmt.Println("Updating cached servers...")
    for _, srv := range cached_servers {
        tx.NamedExec("UPDATE servers SET name=:name, players=:players, maxplayers=:maxplayers, gamemode=:gamemode, map=:map, version=:version WHERE ip=:ip AND port=:port", &srv)
    }

    tx.Commit()
    fmt.Println("Done")
}
