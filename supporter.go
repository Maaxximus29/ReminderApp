package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
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
type ReminderDate []time.Time

// var allTasks Tasks
var reminderDate ReminderDate
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

func reminderChecker(r Reminder, rd ReminderDate, d Tasks) {
	defer wg.Done()
	if len(r) > 0 {
		now := time.Now()
		for _, dd := range rd {
			y1, m1, d1 := dd.Date()
			y2, m2, d2 := now.Date()
			if y1 == y2 && m1 == m2 && d1 == d2 {
				for i, rt := range r {
					if now.After(rt) {
						fmt.Print("REMINDER IS DUE\n")
						r = append(r[:i], r[i+1:]...)
						rd = append(rd[:i], rd[i+1:]...)
						d = append(d[:i], d[i+1:]...)

					}

				}

			}

		}
	} else {
		fmt.Print("All task completed")
	}
}

//for i := 0; i < len(r); i++ {

//reminderChecker := r[i]
//reminderDateChecker := rd[i]
//if formatted == reminderDateChecker.Format("02-01-2006") {
//if time.Now().After(reminderChecker) {
//	fmt.Print("REMINDER IS DUE")
//}
//}

//}

// the above snippet does not work and will only function if I use the correct type from below to do this
func grandIO(wg *sync.WaitGroup) {
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
	//reminder = append(reminder, grandCreatedTask.SpecifiedTime)
	//reminderDate = append(reminderDate, grandCreatedTask.DueDate)
	fmt.Printf("Reminder added successfully: %s %s %s\n", grandCreatedTask.TaskName, grandCreatedTask.DueDate.Format("02-01-2006"), grandCreatedTask.SpecifiedTime.Format("3:04 PM"))

	jsonEncoding(decodeJson)
	decodeJson = jsonDecoding()
	for _, individual := range decodeJson {
		reminder = append(reminder, individual.SpecifiedTime)
		reminderDate = append(reminderDate, individual.DueDate)
	}
	// reminder, reminderDate, decodeJson =
	go reminderChecker(reminder, reminderDate, decodeJson)

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
	//go reminderChecker(reminder, reminderDate)
	scanner.Scan()

	grandYesOrNo := scanner.Text()
	if grandYesOrNo == "yes" {
		grandIO(wg)
	} else {
		fmt.Println("Thank you for your time")
	}

}

//task.dueDate.Format("02-01-2006")) how is the execution work for this? left to right evaluation or right to left and how does it work
