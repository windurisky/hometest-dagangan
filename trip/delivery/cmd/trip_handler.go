package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/windurisky/hometest-dagangan/domain"
	"github.com/windurisky/hometest-dagangan/logger"
	"github.com/windurisky/hometest-dagangan/util"
)

type tripHandler struct {
	logger      logger.Logger
	tripUsecase domain.TripUsecase
}

func NewTripHandler(logger logger.Logger, tripUsecase domain.TripUsecase) domain.TripHandler {
	return &tripHandler{
		logger:      logger,
		tripUsecase: tripUsecase,
	}
}

// stringToDuration will convert string with hh:mm:ss.fff format into time.Duration
func (t *tripHandler) stringToDuration(input string) (result time.Duration, err error) {
	durationParts := strings.Split(input, ":")
	if len(durationParts) != 3 {
		err = domain.ErrInvalidDurationFormat
		t.logger.Error(err.Error())
		return
	}

	hours, err := strconv.Atoi(durationParts[0])
	if err != nil {
		t.logger.Error(err.Error())
		return
	}

	minutes, err := strconv.Atoi(durationParts[1])
	if err != nil {
		t.logger.Error(err.Error())
		return
	}
	if minutes > 59 {
		err = domain.ErrInvalidMinutes
		t.logger.Error(err.Error())
		return
	}

	secondsParts := strings.Split(durationParts[2], ".")
	if len(secondsParts) != 2 {
		err = domain.ErrInvalidDurationFormat
		t.logger.Error(err.Error())
		return
	}

	seconds, err := strconv.Atoi(secondsParts[0])
	if err != nil {
		t.logger.Error(err.Error())
		return
	}
	if seconds > 59 {
		err = domain.ErrInvalidSeconds
		t.logger.Error(err.Error())
		return
	}

	milliseconds, err := strconv.Atoi(secondsParts[1])
	if err != nil {
		t.logger.Error(err.Error())
		return
	}
	if milliseconds > 999 {
		err = domain.ErrInvalidMilliseconds
		t.logger.Error(err.Error())
		return
	}

	result = time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		time.Duration(milliseconds)*time.Millisecond

	lowerLimit, err := strconv.Atoi(util.GetEnvWithDefault("DURATION_LOWER_LIMIT_MINUTES", "2"))
	if err != nil {
		t.logger.Error(err.Error())
		return
	}

	lowerLimitDuration := time.Duration(lowerLimit) * time.Minute

	upperLimit, err := strconv.Atoi(util.GetEnvWithDefault("DURATION_UPPER_LIMIT_MINUTES", "10"))
	if err != nil {
		t.logger.Error(err.Error())
		return
	}

	upperLimitDuration := time.Duration(upperLimit) * time.Minute

	if result < lowerLimitDuration || result > upperLimitDuration {
		err = domain.ErrInvalidDurationRange
		t.logger.Error(err.Error())
	}
	return
}

// durationToString converts time.Duration into hh:mm:ss.fff string format
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
		err = domain.ErrInvalidTripParameterCount
		t.logger.Error(err.Error())
		return
	}

	duration, err := t.stringToDuration(values[1])
	if err != nil {
		return
	}

	result.Location = values[0]
	result.Duration = duration
	result.Mileage, err = strconv.ParseUint(values[2], 10, 64)

	logFields := map[string]interface{}{
		"input":  input,
		"result": result,
	}
	t.logger.Info("Successfully parsed input", logFields)

	return
}

func (t *tripHandler) SummarizeTrip(trips []domain.Trip) (err error) {
	var totalMileAge uint64 = 0
	var totalFareAmount uint64 = 0
	var totalDuration time.Duration = time.Duration(0)

	for _, trip := range trips {
		fareAmount, err := t.tripUsecase.CalculateFare(trip.Mileage)
		if err != nil {
			return err
		}

		trip.FareAmount = &fareAmount
		fmt.Println(trip.Location, t.durationToString(trip.Duration), trip.Mileage, *trip.FareAmount)

		totalMileAge += trip.Mileage
		totalFareAmount += fareAmount
		totalDuration += trip.Duration
	}

	if totalMileAge == 0 {
		err = domain.ErrInvalidTotalMileage
		t.logger.Error(err.Error())
		return
	}

	fmt.Println("Total Mileage:", totalMileAge, "km")
	fmt.Println("Total Fare Amount: Rp", totalFareAmount)
	fmt.Println("Total Duration:", t.durationToString(totalDuration))

	logFields := map[string]interface{}{"input": trips}
	t.logger.Info("Successfully summarized trips", logFields)

	return
}
