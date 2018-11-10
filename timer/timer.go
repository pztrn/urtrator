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
		error_text := "Task '" + task.Name + "' already exist! Ignoring..."
		fmt.Println(error_text)
		return errors.New(error_text)
	}

	task.InProgress = false

	curtime := time.Now()
	nextlaunch := curtime.Add(time.Duration(task.Timeout) * time.Second)
	task.NextLaunch = nextlaunch

	t.tasksMutex.Lock()
	t.tasks[task.Name] = task
	t.tasksMutex.Unlock()

	fmt.Println("Added task '" + task.Name + "' with " + strconv.Itoa(task.Timeout) + " seconds timeout")
	return nil
}

func (t *Timer) executeTasks() {
	t.tasksMutex.Lock()
	for task_name, task := range t.tasks {
		// Check if task should be run.
		curtime := time.Now()
		diff := curtime.Sub(task.NextLaunch)
		//fmt.Println(diff)
		if diff > 0 {
			fmt.Println("Checking task '" + task_name + "'...")
			// Check if task is already running.
			if task.InProgress {
				fmt.Println("Already executing, skipping...")
				continue
			}

			fmt.Println("Launching task '" + task_name + "'...")
			task.InProgress = true
			Eventer.LaunchEvent(task.Callee, map[string]string{})

			curtime = time.Now()
			nextlaunch := curtime.Add(time.Duration(task.Timeout) * time.Second)
			task.NextLaunch = nextlaunch
		}
	}
	t.tasksMutex.Unlock()
}

func (t *Timer) GetTaskStatus(task_name string) bool {
	t.tasksMutex.Lock()
	task, ok := t.tasks[task_name]
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
		for _ = range ticker.C {
			go t.executeTasks()
		}
	}()
}

func (t *Timer) initializeStorage() {
	t.tasks = make(map[string]*TimerTask)
}

func (t *Timer) RemoveTask(task_name string) {
	t.tasksMutex.Lock()
	_, ok := t.tasks[task_name]
	t.tasksMutex.Unlock()
	if !ok {
		return
	}

	t.tasksMutex.Lock()
	delete(t.tasks, task_name)
	t.tasksMutex.Unlock()
}

func (t *Timer) SetTaskNotInProgress(data map[string]string) {
	t.tasksMutex.Lock()
	_, ok := t.tasks[data["task_name"]]
	if !ok {
		return
	}

	t.tasks[data["task_name"]].InProgress = false
	t.tasksMutex.Unlock()
}
