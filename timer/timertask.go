// URTator - Urban Terror server browser and game launcher, written in
// Go.
//
// Copyright (c) 2016, Stanslav N. a.k.a pztrn (or p0z1tr0n)
// All rights reserved.
//
// Licensed under Terms and Conditions of GNU General Public License
// version 3 or any higher.
// ToDo: put full text of license here.
package timer

import (
    // stdlib
    "time"
)

type TimerTask struct {
    // Task name.
    Name        string
    // Task timeout, in seconds.
    Timeout     int
    // What we should call?
    // This should be an event name.
    Callee      string

    // Internal variables, used by Timer.
    // These variables can be defined, but they will be most likely
    // overrided after first task launch.
    // Next task launch time.
    NextLaunch  time.Time
    // Is task currently executed?
    // Kinda alternative to mutex.
    InProgress  bool
}
