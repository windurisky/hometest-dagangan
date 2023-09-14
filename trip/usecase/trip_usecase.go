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

func (t *tripUsecase) CalculateFare(mileage uint64) (result uint64, err error) {
	if mileage == 0 {
		result = 0
		return
	}

	fareConfig, err := t.fareConfigurationUsecase.FindByMileage(mileage)
	if err != nil {
		// TODO: add logging here
		return
	}

	result = fareConfig.FarePerMileage * mileage

	return
}
