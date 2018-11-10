// URTrator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016-2018, Stanslav N. a.k.a pztrn (or p0z1tr0n) and
// URTrator contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject
// to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
// CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package context

import (
	// stdlib
	"errors"
	"fmt"

	// local
	"gitlab.com/pztrn/urtrator/cache"
	"gitlab.com/pztrn/urtrator/clipboardwatcher"
	"gitlab.com/pztrn/urtrator/colorizer"
	"gitlab.com/pztrn/urtrator/configuration"
	"gitlab.com/pztrn/urtrator/database"
	"gitlab.com/pztrn/urtrator/eventer"
	"gitlab.com/pztrn/urtrator/launcher"
	"gitlab.com/pztrn/urtrator/requester"
	"gitlab.com/pztrn/urtrator/timer"
	"gitlab.com/pztrn/urtrator/translator"

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
	// Translator.
	Translator *translator.Translator
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

func (ctx *Context) initializeTranslator() {
	ctx.Translator = translator.New(ctx.Cfg)
	ctx.Translator.Initialize()
}

func (ctx *Context) Initialize() {
	fmt.Println("Initializing application context...")
	ctx.initializeColorizer()
	ctx.initializeConfig()
	ctx.initializeDatabase()
	ctx.initializeTranslator()
	ctx.initializeEventer()
	ctx.initializeCache()
	ctx.initializeLauncher()
	ctx.initializeTimer()
	ctx.initializeRequester()
}
