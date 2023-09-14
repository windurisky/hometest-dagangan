package main

import (
	"fmt"
	"time"

	"github.com/windurisky/hometest-dagangan/domain"
	_fareConfigurationUsecase "github.com/windurisky/hometest-dagangan/fare_configuration/usecase"
	_tripUsecase "github.com/windurisky/hometest-dagangan/trip/usecase"
)

func main() {
	sampleTrips := []domain.Trip{
		{
			Location: "PointA",
			Duration: time.Duration(80) * time.Minute,
			Mileage:  8,
		},
		{
			Location: "PointB",
			Duration: time.Duration(3500) * time.Second,
			Mileage:  11,
		},
	}

	fareConfigurationUsecase := _fareConfigurationUsecase.NewFareConfigurationUsecase()
	tripUsecase := _tripUsecase.NewTripUsecase(fareConfigurationUsecase)
	for _, trip := range sampleTrips {
		tripUsecase.CalculateFare(&trip)
		fmt.Println(trip.Location, trip.Duration, trip.Mileage, *trip.FareAmount)
	}
}
