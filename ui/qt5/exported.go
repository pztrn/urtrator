// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package ui

import (
	// local
	"github.com/pztrn/urtrator/context"
)

var (
	ctx *context.Context
)

func NewMainWindow(c *context.Context) *MainWindow {
	ctx = c
	m := MainWindow{}
	return &m
}
