package main

import (
	"fmt"
	"time"
)

type tasksToBeCompleted struct {
	taskName      string
	dueDate       time.Time
	specifiedTime time.Time
}

type Tasks []tasksToBeCompleted

var allTasks Tasks
var createdInside tasksToBeCompleted

func createTask(taskName string, dueDate time.Time, specifiedTime time.Time) tasksToBeCompleted {
	// var createdInside tasksToBeCompleted
	createdInside.taskName = taskName
	createdInside.dueDate = dueDate
	createdInside.specifiedTime = specifiedTime
	allTasks = append(allTasks, createdInside)
	return createdInside
}

func viewTask() {
	fmt.Printf("%T", allTasks)
}
