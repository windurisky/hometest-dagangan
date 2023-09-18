package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/windurisky/hometest-dagangan/domain"
	"github.com/windurisky/hometest-dagangan/domain/mocks"
	_tripUsecase "github.com/windurisky/hometest-dagangan/trip/usecase"
)

func TestCalculateFare(t *testing.T) {
	testCases := []struct {
		name                             string
		mileage                          uint64
		expectedFare                     uint64
		expectedError                    error
		mockLoggerExpectation            func(mockLogger *mocks.Logger)
		mockFareConfigUsecaseExpectation func(mockFareConfigUsecase *mocks.FareConfigurationUsecase)
	}{
		{
			name:          "Valid Mileage",
			mileage:       10,
			expectedFare:  20,
			expectedError: nil,
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Info", "Successfully calculated fare", mock.Anything).Return()
			},
			mockFareConfigUsecaseExpectation: func(mockFareConfigUsecase *mocks.FareConfigurationUsecase) {
				expectedFareConfig := domain.FareConfiguration{
					UpperLimit:     10,
					FarePerMileage: 2,
				}
				mockFareConfigUsecase.On("FindByMileage", uint64(10)).Return(expectedFareConfig, nil)
			},
		},
		{
			name:                             "Zero Mileage",
			mileage:                          0,
			expectedFare:                     0,
			expectedError:                    nil,
			mockLoggerExpectation:            func(mockLogger *mocks.Logger) {},
			mockFareConfigUsecaseExpectation: func(mockFareConfigUsecase *mocks.FareConfigurationUsecase) {},
		},
		{
			name:                  "FareConfig Error",
			mileage:               5,
			expectedFare:          0,
			expectedError:         domain.ErrFareConfigurationNotFound,
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {},
			mockFareConfigUsecaseExpectation: func(mockFareConfigUsecase *mocks.FareConfigurationUsecase) {
				mockFareConfigUsecase.On("FindByMileage", uint64(5)).Return(domain.FareConfiguration{}, domain.ErrFareConfigurationNotFound)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockLogger := new(mocks.Logger)
			mockFareConfigUsecase := new(mocks.FareConfigurationUsecase)
			tripUsecase := _tripUsecase.NewTripUsecase(mockLogger, mockFareConfigUsecase)

			// set up mock expectations
			tc.mockLoggerExpectation(mockLogger)
			tc.mockFareConfigUsecaseExpectation(mockFareConfigUsecase)

			// execute the function
			result, err := tripUsecase.CalculateFare(tc.mileage)

			// assert the results
			mockLogger.AssertExpectations(t)
			mockFareConfigUsecase.AssertExpectations(t)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedFare, result)
		})
	}
}
