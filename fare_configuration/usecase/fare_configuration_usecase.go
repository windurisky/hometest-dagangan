package usecase

import (
	"path/filepath"

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
	absolutePath, _ := filepath.Abs(relativePath)
	fareConfigurations, err := fc.fareConfigurationRepo.GetAll(absolutePath)
	if err != nil {
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
