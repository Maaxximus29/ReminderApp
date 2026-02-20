package main

import (
	"bufio"
	"fmt"
	"os"
	"text/tabwriter"
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

	createdInside.taskName = taskName
	createdInside.dueDate = dueDate
	createdInside.specifiedTime = specifiedTime
	allTasks = append(allTasks, createdInside)
	return createdInside
}

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
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
		fmt.Fprintf(w, "ID\tTask\tDue Date\tTime\n")
		for i, task := range allTasks {

			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i+1, task.taskName, task.dueDate.Format("02-01-2006"), task.specifiedTime.Format("3:04 PM"))

		}
		w.Flush()

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
