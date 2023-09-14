package main

import (
	"fmt"
	"time"

	_fareConfigurationRepository "github.com/windurisky/hometest-dagangan/fare_configuration/repository"
	_fareConfigurationUsecase "github.com/windurisky/hometest-dagangan/fare_configuration/usecase"
	_tripUsecase "github.com/windurisky/hometest-dagangan/trip/usecase"
)

func main() {
	// TODO: make it into user input assigned values, create a input parser as well

	sampleTrips := []string{
		"PointA 01:20:05.500 2",
		"PointB 00:58:05.500 11",
		"PointC 00:25:00.000 0",
	}

	fareConfigurationRepo := _fareConfigurationRepository.NewFareConfigurationRepository()
	fareConfigurationUsecase := _fareConfigurationUsecase.NewFareConfigurationUsecase(fareConfigurationRepo)
	tripUsecase := _tripUsecase.NewTripUsecase(fareConfigurationUsecase)

	var totalMileAge uint64 = 0
	var totalFareAmount uint64 = 0
	var totalDuration time.Duration = time.Duration(0)

	for _, tripString := range sampleTrips {
		trip, err := tripUsecase.ParseInput(tripString)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		fareAmount, err := tripUsecase.CalculateFare(trip.Mileage)
		if err != nil {
			// TODO: add logging here
			fmt.Println("Error:", err)
			break
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
