package usecase

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/windurisky/hometest-dagangan/domain"
)

type fareConfigurationUsecase struct{}

func NewFareConfigurationUsecase() domain.FareConfigurationUsecase {
	return &fareConfigurationUsecase{}
}

func (fc *fareConfigurationUsecase) GetList(fixturePath string) (result []domain.FareConfiguration, err error) {
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
