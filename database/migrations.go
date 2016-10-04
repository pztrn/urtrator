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
    "fmt"
)

var start_schema = `
DROP TABLE IF EXISTS database;
CREATE TABLE database (
    version         VARCHAR(10)     NOT NULL
);

DROP TABLE IF EXISTS servers;
CREATE TABLE servers (
    ip              VARCHAR(128)    NOT NULL,
    port            VARCHAR(5)      NOT NULL,
    name            VARCHAR(128)    NOT NULL,
    players         VARCHAR(2)      NOT NULL,
    maxplayers      VARCHAR(2)      NOT NULL,
    ping            VARCHAR(4),
    gamemode        VARCHAR(1)      NOT NULL,
    map             VARCHAR(64)     NOT NULL,
    version         VARCHAR(5)      NOT NULL
);

INSERT INTO database (version) VALUES (1);
`

// Migrate database to latest version.
// ToDo: make it more good :).
func migrate_full(db *Database, version int) {
    if version < 1 {
        start_to_one(db)
        version = 1
    }
    if version == 1 {
        one_to_two(db)
        version = 2
    }
    if version == 2 {
        two_to_three(db)
        version = 3
    }
    if version == 3 {
        three_to_four(db)
        version = 4
    }
}

// Initial database structure.
func start_to_one(db *Database) {
    fmt.Println("Upgrading database from 0 to 1...")
    db.Db.MustExec(start_schema)
}

// Favorite server mark.
func one_to_two(db *Database) {
    fmt.Println("Upgrading database from 1 to 2...")
    db.Db.MustExec("ALTER TABLE servers ADD favorite VARCHAR(1) DEFAULT '0'")
    db.Db.MustExec("UPDATE database SET version=2")
}

// URTRator settings and Urban Terror profiles.
func two_to_three(db *Database) {
    fmt.Println("Upgrading database from 2 to 3...")
    db.Db.MustExec("DROP TABLE IF EXISTS settings")
    db.Db.MustExec("CREATE TABLE settings (show_tray_icon VARCHAR(1) NOT NULL DEFAULT '0', enable_autoupdate VARCHAR(1) NOT NULL DEFAULT '0')")
    db.Db.MustExec("DROP TABLE IF EXISTS urt_profiles")
    db.Db.MustExec("CREATE TABLE urt_profiles (name VARCHAR(128) NOT NULL, version VARCHAR(5) NOT NULL DEFAULT '4.3', binary VARCHAR(1024) NOT NULL, second_x_session VARCHAR(1) NOT NULL DEFAULT '0', additional_parameters VARCHAR(1024) NOT NULL DEFAULT '')")
    db.Db.MustExec("UPDATE database SET version=3")
}

// UrT version inconsistency.
func three_to_four(db *Database) {
    fmt.Println("Upgrading database from 3 to 4...")
    db.Db.MustExec("UPDATE urt_profiles SET version='4.3.0' WHERE version='4.3.000'")
}
