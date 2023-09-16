package domain

import "errors"

var (
	ErrFareConfigurationNotFound = errors.New("fare fonfiguration not found")
	ErrInvalidDurationFormat     = errors.New("duration format must be hh:mm:ss.fff")
	ErrInvalidDurationRange      = errors.New("duration must be ranged between 2 to 10 minutes")
	ErrInvalidTripParameterCount = errors.New("should input 3 parameters on trip")
	ErrInvalidTotalMileage       = errors.New("total mileage must be greater than zero")
)
