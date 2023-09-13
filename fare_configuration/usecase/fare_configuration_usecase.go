package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/windurisky/hometest-dagangan/domain"
)

type fareConfigurationUsecase struct{}

func NewFareConfigurationUsecase() domain.FareConfigurationUsecase {
	return &fareConfigurationUsecase{}
}

func (fc *fareConfigurationUsecase) GetList(fixturePath string) (result []domain.FareConfiguration, err error) {
	// on real life use cases, this will probably be in a database
	// in that case, it will call something like fareConfigurationRepository.GetList()
	// for demo purpose, changing the json fixture as needed will suffice
	file, err := os.Open(fixturePath)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&result); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	return
}

func (fc *fareConfigurationUsecase) FindByMileage(mileage int64) (result domain.FareConfiguration, err error) {
	relativePath := "fare_configuration/fixtures/fare_configuration.json"
	absolutePath, _ := filepath.Abs(relativePath)
	fareConfigurations, err := fc.GetList(absolutePath)

	for _, config := range fareConfigurations {
		if config.UpperLimit >= mileage {
			result = config
			break
		}
	}

	if result == (domain.FareConfiguration{}) {
		err = errors.New("fare fonfiguration not found")
	}
	return
}
