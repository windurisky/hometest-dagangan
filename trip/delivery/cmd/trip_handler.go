package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/windurisky/hometest-dagangan/domain"
)

type tripHandler struct {
	tripUsecase domain.TripUsecase
}

func NewTripHandler(tripUsecase domain.TripUsecase) domain.TripHandler {
	return &tripHandler{
		tripUsecase: tripUsecase,
	}
}

func (t *tripHandler) stringToDuration(input string) (result time.Duration, err error) {
	// input string format must be hh:mm:ss.fff
	durationParts := strings.Split(input, ":")
	if len(durationParts) != 3 {
		err = errors.New("duration format must be hh:mm:ss.fff")
		return
	}

	hours, err := strconv.Atoi(durationParts[0])
	if err != nil {
		return
	}

	minutes, err := strconv.Atoi(durationParts[1])
	if err != nil {
		return
	}

	secondsParts := strings.Split(durationParts[2], ".")
	if len(secondsParts) != 2 {
		// TODO: redundant with above error, should use custom error library
		err = errors.New("duration format must be hh:mm:ss.fff")
		return
	}

	seconds, err := strconv.Atoi(secondsParts[0])
	if err != nil {
		// TODO: add logging here
		return
	}

	milliseconds, err := strconv.Atoi(secondsParts[1])
	if err != nil {
		// TODO: add logging here
		return
	}

	result = time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		time.Duration(milliseconds)*time.Millisecond

	// TODO: put in env
	lowerLimit := 2 * time.Minute
	upperLimit := 10 * time.Minute

	if result < lowerLimit || result > upperLimit {

		err = errors.New("duration must be between 2 to 10 minutes")
		// TODO: add logging here
	}
	return
}

func (t *tripHandler) durationToString(input time.Duration) (result string) {
	hours := input / time.Hour
	input -= hours * time.Hour

	minutes := input / time.Minute
	input -= minutes * time.Minute

	seconds := input / time.Second
	input -= seconds * time.Second

	milliseconds := input / time.Millisecond

	return fmt.Sprintf("%02d:%02d:%02d.%03d", hours, minutes, seconds, milliseconds)
}

func (t *tripHandler) ParseInput(input string) (result domain.Trip, err error) {
	values := strings.Split(input, " ")

	if len(values) != 3 {
		// TODO: add to custom error
		// TODO: add logging here
		err = errors.New("should provide 3 parameters")
		return
	}

	duration, err := t.stringToDuration(values[1])
	if err != nil {
		// TODO: add logging here
		return
	}

	result.Location = values[0]
	result.Duration = duration
	result.Mileage, err = strconv.ParseUint(values[2], 10, 64)

	return
}

func (t *tripHandler) SummarizeTrip(trips []domain.Trip) (err error) {
	var totalMileAge uint64 = 0
	var totalFareAmount uint64 = 0
	var totalDuration time.Duration = time.Duration(0)

	for _, trip := range trips {
		fareAmount, err := t.tripUsecase.CalculateFare(trip.Mileage)
		if err != nil {
			// TODO: add logging here
			return err
		}

		trip.FareAmount = &fareAmount
		fmt.Println(trip.Location, t.durationToString(trip.Duration), trip.Mileage, *trip.FareAmount)

		totalMileAge += trip.Mileage
		totalFareAmount += fareAmount
		totalDuration += trip.Duration
	}

	if totalMileAge == 0 {
		err = errors.New("total mileage must be greater than zero")
		// TODO: add logging here
		return
	}

	fmt.Println("Total Mileage:", totalMileAge, "km")
	fmt.Println("Total Fare Amount: Rp", totalFareAmount)
	fmt.Println("Total Duration:", t.durationToString(totalDuration))
	return
}
