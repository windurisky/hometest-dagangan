package usecase

import (
	"github.com/windurisky/hometest-dagangan/domain"
	"github.com/windurisky/hometest-dagangan/logger"
)

type fareConfigurationUsecase struct {
	logger                logger.Logger
	fareConfigurationRepo domain.FareConfigurationRepository
}

func NewFareConfigurationUsecase(logger logger.Logger, fcRepo domain.FareConfigurationRepository) domain.FareConfigurationUsecase {
	return &fareConfigurationUsecase{
		logger:                logger,
		fareConfigurationRepo: fcRepo,
	}
}

func (fc *fareConfigurationUsecase) FindByMileage(mileage uint64) (result domain.FareConfiguration, err error) {
	relativePath := "fare_configuration/fixtures/fare_configuration.json"

	fareConfigurations, err := fc.fareConfigurationRepo.GetAll(relativePath)
	if err != nil {
		fc.logger.Error(err.Error())
		return
	}

	for _, config := range fareConfigurations {
		if config.UpperLimit >= mileage {
			result = config
			return
		}
	}

	if result == (domain.FareConfiguration{}) {
		err = domain.ErrFareConfigurationNotFound
		fc.logger.Error(err.Error())
	}
	return
}
