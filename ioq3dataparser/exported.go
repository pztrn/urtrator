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
    "fmt"
    "strings"
)

func ParseInfoToMap(data string) map[string]string {
    fmt.Println(data)

    parsed_data := make(map[string]string)

    srv_config := strings.Split(data, "\\")
    srv_config = srv_config[1:]
    // Parse server configuration into passed server's datamodel.
    for i := 0; i < len(srv_config[1:]); i = i + 2 {
        parsed_data[srv_config[i]] = srv_config[i + 1]
        fmt.Println(srv_config[i] + " => " + srv_config[i + 1])
    }

    fmt.Println(parsed_data)

    return parsed_data
}
