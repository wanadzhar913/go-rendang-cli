package rendang

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

const (
	DefaultEventName = "Rendang Factory"
	DefaultStock     = uint(100)
)

type Order struct {
	FirstName string
	LastName  string
	Email     string
	Packs     uint
}

type Summary struct {
	EventName      string
	TotalStock     uint
	RemainingStock uint
	OrderCount     int
	Bookings       []Order
}

type Store struct {
	eventName      string
	totalStock     uint
	remainingStock uint
	bookings       []Order
	mu             sync.Mutex
}

func NewStore(eventName string, totalStock uint) *Store {
	if strings.TrimSpace(eventName) == "" {
		eventName = DefaultEventName
	}
	if totalStock == 0 {
		totalStock = DefaultStock
	}

	return &Store{
		eventName:      eventName,
		totalStock:     totalStock,
		remainingStock: totalStock,
		bookings:       make([]Order, 0),
	}
}

func (s *Store) PlaceOrder(order Order) (Summary, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order = sanitizeOrder(order)
	issues := validateOrder(order, s.remainingStock)
	if len(issues) > 0 {
		return Summary{}, errors.New(strings.Join(issues, "\n"))
	}

	s.remainingStock -= order.Packs
	s.bookings = append(s.bookings, order)

	return s.summaryLocked(), nil
}

func (s *Store) Summary() Summary {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.summaryLocked()
}

func (s *Store) summaryLocked() Summary {
	bookings := make([]Order, len(s.bookings))
	copy(bookings, s.bookings)

	return Summary{
		EventName:      s.eventName,
		TotalStock:     s.totalStock,
		RemainingStock: s.remainingStock,
		OrderCount:     len(s.bookings),
		Bookings:       bookings,
	}
}

func sanitizeOrder(order Order) Order {
	order.FirstName = strings.TrimSpace(order.FirstName)
	order.LastName = strings.TrimSpace(order.LastName)
	order.Email = strings.TrimSpace(strings.ToLower(order.Email))
	return order
}

func validateOrder(order Order, remainingStock uint) []string {
	issues := make([]string, 0, 4)

	if len(order.FirstName) < 2 || len(order.LastName) < 2 {
		issues = append(issues, "- first and last names must be at least 2 characters")
	}

	if !strings.Contains(order.Email, "@") || strings.HasSuffix(order.Email, "@") {
		issues = append(issues, "- email must look like a real address")
	}

	if order.Packs == 0 {
		issues = append(issues, "- packs must be greater than 0")
	}

	if order.Packs > remainingStock {
		issues = append(issues, fmt.Sprintf("- only %d pack(s) remaining in stock", remainingStock))
	}

	return issues
}
