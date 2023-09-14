package domain

type FareConfiguration struct {
	UpperLimit     uint64 `json:"upper_limit"`
	FarePerMileage uint64 `json:"fare_per_mileage"`
}

type FareConfigurationUsecase interface {
	GetList(string) ([]FareConfiguration, error)
	FindByMileage(uint64) (FareConfiguration, error)
}
