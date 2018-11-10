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

package timer

import (
	// stdlib
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Timer struct {
	// Tasks.
	tasks map[string]*TimerTask
	// Tasks map mutex.
	tasksMutex sync.Mutex
}

func (t *Timer) AddTask(task *TimerTask) error {
	fmt.Println("Adding task '" + task.Name + "'...")

	_, ok := t.tasks[task.Name]
	if ok {
		errorText := "Task '" + task.Name + "' already exist! Ignoring..."
		fmt.Println(errorText)
		return errors.New(errorText)
	}

	task.InProgress = false

	curTime := time.Now()
	nextLaunch := curTime.Add(time.Duration(task.Timeout) * time.Second)
	task.NextLaunch = nextLaunch

	t.tasksMutex.Lock()
	t.tasks[task.Name] = task
	t.tasksMutex.Unlock()

	fmt.Println("Added task '" + task.Name + "' with " + strconv.Itoa(task.Timeout) + " seconds timeout")
	return nil
}

func (t *Timer) executeTasks() {
	t.tasksMutex.Lock()
	defer t.tasksMutex.Unlock()
	for taskName, task := range t.tasks {
		// Check if task should be run.
		curtime := time.Now()
		diff := curtime.Sub(task.NextLaunch)
		//fmt.Println(diff)
		if diff > 0 {
			fmt.Println("Checking task '" + taskName + "'...")
			// Check if task is already running.
			if task.InProgress {
				fmt.Println("Already executing, skipping...")
				continue
			}

			fmt.Println("Launching task '" + taskName + "'...")
			task.InProgress = true
			Eventer.LaunchEvent(task.Callee, map[string]string{})

			curtime = time.Now()
			nextlaunch := curtime.Add(time.Duration(task.Timeout) * time.Second)
			task.NextLaunch = nextlaunch
		}
	}
}

func (t *Timer) GetTaskStatus(taskName string) bool {
	t.tasksMutex.Lock()
	task, ok := t.tasks[taskName]
	t.tasksMutex.Unlock()
	if !ok {
		return false
	}

	return task.InProgress
}

func (t *Timer) Initialize() {
	fmt.Println("Initializing timer...")

	t.initializeStorage()
	Eventer.AddEventHandler("taskDone", t.SetTaskNotInProgress)

	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for range ticker.C {
			go t.executeTasks()
		}
	}()
}

func (t *Timer) initializeStorage() {
	t.tasks = make(map[string]*TimerTask)
}

func (t *Timer) RemoveTask(taskName string) {
	t.tasksMutex.Lock()
	_, ok := t.tasks[taskName]
	t.tasksMutex.Unlock()
	if !ok {
		return
	}

	t.tasksMutex.Lock()
	delete(t.tasks, taskName)
	t.tasksMutex.Unlock()
}

func (t *Timer) SetTaskNotInProgress(data map[string]string) {
	t.tasksMutex.Lock()
	defer t.tasksMutex.Unlock()
	_, ok := t.tasks[data["taskName"]]
	if !ok {
		return
	}

	t.tasks[data["taskName"]].InProgress = false
}
