package main

import (
	"fmt"
	"path/filepath"

	_fareConfigurationUsecase "github.com/windurisky/hometest-dagangan/fare_configuration/usecase"
)

func main() {
	relativePath := "fare_configuration/fixtures/fare_configuration.json"
	absolutePath, _ := filepath.Abs(relativePath)
	fmt.Println(_fareConfigurationUsecase.NewFareConfigurationUsecase().GetList(absolutePath))
}
