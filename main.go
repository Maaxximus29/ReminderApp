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

	fmt.Println("Specified Time on the Due Date")
	scanner.Scan()
	inputSpecifiedTime := scanner.Text()
	//var createdTask tasksToBeCompleted
	parsedDate, err := time.Parse("02-01-2006", inputDueDate)
	if err != nil {
		fmt.Println("The Due Date is invalid")
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
	fmt.Println("Would you like to make new reminders? Please reply with yes or no")
	//scanner.Scan()
	if scanner.Scan() {
		yesOrNo := scanner.Text()
		if yesOrNo == "yes" {
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
	fmt.Println("Would you like to make new reminders? Please reply with yes or no")
	scanner.Scan()
	newYesOrNo := scanner.Text()
	confirmation(newYesOrNo)

}

//I could use a loop but that would mean that I ask the user on how many reminders he is going to add which is not a good idea
//I could use a grand if else block
//the entire code would be inside a giant function on the other page
//there would be a grand scanner which stores the value of the very first yes then at the very end of the function with the grand scanner
//the function will call itself recursively
//if the user types in yes then the same loop would continue but if the user replies with no then the loop will terminate
