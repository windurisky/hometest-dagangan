package domain

import "time"

type Trip struct {
	Location   string
	Duration   time.Duration
	Mileage    int64
	FareAmount *int64
}

type TripUsecase interface {
	CalculateFare(trip *Trip) error
}
