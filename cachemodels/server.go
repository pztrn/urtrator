// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package cachemodels

import (
    // local
    "github.com/pztrn/urtrator/datamodels"

    // Other
    "github.com/mattn/go-gtk/gtk"
)

type Server struct {
    Server *datamodels.Server
    AllServersIter *gtk.TreeIter
    AllServersIterSet bool
    AllServersIterInList bool

    FavServersIter *gtk.TreeIter
    FavServersIterSet bool
    FavServersIterInList bool
}
