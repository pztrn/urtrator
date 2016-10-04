// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package datamodels

type Server struct {
    // Server's address
    Ip                  string          `db:"ip"`
    // Server's port
    Port                string          `db:"port"`
    // Server's name
    Name                string          `db:"name"`
    // Current players count
    Players             string          `db:"players"`
    // Maximum players
    Maxplayers          string          `db:"maxplayers"`
    // Ping
    Ping                string          `db:"ping"`
    // Gametype. See Urban Terror documentation on relationship.
    Gamemode            string          `db:"gamemode"`
    // Current map
    Map                 string          `db:"map"`
    // Server's software version
    Version             string          `db:"version"`
    // Is server was favorited?
    Favorite            string          `db:"favorite"`
}
