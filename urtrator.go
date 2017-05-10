// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package main

import (
    // local
    "github.com/pztrn/urtrator/common"
    "github.com/pztrn/urtrator/context"
    "github.com/pztrn/urtrator/ui/gtk2"
    //"github.com/pztrn/urtrator/ui/qt5"

    // stdlib
    "fmt"
    "runtime"
)

func main() {
    fmt.Println("This is URTrator, version " + common.URTRATOR_VERSION)

    numCPUs := runtime.NumCPU()
    runtime.GOMAXPROCS(numCPUs)

    ctx := context.New()
    ctx.Initialize()

    ui := ui.NewMainWindow(ctx)
    ui.Initialize()
}
