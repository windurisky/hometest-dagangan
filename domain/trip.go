package domain

import "time"

type Trip struct {
	Location   string
	Duration   time.Duration
	Mileage    uint64
	FareAmount *uint64
}

type TripUsecase interface {
	ParseInput(string) (Trip, error)
	CalculateFare(uint64) (uint64, error)
}
