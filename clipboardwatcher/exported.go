// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package clipboardwatcher

import(
    // local
    "github.com/pztrn/urtrator/cache"
    "github.com/pztrn/urtrator/eventer"
)

var (
    Cache *cache.Cache
    Eventer *eventer.Eventer
)

func New(c *cache.Cache, e *eventer.Eventer) *ClipboardWatcher {
    Cache = c
    Eventer = e
    cw := ClipboardWatcher{}
    return &cw
}
