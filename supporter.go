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
// func confirmation(yesOrNo string) {
// 	if yesOrNo == "yes" {
// 		scanner := bufio.NewScanner(os.Stdin)
// 		fmt.Println("Please enter the following details- ")
// 		fmt.Println("Task Name")
// 		scanner.Scan()
// 		newInputTaskName := scanner.Text()
// 		fmt.Println("Due Date")
// 		scanner.Scan()
// 		newDueDate := scanner.Text()
// 		fmt.Println("Specified Time on the Due Date")
// 		scanner.Scan()
// 		newInputSpecifiedTime := scanner.Text()
// 		newParsedDate, err := time.Parse("02-01-2006", newDueDate)
// 		if err != nil {
// 			fmt.Println("The Due Date is invalid")
// 		}
// 		newParsedTime, err := time.Parse("3:04 PM", newInputSpecifiedTime)
// 		if err != nil {
// 			fmt.Println("The time is invalid")
// 		}

// 		newCreatedTask := createTask(newInputTaskName, newParsedDate, newParsedTime)
// 		fmt.Printf("Reminder added successfully: %s %s %s\n", newCreatedTask.taskName, newCreatedTask.dueDate.Format("02-01-2006"), newCreatedTask.specifiedTime.Format("3:04 PM"))

// 		fmt.Println("To view the created task please type yes")
// 		var newConfirmationToViewAgain string
// 		fmt.Scanln(&newConfirmationToViewAgain)
// 		if newConfirmationToViewAgain == "yes" {
// 			for _, newTask := range allTasks {
// 				fmt.Printf("%s\n", newTask.taskName)
// 				fmt.Printf("%s\n", newTask.dueDate)
// 				fmt.Printf("%s\n", newTask.specifiedTime)
// 			}
// 		} else {
// 			fmt.Println("Thank you for your time")
// 		}

// 	} else {

// 		fmt.Println("Thank you for your time")

// 	}
// }

// func printer() {
// 	fmt.Println("Would you like to make new reminders? Please reply with yes or no")
// }

func grandIO() {
	fmt.Println("Please enter the following details-")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Task Name")
	scanner.Scan()
	grandTaskName := scanner.Text()
	fmt.Println("Due Date")
	scanner.Scan()
	grandDueDate := scanner.Text()
	grandParsedDate, err := time.Parse("02-01-2006", grandDueDate)
	if err != nil {
		fmt.Println("The Due Date is invalid")
	}
	fmt.Println("Specified Time on the Due Date")
	scanner.Scan()
	grandSpecifiedTime := scanner.Text()
	grandParsedTime, err := time.Parse("3:04 PM", grandSpecifiedTime)
	if err != nil {
		fmt.Println("The time is invalid")
	}

	grandCreatedTask := createTask(grandTaskName, grandParsedDate, grandParsedTime)
	fmt.Printf("Reminder added successfully: %s %s %s\n", grandCreatedTask.taskName, grandCreatedTask.dueDate.Format("02-01-2006"), grandCreatedTask.specifiedTime.Format("3:04 PM"))

	fmt.Println("To view the created task please type yes")
	var grandConfirmationToViewAgain string
	fmt.Scanln(&grandConfirmationToViewAgain)
	if grandConfirmationToViewAgain == "yes" {
		//fmt.Printf("%s %s %s", allTasks)
		for _, task := range allTasks {
			fmt.Printf("%s\n", task.taskName)
			fmt.Printf("%s\n", task.dueDate.Format("02-01-2006"))
			fmt.Printf("%s\n", task.specifiedTime.Format("3:04 PM"))

		}

	} else {
		fmt.Println("Thank you for your time")
	}
	fmt.Println("would you like to add more reminders? Please reply with yes or no")
	scanner.Scan()

	grandYesOrNo := scanner.Text()
	if grandYesOrNo == "yes" {
		grandIO()
	} else {
		fmt.Println("Thank you for your time")
	}

}

//task.dueDate.Format("02-01-2006")) how is the execution work for this? I mean the operand evaluation
