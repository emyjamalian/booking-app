package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

func getUserInfo() (Ticket, error) {
	var firstName string
	var quantity int
	var email string

	fmt.Println("Please enter your name:")
	_, err := fmt.Scanln(&firstName)
	if err != nil {
		return Ticket{}, err
	}
	if strings.TrimSpace(firstName) == "" {
		return Ticket{}, errors.New("name cannot be empty")
	}

	fmt.Println("Please enter your email address:")
	_, err = fmt.Scanln(&email)
	if err != nil {
		return Ticket{}, err
	}
	if err := checkmail.ValidateFormat(email); err != nil {
		return Ticket{}, fmt.Errorf("invalid email: %q", email)
	}

	fmt.Println("Please enter the number of tickets you want to book:")
	_, err = fmt.Scanln(&quantity)
	if err != nil {
		return Ticket{}, err
	}

	return Ticket{
		Id:       uuid.New(),
		Quantity: quantity,
		Date:     time.Now(),
		User:     User{FirstName: firstName, Email: email}}, nil
}

func greetUsers() {
	fmt.Println("Welcome to Devfest booking!")
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

func bookTickets(totalTickets int, tickets []Ticket) (int, []Ticket) {

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
		tickets = append(tickets, userDetails)

		fmt.Printf("Thank you %s for booking %v tickets for Devfest!\n", userDetails.FirstName, userDetails.Quantity)
	}
	return totalTickets, tickets
}
