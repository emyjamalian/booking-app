package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func getUserInfo() (Ticket, error) {
	var firstName string
	var quantity int
	var email string

	fmt.Println("Please enter your name:")
	fmt.Scanln(&firstName)
	if strings.TrimSpace(firstName) == "" {
		return Ticket{}, errors.New("name cannot be empty")
	}

	fmt.Println("Please enter your email address:")
	fmt.Scanln(&email)
	if !strings.Contains(email, "@") {
		return Ticket{}, fmt.Errorf("invalid email: %q", email)
	}

	fmt.Println("Please enter the number of tickets you want to book:")
	fmt.Scanln(&quantity)

	return Ticket{Quantity: quantity, User: User{FirstName: firstName, Email: email}}, nil
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

func isValidTicket(totalTickets int, quantity int) error {
	if quantity <= 0 {
		return errors.New("tickets must be greater than zero")
	}
	if quantity > totalTickets {
		return fmt.Errorf("only %d tickets remaining", totalTickets)
	}
	return nil
}

func bookTickets(totalTickets int, bookingNames []User) (int, []User) {

	for totalTickets > 0 {
		userDetails, err := getUserInfo()
		if err != nil {
			fmt.Println(color.RedString("Invalid input:"), err)
			continue
		}

		if err := isValidTicket(totalTickets, userDetails.Quantity); err != nil {
			fmt.Println(color.RedString("Booking failed:"), err)
			continue
		}

		totalTickets -= userDetails.Quantity
		bookingNames = append(bookingNames, userDetails.User)

		fmt.Printf("Thank you %s for booking %v tickets for Devfest!\n", userDetails.FirstName, userDetails.Quantity)

		fmt.Printf("\n*******************************\n")
		fmt.Printf("List of all attendees: %v\n", strings.Join(namesOnly(bookingNames), ", "))
		fmt.Printf("\n*******************************\n")
	}
	return totalTickets, bookingNames
}
