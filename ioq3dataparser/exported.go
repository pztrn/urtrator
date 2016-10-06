// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
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
        parsed_data[srv_config[i]] = srv_config[i + 1]
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
            player_data["nick"] = string([]byte(nickname)[1:len(nickname)-1])
            player_data["ping"] = raw_player_data[1]
            player_data["frags"] = raw_player_data[0]
            parsed_data[player_data["nick"]] = player_data
        }
    }

    return parsed_data
}
