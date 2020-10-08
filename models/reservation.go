package models

import (
	"time"
)

type Reservation struct {
	// ID        primitive.ObjectID `json:"_id,omitempty"` // TODO: change to ObjectID ???
	ID        string         `json:"id,omitempty"`
	Dates     []time.Time    `json:"dates"`
	Timestamp time.Time      `json:"timestamp,omitempty"` // TODO: get from http request?
	Email     string         `json:"email"`
	Items     map[string]int `json:"items"`
}

type ReservationManager struct {
	ReservedItems map[time.Time]map[string]int
	Reservations  map[string]Reservation
}

func NewReservationManager() *ReservationManager {
	return &ReservationManager{
		ReservedItems: make(map[time.Time]map[string]int),
		Reservations:  make(map[string]Reservation),
	}
}

func (rm ReservationManager) CreateReservation(r Reservation) bool {
	// check if there are enough items in the stock
	for _, date := range r.Dates {
		rsvd := rm.ReservedItems[date]
		for id, amount := range r.Items {
			total := GetProduct(id).P.Amount
			if rsvd[id]+amount > int(total) {
				return false
			}
		}
	}
	// remove amount from items
	for _, date := range r.Dates {
		rsvd, ok := rm.ReservedItems[date]
		if !ok {
			rm.ReservedItems[date] = make(map[string]int)
			rsvd = rm.ReservedItems[date]
		}
		for id, amount := range r.Items {
			_, ok := rsvd[id]
			if !ok {
				rsvd[id] = 0
			}
			rsvd[id] += amount
		}
	}
	rm.Reservations[r.ID] = r
	return true
}

func (rm ReservationManager) DeleteReservation(id string) {
	r, ok := rm.Reservations[id]
	if ok {
		for _, date := range r.Dates {
			rsvd := rm.ReservedItems[date]
			for id, amount := range r.Items {
				rsvd[id] -= amount
			}
		}
		delete(rm.Reservations, id)
	}
}
