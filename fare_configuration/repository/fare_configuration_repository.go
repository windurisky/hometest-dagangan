package repository

import (
	"encoding/json"
	"os"

	"github.com/windurisky/hometest-dagangan/domain"
	"github.com/windurisky/hometest-dagangan/logger"
)

type fareConfigurationRepository struct {
	logger logger.Logger
}

func NewFareConfigurationRepository(logger logger.Logger) domain.FareConfigurationRepository {
	return &fareConfigurationRepository{
		logger: logger,
	}
}

func (fc *fareConfigurationRepository) GetAll(fixturePath string) (result []domain.FareConfiguration, err error) {
	// on real life use cases, the data will most likely be retrieved from a database
	// for demo purpose, changing the json fixture as needed will suffice
	file, err := os.Open(fixturePath)
	if err != nil {
		fc.logger.Error(err.Error())
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&result); err != nil {
		fc.logger.Error(err.Error())
		return
	}

	return
}
