package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

const fileName = "tickets.json"

type User struct {
	FirstName string `json:"firstName"`
	Email     string `json:"email"`
}

type Ticket struct {
	User
	Id       uuid.UUID `json:"id"`
	Date     time.Time `json:"date"`
	Quantity int       `json:"quantity"`
	Total    int       `json:"total"`
}

type TicketStorage struct{}

func (s *TicketStorage) FindAll() ([]Ticket, error) {
	return loadTickets(), nil
}

func (s *TicketStorage) FindByID(id uuid.UUID) (Ticket, error) {
	tickets := loadTickets()
	for _, t := range tickets {
		if t.Id == id {
			return t, nil
		}
	}
	return Ticket{}, fmt.Errorf("ticket not found: %s", id)
}

func (s *TicketStorage) Create(t Ticket) (Ticket, error) {
	tickets := loadTickets()
	tickets = append(tickets, t)
	return t, saveTickets(tickets)
}

func (s *TicketStorage) Update(t Ticket) (Ticket, error) {
	tickets := loadTickets()
	for i, existing := range tickets {
		if existing.Id == t.Id {
			tickets[i] = t
			return t, saveTickets(tickets)
		}
	}
	return Ticket{}, fmt.Errorf("ticket not found: %s", t.Id)
}

func (s *TicketStorage) Delete(id uuid.UUID) error {
	tickets := loadTickets()
	for i, t := range tickets {
		if t.Id == id {
			tickets = append(tickets[:i], tickets[i+1:]...)
			return saveTickets(tickets)
		}
	}
	return fmt.Errorf("ticket not found: %s", id)
}

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
		return fmt.Errorf("error: %q", err)
	}
	err = os.WriteFile(fileName, data, 0644)
	return err
}
