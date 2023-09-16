package domain

import "errors"

var (
	ErrFareConfigurationNotFound = errors.New("fare fonfiguration not found")
	ErrInvalidDurationFormat     = errors.New("duration format must be hh:mm:ss.fff")
	ErrInvalidDurationRange      = errors.New("duration must be ranged between 2 to 10 minutes")
	ErrInvalidMinutes            = errors.New("minutes must be ranged between 0 to 59")
	ErrInvalidSeconds            = errors.New("seconds must be ranged between 0 to 59")
	ErrInvalidMilliseconds       = errors.New("milliseconds must be ranged between 0 to 999")
	ErrInvalidTripParameterCount = errors.New("should input 3 parameters on trip")
	ErrInvalidTotalMileage       = errors.New("total mileage must be greater than zero")
)
