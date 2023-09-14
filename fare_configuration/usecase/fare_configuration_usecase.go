package usecase

import (
	"errors"
	"path/filepath"

	"github.com/windurisky/hometest-dagangan/domain"
)

type fareConfigurationUsecase struct {
	fareConfigurationRepo domain.FareConfigurationRepository
}

func NewFareConfigurationUsecase(fcRepo domain.FareConfigurationRepository) domain.FareConfigurationUsecase {
	return &fareConfigurationUsecase{
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
			break
		}
	}

	if result == (domain.FareConfiguration{}) {
		// TODO: error message library
		err = errors.New("fare fonfiguration not found")
	}
	return
}
