package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/windurisky/hometest-dagangan/domain"
	_fareConfigurationRepository "github.com/windurisky/hometest-dagangan/fare_configuration/repository"
	_fareConfigurationUsecase "github.com/windurisky/hometest-dagangan/fare_configuration/usecase"
	"github.com/windurisky/hometest-dagangan/logger"
	_tripHandler "github.com/windurisky/hometest-dagangan/trip/delivery/cmd"
	_tripUsecase "github.com/windurisky/hometest-dagangan/trip/usecase"
)

func main() {
	logger, err := logger.NewZapLogger()
	if err != nil {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	fareConfigurationRepo := _fareConfigurationRepository.NewFareConfigurationRepository(logger)
	fareConfigurationUsecase := _fareConfigurationUsecase.NewFareConfigurationUsecase(logger, fareConfigurationRepo)
	tripUsecase := _tripUsecase.NewTripUsecase(logger, fareConfigurationUsecase)
	tripHandler := _tripHandler.NewTripHandler(logger, tripUsecase)

	var input string

	fmt.Print("Enter number of trips: ")
	_, err = fmt.Scanln(&input)
	if err != nil {
		logger.Error(err.Error())
		fmt.Println("Error reading input:", err)
		return
	}

	tripLength, err := strconv.Atoi(input)
	if err != nil {
		logger.Error(err.Error())
		fmt.Println("Error reading input:", err)
		return
	}

	var trips []domain.Trip
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < tripLength; i++ {
		fmt.Printf("Enter trip %d: ", i+1)
		if scanner.Scan() {
			input = scanner.Text()
		} else {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}

		trip, err := tripHandler.ParseInput(input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		trips = append(trips, trip)
	}

	err = tripHandler.SummarizeTrip(trips)
	if err != nil {
		fmt.Println("Error processing data:", err)
	}
}
