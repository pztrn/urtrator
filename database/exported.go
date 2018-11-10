// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package database

import (
	// Local
	"gitlab.com/pztrn/urtrator/configuration"
)

var (
	cfg *configuration.Config
)

func New(c *configuration.Config) *Database {
	cfg = c
	d := Database{}
	return &d
}
