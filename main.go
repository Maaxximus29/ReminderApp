package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Please enter the following details-")

	fmt.Println("Task Name")
	var inputTaskName string
	fmt.Scanln(&inputTaskName)

	fmt.Println("Due Date")
	var inputDueDate string
	fmt.Scanln(&inputDueDate)

	fmt.Println("Specified Time on the due date")
	var inputSpecifiedTime string
	fmt.Scanln(&inputSpecifiedTime)
	//var createdTask tasksToBeCompleted
	parsedDate, err := time.Parse("02-01-2006", inputDueDate)
	if err != nil {
		fmt.Println("The date is invalid")
	}
	parsedTime, err := time.Parse("3:04 PM", inputSpecifiedTime)
	if err != nil {
		fmt.Println("The time is invalid")
	}

	createdTask := createTask(inputTaskName, parsedDate, parsedTime)
	fmt.Printf("%T", createdTask)
}
