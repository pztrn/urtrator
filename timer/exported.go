// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package timer

import (
    // stdlib
    "fmt"

    // local
    "github.com/pztrn/urtrator/configuration"
    "github.com/pztrn/urtrator/eventer"
)

var (
    Cfg *configuration.Config
    Eventer *eventer.Eventer
)

func New(e *eventer.Eventer, cc *configuration.Config) *Timer {
    Cfg = cc
    Eventer = e
    fmt.Println("Creating Timer object...")
    t := Timer{}
    return &t
}
