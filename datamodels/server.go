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

package datamodels

type Server struct {
	// Server's address
	Ip string `db:"ip"`
	// Server's port
	Port string `db:"port"`
	// Server's name
	Name string `db:"name"`
	// Current players count
	Players string `db:"players"`
	// Bots count
	Bots string `db:"bots"`
	// Maximum players
	Maxplayers string `db:"maxplayers"`
	// Ping
	Ping string `db:"ping"`
	// Gametype. See Urban Terror documentation on relationship.
	Gamemode string `db:"gamemode"`
	// Current map
	Map string `db:"map"`
	// Server's software version
	Version string `db:"version"`
	// Is server was favorited?
	Favorite string `db:"favorite"`
	// Server's password.
	Password string `db:"password"`
	// Profile to use with server.
	ProfileToUse string `db:"profile_to_use"`
	// Extended server's configuration.
	ExtendedConfig string `db:"extended_config"`
	// Players information.
	PlayersInfo string `db:"players_info"`
	// Is server private?
	IsPrivate string `db:"is_private"`
}
