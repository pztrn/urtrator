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

package database

import (
	// stdlib
	//"database/sql"
	"fmt"
	"path"
	"runtime"
	"strconv"

	// local
	"go.dev.pztrn.name/urtrator/configuration"
	"go.dev.pztrn.name/urtrator/datamodels"

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
