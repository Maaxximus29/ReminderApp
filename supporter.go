package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type TasksToBeCompleted struct {
	TaskName      string    `jason:"task"`
	DueDate       time.Time `jason:"date"`
	SpecifiedTime time.Time `jason:"time"`
}

type Tasks []TasksToBeCompleted
type Reminder []time.Time

// var allTasks Tasks
var reminder Reminder
var decodeJson Tasks
var createdInside TasksToBeCompleted

func createTask(taskName string, dueDate time.Time, specifiedTime time.Time) TasksToBeCompleted {

	createdInside.TaskName = taskName
	createdInside.DueDate = dueDate
	createdInside.SpecifiedTime = specifiedTime
	decodeJson = append(decodeJson, createdInside)
	return createdInside
}

func jsonEncoding(t Tasks) {

	data := t
	finalJson, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("storage.json", finalJson, 0644)
	if err != nil {
		panic(err)
	}
}

func jsonDecoding() []TasksToBeCompleted {
	newData, err := os.ReadFile("storage.json")
	if err != nil {
		panic(err)
	}

	var decodedTasks []TasksToBeCompleted
	err = json.Unmarshal(newData, &decodedTasks)
	if err != nil {
		panic(err)
	}

	return decodedTasks

}

func checker(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return true
}

func reminderChecker(r Reminder) {
	for i := 0; i < len(r); i++ {
		reminderChecker := time.Until(r[i])

		if reminderChecker <= 0 {
			fmt.Println("Reminder due")
			break
		}
	}
}

func grandIO() {
	checker := checker("storage.json")
	if checker == true {
		checkerNewData, err := os.ReadFile("storage.json")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(checkerNewData, &decodeJson)
	}

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
	reminder = append(reminder, grandCreatedTask.SpecifiedTime)
	fmt.Printf("Reminder added successfully: %s %s %s\n", grandCreatedTask.TaskName, grandCreatedTask.DueDate.Format("02-01-2006"), grandCreatedTask.SpecifiedTime.Format("3:04 PM"))

	jsonEncoding(decodeJson)
	decodeJson = jsonDecoding()

	fmt.Println("To view the created task please type Yes else No")
	var grandConfirmationToViewAgain string
	fmt.Scanln(&grandConfirmationToViewAgain)
	if grandConfirmationToViewAgain == "yes" {
		//fmt.Printf("%s %s %s", allTasks)
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)
		fmt.Fprintf(w, "ID\tTask\tDue Date\tTime\n")
		for i, task := range decodeJson {

			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", i+1, task.TaskName, task.DueDate.Format("02-01-2006"), task.SpecifiedTime.Format("3:04 PM"))

		}
		w.Flush()

	} else {
		fmt.Println("Thank you for your time")
	}
	fmt.Println("would you like to add more reminders? Please reply with Yes or No")
	go reminderChecker(reminder)
	scanner.Scan()

	grandYesOrNo := scanner.Text()
	if grandYesOrNo == "yes" {
		grandIO()
	} else {
		fmt.Println("Thank you for your time")
	}

}

//task.dueDate.Format("02-01-2006")) how is the execution work for this? left to right evaluation or right to left and how does it work
