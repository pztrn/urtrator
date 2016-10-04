// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package context

import (
    // stdlib
    "fmt"

    // local
    "github.com/pztrn/urtrator/configuration"
    "github.com/pztrn/urtrator/database"
    "github.com/pztrn/urtrator/eventer"
    "github.com/pztrn/urtrator/launcher"
    "github.com/pztrn/urtrator/requester"

    // Github
    "github.com/mattn/go-gtk/gtk"
)

type Context struct {
    // Configuration.
    Cfg *configuration.Config
    // Database.
    Database *database.Database
    // Eventer.
    Eventer *eventer.Eventer
    // Game launcher.
    Launcher *launcher.Launcher
    // Requester, which requests server's information.
    Requester *requester.Requester
}

func (ctx *Context) Close() {
    fmt.Println("Closing URTrator...")

    ctx.Database.Close()

    // At last, close main window.
    gtk.MainQuit()
}

func (ctx *Context) initializeConfig() {
    ctx.Cfg = configuration.New()
    ctx.Cfg.Initialize()
}

func (ctx *Context) initializeDatabase() {
    ctx.Database = database.New(ctx.Cfg)
    ctx.Database.Initialize(ctx.Cfg)
    ctx.Database.Migrate()
}

func (ctx *Context) initializeEventer() {
    ctx.Eventer = eventer.New()
    ctx.Eventer.Initialize()
}

func (ctx *Context) initializeLauncher() {
    ctx.Launcher = launcher.New()
    ctx.Launcher.Initialize()
}

func (ctx *Context) initializeRequester() {
    ctx.Requester = requester.New()
    ctx.Requester.Initialize()
}

func (ctx *Context) Initialize() {
    fmt.Println("Initializing application context...")
    ctx.initializeConfig()
    ctx.initializeDatabase()
    ctx.initializeEventer()
    ctx.initializeLauncher()
    ctx.initializeRequester()
}
