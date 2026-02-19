package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	fmt.Println("Please enter the following details-")
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Task Name")
	scanner.Scan()
	inputTaskName := scanner.Text()

	fmt.Println("Due Date")
	scanner.Scan()
	inputDueDate := scanner.Text()

	fmt.Println("Specified Time on the due date")
	scanner.Scan()
	inputSpecifiedTime := scanner.Text()
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

	fmt.Printf("Reminder added successfully: %s %s %s\n", createdTask.taskName, createdTask.dueDate.Format("02-01-2006"), createdTask.specifiedTime.Format("3:04 PM"))

	fmt.Println("To view the created task please type yes")
	var confirmationToViewAgain string
	fmt.Scanln(&confirmationToViewAgain)
	if confirmationToViewAgain == "yes" {
		//fmt.Printf("%s %s %s", allTasks)
		for _, task := range allTasks {
			fmt.Printf("%s\n", task.taskName)
			fmt.Printf("%s\n", task.dueDate.Format("02-01-2006"))
			fmt.Printf("%s\n", task.specifiedTime.Format("3:04 PM"))

		}

	} else {
		fmt.Println("Thank you for your time")
	}

}
