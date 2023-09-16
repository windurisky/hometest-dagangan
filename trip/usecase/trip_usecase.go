package usecase

import (
	"github.com/windurisky/hometest-dagangan/domain"
	"github.com/windurisky/hometest-dagangan/logger"
)

type tripUsecase struct {
	logger                   logger.Logger
	fareConfigurationUsecase domain.FareConfigurationUsecase
}

func NewTripUsecase(logger logger.Logger, fc domain.FareConfigurationUsecase) domain.TripUsecase {
	return &tripUsecase{
		logger:                   logger,
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
		return
	}

	result = fareConfig.FarePerMileage * mileage

	return
}
