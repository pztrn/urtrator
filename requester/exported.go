// URTrator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016-2020, Stanslav N. a.k.a pztrn (or p0z1tr0n) and
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

package requester

import (
	// stdlib
	"fmt"

	// local
	"go.dev.pztrn.name/urtrator/cache"
	"go.dev.pztrn.name/urtrator/configuration"
	"go.dev.pztrn.name/urtrator/eventer"
	"go.dev.pztrn.name/urtrator/timer"
)

var (
	Cache   *cache.Cache
	Cfg     *configuration.Config
	Eventer *eventer.Eventer
	Timer   *timer.Timer
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
