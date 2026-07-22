package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func getUserInfo() User {
	var firstName string
	var userTickets int
	var email string
	//shows the question, take the user input and take that value with the pointer and put it in the variable
	fmt.Println("Please enter your name:")
	fmt.Scanln(&firstName)

	fmt.Println("Please enter your email address:")
	fmt.Scanln(&email)

	fmt.Println("Please enter the number of tickets you want to book:")
	fmt.Scanln(&userTickets)

	return User{FirstName: firstName, UserTickets: userTickets, Email: email}
}

func greetUsers() {
	fmt.Println("Welcome to Devfest booking!")
}

func namesOnly(users []User) []string {
	names := make([]string, len(users))
	for i, u := range users {
		names[i] = u.FirstName
	}
	return names
}

func isValidTicket(totalTickets int, userTickets int) bool {
	return userTickets > 0 && userTickets <= totalTickets
}

func bookAndSendTickets(totalTickets int, bookingNames []User) (int, []User) {

	for totalTickets > 0 {
		var userDetails = getUserInfo()

		if !isValidTicket(totalTickets, userDetails.UserTickets) {
			fmt.Println(color.RedString("Not enough tickets. Available:"), totalTickets)
			continue
		}

		totalTickets -= userDetails.UserTickets
		bookingNames = append(bookingNames, userDetails)

		fmt.Printf("Thank you %s for booking %v tickets for Devfest!\n", userDetails.FirstName, userDetails.UserTickets)

		fmt.Printf("\n*******************************\n")
		fmt.Printf("List of all attendees: %v\n", strings.Join(namesOnly(bookingNames), ", "))
		fmt.Printf("\n*******************************\n")

	}
	return totalTickets, bookingNames
}
