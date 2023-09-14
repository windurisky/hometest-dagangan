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
		return 0, err
	}

	fareConfig, err := t.fareConfigurationUsecase.FindByMileage(mileage)
	if err != nil {
		return
	}

	result = fareConfig.FarePerMileage * mileage

	return
}
