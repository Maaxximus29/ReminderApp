package main

func main() {

	grandIO()

	select {}

}

//I could use a loop but that would mean that I ask the user on how many reminders he is going to add which is not a good idea
//I could use a grand if else block
//the entire code would be inside a giant function on the other page
//there would be a grand scanner which stores the value of the very first yes then at the very end of the function with the grand scanner
//the function will call itself recursively
//if the user types in yes then the same loop would continue but if the user replies with no then the loop will terminate
//20 Feb Code sprint-
//Write the operations inside the granIO function in supporter.go
//It worked

//Ran into a new issue with persistent storage, I need to somehow get the program to store the values once I exit the program
//ideas
//I want the JSON value to be persistent in order to append to it and later fetch it
//I enter the program and check if the JSON file is empty or now, if it isnt empty then I immediately append to the slice
//Not sure if this works or not, I will use this method to get things working, if not then I will check it out tomorrow.
//current code sprint goal is to just try out the method above and then move onto other taks
//the problem I will face is that now the code only reflects the index of the task added in the current session

//It worked yes ir needs a lot of fine tuning but it worked
//problems I am facing right now
//how do I integrate the code with the timer
//this is why I was overwhelmed
//what to do
//so, first task, access the decodeJson slice and use its time.Time?
//how does time.Time actually work in simple language
//how to use time.Time to send a remeinder
//iterate through a slice of time.Time()
//time.Until(slice[i])
// when time.Now() == time.Until(slice[i])
//send reminder
