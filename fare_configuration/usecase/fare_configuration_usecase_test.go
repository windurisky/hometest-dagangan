package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/windurisky/hometest-dagangan/domain"
	"github.com/windurisky/hometest-dagangan/domain/mocks"
	_fareConfigUsecase "github.com/windurisky/hometest-dagangan/fare_configuration/usecase"
)

func TestFindByMileage(t *testing.T) {
	defaultDomainConfiguration := []domain.FareConfiguration{
		{
			UpperLimit:     1,
			FarePerMileage: 1000,
		},
		{
			UpperLimit:     10,
			FarePerMileage: 1500,
		},
		{
			UpperLimit:     100,
			FarePerMileage: 2000,
		},
	}
	testCases := []struct {
		name                                 string
		mileage                              uint64
		expectedResult                       domain.FareConfiguration
		expectedError                        error
		mockLoggerExpectation                func(mockLogger *mocks.Logger)
		mockFareConfigurationRepoExpectation func(mockFareConfigurationRepo *mocks.FareConfigurationRepository)
	}{
		{
			name:    "Valid Mileage",
			mileage: 10,
			expectedResult: domain.FareConfiguration{
				UpperLimit:     10,
				FarePerMileage: 1500,
			},
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {},
			mockFareConfigurationRepoExpectation: func(mockFareConfigurationRepo *mocks.FareConfigurationRepository) {
				mockFareConfigurationRepo.On("GetAll", mock.Anything).Return(defaultDomainConfiguration, nil)
			},
		},
		{
			name:    "Fail GetAll Fare Config",
			mileage: 10,
			expectedResult: domain.FareConfiguration{
				UpperLimit:     10,
				FarePerMileage: 1500,
			},
			expectedError: errors.New("Something is wrong"),
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", "Something is wrong").Return()
			},
			mockFareConfigurationRepoExpectation: func(mockFareConfigurationRepo *mocks.FareConfigurationRepository) {
				mockFareConfigurationRepo.On("GetAll", mock.Anything).Return(defaultDomainConfiguration, errors.New("Something is wrong"))
			},
		},
		{
			name:           "Fail GetAll Fare Config",
			mileage:        120,
			expectedResult: domain.FareConfiguration{},
			expectedError:  domain.ErrFareConfigurationNotFound,
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", domain.ErrFareConfigurationNotFound.Error()).Return()
			},
			mockFareConfigurationRepoExpectation: func(mockFareConfigurationRepo *mocks.FareConfigurationRepository) {
				mockFareConfigurationRepo.On("GetAll", mock.Anything).Return(defaultDomainConfiguration, nil)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockLogger := new(mocks.Logger)
			mockFareConfigurationRepo := new(mocks.FareConfigurationRepository)

			usecase := _fareConfigUsecase.NewFareConfigurationUsecase(mockLogger, mockFareConfigurationRepo)

			// set up mock expectations
			tc.mockLoggerExpectation(mockLogger)
			tc.mockFareConfigurationRepoExpectation(mockFareConfigurationRepo)

			// execute the function
			result, err := usecase.FindByMileage(tc.mileage)

			// assert the error result
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}
