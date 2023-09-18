package repository_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/windurisky/hometest-dagangan/domain"
	"github.com/windurisky/hometest-dagangan/domain/mocks"
	_fareConfigRepo "github.com/windurisky/hometest-dagangan/fare_configuration/repository"
)

func TestGetAll(t *testing.T) {
	testCases := []struct {
		name                  string
		relativePath          string
		expectedResult        []domain.FareConfiguration
		expectedError         error
		mockLoggerExpectation func(mockLogger *mocks.Logger)
	}{
		{
			name:         "Valid Fixture Data",
			relativePath: "../fixtures/fare_configuration.json",
			expectedResult: []domain.FareConfiguration{
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
			},
			expectedError:         nil,
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {},
		},
		{
			name:           "Invalid Fixture File",
			relativePath:   "some/wrong/path.json",
			expectedResult: nil,
			expectedError:  errors.New("open some/wrong/path.json: no such file or directory"),
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", "open some/wrong/path.json: no such file or directory").Return()
			},
		},
		{
			name:           "JSON Decode error",
			relativePath:   "../fixtures/wrong_fare_configuration.json",
			expectedResult: nil,
			expectedError:  errors.New("invalid character 'i' looking for beginning of value"),
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", "invalid character 'i' looking for beginning of value").Return()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockLogger := new(mocks.Logger)

			fareConfigRepo := _fareConfigRepo.NewFareConfigurationRepository(mockLogger)

			// set up mock expectations
			tc.mockLoggerExpectation(mockLogger)

			// execute the function
			result, err := fareConfigRepo.GetAll(tc.relativePath)

			// assert the results
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}
