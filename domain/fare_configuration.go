package domain

type FareConfiguration struct {
	UpperLimit     int64 `json:"upper_limit"`
	FarePerMileage int64 `json:"fare_per_mileage"`
}

type FareConfigurationUsecase interface {
	GetList(string) ([]FareConfiguration, error)
}
