package main

import "time"

type tasksToBeCompleted struct {
	taskName      string
	dueDate       time.Time
	specifiedTime time.Time
}

// type Tasks []tasksToBeCompleted

func createTask(taskName string, dueDate time.Time, specifiedTime time.Time) tasksToBeCompleted {
	var createdInside tasksToBeCompleted
	createdInside.taskName = taskName
	createdInside.dueDate = dueDate
	createdInside.specifiedTime = specifiedTime
	return createdInside
}
