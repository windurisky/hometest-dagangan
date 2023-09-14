package usecase

import (
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

func (t *tripUsecase) CalculateFare(trip *domain.Trip) (err error) {
	fareConfig, err := t.fareConfigurationUsecase.FindByMileage(trip.Mileage)
	if err != nil {
		return
	}

	fareAmount := fareConfig.FarePerMileage * trip.Mileage
	trip.FareAmount = &fareAmount

	return
}
