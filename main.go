package main

import "fmt"

type User struct {
	FirstName   string
	UserTickets int
	Email       string
}

func main() {

	bookingNames := make([]User, 0)
	totalTickets := 50

	greetUsers()
	totalTickets, bookingNames = bookAndSendTickets(totalTickets, bookingNames)
	fmt.Printf("test? %v and %v", totalTickets, bookingNames)

}
