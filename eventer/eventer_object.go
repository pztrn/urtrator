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

package eventer

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	//"reflect"

	// github
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

type Eventer struct {
	// Events
	events map[string]map[string]func(data map[string]string)
}

func (e *Eventer) AddEventHandler(event string, handler func(data map[string]string)) {
	_, ok := e.events[event]
	if !ok {
		e.events[event] = make(map[string]func(data map[string]string))
	}
	event_id_raw := make([]byte, 16)
	crand.Read(event_id_raw)
	event_id := fmt.Sprintf("%x", event_id_raw)
	e.events[event][event_id] = handler
}

func (e *Eventer) Initialize() {
	e.initializeStorage()
}

func (e *Eventer) initializeStorage() {
	e.events = make(map[string]map[string]func(data map[string]string))
}

func (e *Eventer) LaunchEvent(event string, data map[string]string) error {
	_, ok := e.events[event]
	if !ok {
		return errors.New("Event " + event + " not found!")
	}

	fmt.Println("Launching event " + event)
	glib.IdleAdd(func() bool {
		e.reallyLaunchEvent(event, data)
		return false
	})

	for {
		if gtk.EventsPending() {
			gtk.MainIteration()
		} else {
			break
		}
	}

	return nil
}

func (e *Eventer) reallyLaunchEvent(event string, data map[string]string) {
	fmt.Println("Really launching event " + event + "...")
	for _, val := range e.events[event] {
		val(data)
	}
}
