// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package cache

import (
    // local
    event "github.com/pztrn/urtrator/eventer"
    "github.com/pztrn/urtrator/database"
)

var (
    Database *database.Database
    Eventer *event.Eventer
)

func New(d *database.Database, e *event.Eventer) *Cache {
    Database = d
    Eventer = e
    c := Cache{}
    return &c
}
