// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package eventer

import (
    crand "crypto/rand"
    "errors"
    "fmt"
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

    for _, val := range e.events[event] {
        val(data)
    }

    return nil
}
