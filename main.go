package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	FirstName string
	Email     string
}

type Ticket struct {
	User
	Id       uuid.UUID
	Date     time.Time
	Quantity int
	Total    int
}

func main() {
	tickets := loadTickets()
	totalTickets := 50

	greetUsers()
	totalTickets, tickets = bookTickets(totalTickets, tickets)
	err := saveTickets(tickets)
	if err != nil {
		return
	}
	fmt.Printf("test? %v and %v", totalTickets, tickets)
}
