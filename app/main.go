package main

import (
	"fmt"
	"time"

	"github.com/windurisky/hometest-dagangan/domain"
	_fareConfigurationUsecase "github.com/windurisky/hometest-dagangan/fare_configuration/usecase"
	_tripUsecase "github.com/windurisky/hometest-dagangan/trip/usecase"
)

func main() {
	// TODO: make it into user input assigned values, create a input parser as well
	sampleTrips := []domain.Trip{
		{
			Location: "PointA",
			Duration: time.Duration(80) * time.Minute,
			Mileage:  2,
		},
		{
			Location: "PointB",
			Duration: time.Duration(3500) * time.Second,
			Mileage:  11,
		},
		{
			Location: "Pointc",
			Duration: time.Duration(1500) * time.Second,
			Mileage:  0,
		},
	}

	fareConfigurationUsecase := _fareConfigurationUsecase.NewFareConfigurationUsecase()
	tripUsecase := _tripUsecase.NewTripUsecase(fareConfigurationUsecase)
	var totalMileAge uint64 = 0
	var totalFareAmount uint64 = 0
	var totalDuration time.Duration = time.Duration(0)

	for _, trip := range sampleTrips {
		fareAmount, err := tripUsecase.CalculateFare(trip.Mileage)
		if err != nil {
			// TODO: add logging here
			fmt.Println("Error:", err)
		}
		trip.FareAmount = &fareAmount
		fmt.Println(trip.Location, trip.Duration, trip.Mileage, *trip.FareAmount)

		totalMileAge += trip.Mileage
		totalFareAmount += fareAmount
		totalDuration += trip.Duration
	}
	fmt.Println("Total Mileage:", totalMileAge)
	fmt.Println("Total Fare Amount:", totalFareAmount)
	fmt.Println("Total Duration", totalDuration)
}
