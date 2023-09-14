package usecase

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/windurisky/hometest-dagangan/domain"
)

type tripUsecase struct {
	fareConfigurationUsecase domain.FareConfigurationUsecase
}

func NewTripUsecase(fc domain.FareConfigurationUsecase) domain.TripUsecase {
	return &tripUsecase{
		fareConfigurationUsecase: fc,
	}
}

func (t *tripUsecase) stringToDuration(input string) (result time.Duration, err error) {
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
		return
	}

	milliseconds, err := strconv.Atoi(secondsParts[1])
	if err != nil {
		return
	}

	result = time.Duration(hours)*time.Hour +
		time.Duration(minutes)*time.Minute +
		time.Duration(seconds)*time.Second +
		time.Duration(milliseconds)*time.Millisecond

	return
}

func (t *tripUsecase) ParseInput(input string) (result domain.Trip, err error) {
	values := strings.Split(input, " ")

	if len(values) != 3 {
		err = errors.New("should provide 3 parameters")
		return
	}

	duration, err := t.stringToDuration(values[1])
	if err != nil {
		return
	}

	result.Location = values[0]
	result.Duration = duration
	result.Mileage, err = strconv.ParseUint(values[2], 10, 64)

	return
}

func (t *tripUsecase) CalculateFare(mileage uint64) (result uint64, err error) {
	if mileage == 0 {
		return 0, err
	}

	fareConfig, err := t.fareConfigurationUsecase.FindByMileage(mileage)
	if err != nil {
		return
	}

	result = fareConfig.FarePerMileage * mileage

	return
}
