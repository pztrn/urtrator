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
    "runtime"
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

    runtime.UnlockOSThread()
}

func (d *Database) Initialize(cfg *configuration.Config) {
    fmt.Println("Initializing database...")

    runtime.LockOSThread()

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
