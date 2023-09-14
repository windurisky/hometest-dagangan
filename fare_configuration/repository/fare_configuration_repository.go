package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/windurisky/hometest-dagangan/domain"
)

type fareConfigurationRepository struct{}

func NewFareConfigurationRepository() domain.FareConfigurationRepository {
	return &fareConfigurationRepository{}
}

func (fc *fareConfigurationRepository) GetAll(fixturePath string) (result []domain.FareConfiguration, err error) {
	// on real life use cases, the data will most likely be retrieved from a database
	// for demo purpose, changing the json fixture as needed will suffice
	file, err := os.Open(fixturePath)
	if err != nil {
		// TODO: change with logging
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&result); err != nil {
		// TODO: change with logging
		fmt.Println("Error decoding JSON:", err)
		return
	}

	return
}
