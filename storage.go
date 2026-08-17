package main

import (
	"encoding/json"
	"os"
)

const fileName = "tickets.json"

func loadTickets() []Ticket {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return []Ticket{}
	}
	var tickets []Ticket
	if err = json.Unmarshal(data, &tickets); err != nil {
	}
	return tickets
}

func saveTickets(tickets []Ticket) error {
	data, err := json.Marshal(tickets)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileName, data, 0644)
	return err
}
