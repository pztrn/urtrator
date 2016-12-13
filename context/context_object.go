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
    "errors"
    "fmt"

    // local
    "github.com/pztrn/urtrator/cache"
    "github.com/pztrn/urtrator/clipboardwatcher"
    "github.com/pztrn/urtrator/colorizer"
    "github.com/pztrn/urtrator/configuration"
    "github.com/pztrn/urtrator/database"
    "github.com/pztrn/urtrator/eventer"
    "github.com/pztrn/urtrator/launcher"
    "github.com/pztrn/urtrator/requester"
    "github.com/pztrn/urtrator/timer"

    // Github
    "github.com/mattn/go-gtk/gtk"
)

type Context struct {
    // Caching.
    Cache *cache.Cache
    // Clipboard watcher.
    Clipboard *clipboardwatcher.ClipboardWatcher
    // Colors parser and prettifier.
    Colorizer *colorizer.Colorizer
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
    // Timer.
    Timer *timer.Timer
}

func (ctx *Context) Close() error {
    fmt.Println("Closing URTrator...")

    launched := ctx.Launcher.CheckForLaunchedUrbanTerror()
    if launched != nil {
        return errors.New("Urban Terror is launched!")
    }
    ctx.Cache.FlushProfiles(map[string]string{})
    ctx.Cache.FlushServers(map[string]string{})
    ctx.Database.Close()

    // At last, close main window.
    gtk.MainQuit()
    return nil
}

func (ctx *Context) initializeCache() {
    ctx.Cache = cache.New(ctx.Database, ctx.Eventer)
    ctx.Cache.Initialize()
}

func (ctx *Context) InitializeClipboardWatcher() {
    ctx.Clipboard = clipboardwatcher.New(ctx.Cache, ctx.Eventer)
    ctx.Clipboard.Initialize()
}

func (ctx *Context) initializeColorizer() {
    ctx.Colorizer = colorizer.New()
    ctx.Colorizer.Initialize()
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
    ctx.Requester = requester.New(ctx.Cache, ctx.Eventer, ctx.Cfg, ctx.Timer)
    ctx.Requester.Initialize()
}

func (ctx *Context) initializeTimer() {
    ctx.Timer = timer.New(ctx.Eventer, ctx.Cfg)
    ctx.Timer.Initialize()
}

func (ctx *Context) Initialize() {
    fmt.Println("Initializing application context...")
    ctx.initializeColorizer()
    ctx.initializeConfig()
    ctx.initializeDatabase()
    ctx.initializeEventer()
    ctx.initializeCache()
    ctx.initializeLauncher()
    ctx.initializeTimer()
    ctx.initializeRequester()
}
