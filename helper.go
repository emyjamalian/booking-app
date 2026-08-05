package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func getUserInfo() (User, error) {
	var firstName string
	var userTickets int
	var email string

	fmt.Println("Please enter your name:")
	fmt.Scanln(&firstName)
	if strings.TrimSpace(firstName) == "" {
		return User{}, errors.New("name cannot be empty")
	}

	fmt.Println("Please enter your email address:")
	fmt.Scanln(&email)
	if !strings.Contains(email, "@") {
		return User{}, fmt.Errorf("invalid email: %q", email)
	}

	fmt.Println("Please enter the number of tickets you want to book:")
	fmt.Scanln(&userTickets)

	return User{FirstName: firstName, UserTickets: userTickets, Email: email}, nil
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

func isValidTicket(totalTickets int, userTickets int) error {
	if userTickets <= 0 {
		return errors.New("tickets must be greater than zero")
	}
	if userTickets > totalTickets {
		return fmt.Errorf("only %d tickets remaining", totalTickets)
	}
	return nil
}

func bookAndSendTickets(totalTickets int, bookingNames []User) (int, []User) {

	for totalTickets > 0 {
		userDetails, err := getUserInfo()
		if err != nil {
			fmt.Println(color.RedString("Invalid input:"), err)
			continue
		}

		if err := isValidTicket(totalTickets, userDetails.UserTickets); err != nil {
			fmt.Println(color.RedString("Booking failed:"), err)
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
