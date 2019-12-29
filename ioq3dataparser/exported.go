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

package ioq3dataparser

import (
	// stdlib
	"strings"
)

func ParseInfoToMap(data string) map[string]string {
	parsed_data := make(map[string]string)

	srv_config := strings.Split(data, "\\")
	srv_config = srv_config[1:]
	// Parse server configuration into passed server's datamodel.
	for i := 0; i < len(srv_config[1:]); i = i + 2 {
		parsed_data[srv_config[i]] = srv_config[i+1]
	}

	return parsed_data
}

func ParsePlayersInfoToMap(data string) map[string]map[string]string {
	parsed_data := make(map[string]map[string]string)

	// Structure: frags|ping|nick
	raw_data := strings.Split(data, "\\")
	for i := range raw_data {
		raw_player_data := strings.Split(raw_data[i], " ")
		player_data := make(map[string]string)
		if len(raw_player_data) > 1 {
			nickname := strings.Join(raw_player_data[2:], " ")
			player_data["nick"] = string([]byte(nickname)[1 : len(nickname)-1])
			player_data["ping"] = raw_player_data[1]
			player_data["frags"] = raw_player_data[0]
			parsed_data[player_data["nick"]] = player_data
		}
	}

	return parsed_data
}
