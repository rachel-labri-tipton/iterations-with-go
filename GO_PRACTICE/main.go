package main

import (
	"fmt"
	"strings"
)

//package level variables available to main function and all other functions in the package
//can also be accesed in all files in the package 
const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}


func main() {
	
	greetUsers()
	firstNames := getFirstNames(bookings)
	fmt.Printf("These people have already signed up for the conference: %v \n", firstNames)

	for {
		fistName, lastName, email, userTickets := getUserInput()
		isValideName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)
	}
	
	}
}



func greetUsers() {
			fmt.Printf("Welcome to %v booking application.", conferenceName)
  		fmt.Printf("We have a total of %v tickets available. \n", remainingTickets)
			fmt.Printf("Get your tickets here to attend the %v! \n", conferenceName)
	}

//this function takes a slice of strings and returns a slice of strings
//go functions have input and output parameters and here we specify the output type
func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
	var names = strings.Fields(booking)
	firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput() (bool, bool, bool) {
	isValidName := len(firstName)>= 2 && len(lastName)>= 2 
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName && isValidEmail && isValidTicketNumber
}

func getUserInput () (string, string, string, uint) {
	for {

	//declare variables
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Printf("Please enter your first name: \n")
		fmt.Scan(&firstName)

		fmt.Printf("Please enter your last name:  \n")
		fmt.Scan(&lastName)

		fmt.Printf("Please enter your email:  \n")
		fmt.Scan(&email)

		fmt.Printf("How many tickets do you want to buy? \n")
		fmt.Scan(&userTickets)

		// it's implicit that this is a boolean type
		// := is a short variable declaration, quickly declaring and initializing a variable
		isValidname, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
		


		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets for the %v conference. An email will be sent to %v with the details of your order. \n", firstName, lastName, userTickets, conferenceName, email)
			fmt.Printf("There are %v tickets remaining. \n", remainingTickets)
	
	//_ is a blank identifier. it's used to tell the compiler that we don't need this value. 
	// with Go you need to make unused varialbes explicit
			

			if remainingTickets ==0 {
				fmt.Println("Our conference is completely booked. Please try to register next year. Thank you.")
				break
			}
			
		} else {
			if !isValidName {
				fmt.Println("Please enter a valid name that's more than 2 letters. \n")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email with an @ sign. \n")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets must be more than 0. \n")
			}
	}
}

func bookTickets() (string, string, string, uint){
	remainingTickets = remainingTickets - userTickets
	return firstName, lastName, email, userTickets
}

}

// go only has for loops. no while loops. to stop the loop, you need to use break, which is a little like the case switch statement in JS
// continue is another instruction that can be used to skip the current iteration of the loop and move to the next one
// if x condition is true, then x logic should be executed. if x condition is false, then y logic should be executed
// an application should be able to handle bad input and unexpected values
// in the real world, we'd have a UI and a bookings would be persisted in a database
// languages provide various control structures to handle the flow of an application
// A loop statement allows us to execute code multiple times
//strings.Fields() splits the string with white space as a separator and returns a slice with the split elements
	//ask user for their name
	//pointer => value saved in a variable. pointers are used in the C and C++ languages.
//this shows where the memory address is for the variable 
//we need to declare the entry point of the program. where to start execute the code. go compiler
//needs an entry point to start the execution of the program. the entry point is the main function
//packages have functions that can be used. packages=functions of various functionalities
//need to look up in documentation to know the package you need to look up. it's all in the docs
//each data type can de different htings and behaves differently
//there are integer data types and float types that are specific and even strict

			

		