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
	if version == 4 {
		four_to_five(db)
		version = 5
	}
	if version == 5 {
		five_to_six(db)
		version = 6
	}
	if version == 6 {
		six_to_seven(db)
		version = 7
	}
	if version == 7 {
		seven_to_eight(db)
		version = 8
	}
	if version == 8 {
		eight_to_nine(db)
		version = 9
	}
	if version == 9 {
		nine_to_ten(db)
		version = 10
	}
	if version == 10 {
		ten_to_eleven(db)
		version = 11
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
	db.Db.MustExec("UPDATE database SET version=4")
}

// Server's passwords.
func four_to_five(db *Database) {
	fmt.Println("Upgrading database from 4 to 5...")
	db.Db.MustExec("ALTER TABLE servers ADD password VARCHAR(64) DEFAULT ''")
	db.Db.MustExec("UPDATE database SET version=5")
}

// Profile for server.
func five_to_six(db *Database) {
	fmt.Println("Upgrading database from 5 to 6...")
	db.Db.MustExec("ALTER TABLE servers ADD profile_to_use VARCHAR(128) DEFAULT ''")
	db.Db.MustExec("UPDATE database SET version=6")
}

// Configuration storage.
func six_to_seven(db *Database) {
	fmt.Println("Upgrading database from 6 to 7...")
	db.Db.MustExec("CREATE TABLE configuration (key VARCHAR(128) NOT NULL, value VARCHAR(1024) NOT NULL)")
	db.Db.MustExec("UPDATE database SET version=7")
}

// Server's extended information.
func seven_to_eight(db *Database) {
	fmt.Println("Upgrading database from 7 to 8...")
	db.Db.MustExec("ALTER TABLE servers ADD extended_config VARCHAR(4096) NOT NULL DEFAULT ''")
	db.Db.MustExec("ALTER TABLE servers ADD players_info VARCHAR(8192) NOT NULL DEFAULT ''")
	db.Db.MustExec("UPDATE database SET version=8")
}

// Is server private flag.
func eight_to_nine(db *Database) {
	fmt.Println("Upgrading database from 8 to 9...")
	db.Db.MustExec("ALTER TABLE servers ADD is_private VARCHAR(1) NOT NULL DEFAULT '0'")
	db.Db.MustExec("UPDATE database SET version=9")
}

// Bots count.
func nine_to_ten(db *Database) {
	fmt.Println("Upgrading database from 9 to 10...")
	db.Db.MustExec("ALTER TABLE servers ADD bots VARCHAR(2) NOT NULL DEFAULT '0'")
	db.Db.MustExec("UPDATE database SET version=10")
}

// Urban terror's profile path.
func ten_to_eleven(db *Database) {
	fmt.Println("Upgrading database from 10 to 11...")
	db.Db.MustExec("ALTER TABLE urt_profiles ADD profile_path VARCHAR(4096) NOT NULL DEFAULT '~/.q3ut4'")
	db.Db.MustExec("UPDATE database SET version=11")
}
