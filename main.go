package main

import "fmt"

type User struct {
	FirstName string
	Email     string
}

type Ticket struct {
	User
	Quantity int
	Total    int
}

func main() {

	bookingNames := make([]User, 0) // to be in json file
	totalTickets := 50

	greetUsers()
	totalTickets, bookingNames = bookTickets(totalTickets, bookingNames)
	fmt.Printf("test? %v and %v", totalTickets, bookingNames)

}
