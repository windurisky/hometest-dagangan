package domain

type FareConfiguration struct {
	UpperLimit     uint64 `json:"upper_limit"`
	FarePerMileage uint64 `json:"fare_per_mileage"`
}

type FareConfigurationUsecase interface {
	FindByMileage(uint64) (FareConfiguration, error)
}

type FareConfigurationRepository interface {
	GetAll(string) ([]FareConfiguration, error)
}
