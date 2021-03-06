// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package requester

import (
    // stdlib
    "fmt"

    // local
    "github.com/pztrn/urtrator/cache"
    "github.com/pztrn/urtrator/configuration"
    "github.com/pztrn/urtrator/eventer"
    "github.com/pztrn/urtrator/timer"
)

var (
    Cache *cache.Cache
    Cfg *configuration.Config
    Eventer *eventer.Eventer
    Timer *timer.Timer
)

func New(c *cache.Cache, e *eventer.Eventer, cc *configuration.Config, t *timer.Timer) *Requester {
    Cache = c
    Cfg = cc
    Eventer = e
    Timer = t
    fmt.Println("Creating Requester object...")
    r := Requester{}
    return &r
}
