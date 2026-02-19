package main

import (
	"bufio"
	"fmt"
	"os"
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

//	func viewTask() {
//		fmt.Printf("%T", allTasks)
//	}
func confirmation(yesOrNo string) {
	if yesOrNo == "yes" {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please enter the following details- ")
		fmt.Println("Task Name")
		scanner.Scan()
		newInputTaskName := scanner.Text()
		fmt.Println("Due Date")
		scanner.Scan()
		newDueDate := scanner.Text()
		fmt.Println("Specified Time on the Due Date")
		scanner.Scan()
		newInputSpecifiedTime := scanner.Text()
		newParsedDate, err := time.Parse("02-01-2006", newDueDate)
		if err != nil {
			fmt.Println("The Due Date is invalid")
		}
		newParsedTime, err := time.Parse("3:04 PM", newInputSpecifiedTime)
		if err != nil {
			fmt.Println("The time is invalid")
		}

		newCreatedTask := createTask(newInputTaskName, newParsedDate, newParsedTime)
		fmt.Printf("Reminder added successfully: %s %s %s\n", newCreatedTask.taskName, newCreatedTask.dueDate.Format("02-01-2006"), newCreatedTask.specifiedTime.Format("3:04 PM"))

		fmt.Println("To view the created task please type yes")
		var newConfirmationToViewAgain string
		fmt.Scanln(&newConfirmationToViewAgain)
		if newConfirmationToViewAgain == "yes" {
			for _, newTask := range allTasks {
				fmt.Printf("%s\n", newTask.taskName)
				fmt.Printf("%s\n", newTask.dueDate)
				fmt.Printf("%s\n", newTask.specifiedTime)
			}
		} else {
			fmt.Println("Thank you for your time")
		}

	} else {

		fmt.Println("Thank you for your time")

	}
}
