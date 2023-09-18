package cmd_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/windurisky/hometest-dagangan/domain"
	"github.com/windurisky/hometest-dagangan/domain/mocks"
	_triphandler "github.com/windurisky/hometest-dagangan/trip/delivery/cmd"
)

func TestParseInput(t *testing.T) {
	testCases := []struct {
		name                  string
		input                 string
		expectedTrip          domain.Trip
		expectedError         error
		mockLoggerExpectation func(mockLogger *mocks.Logger)
	}{
		{
			name:  "Valid Input",
			input: "Location 00:02:10.100 10",
			expectedTrip: domain.Trip{
				Location: "Location",
				Duration: 2*time.Minute + 10*time.Second + 100*time.Millisecond,
				Mileage:  10,
			},
			expectedError: nil,
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Info", "Successfully parsed input", mock.Anything).Return()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockLogger := new(mocks.Logger)
			handler := _triphandler.NewTripHandler(mockLogger, nil)

			// set up mock expectations
			tc.mockLoggerExpectation(mockLogger)

			// execute the function
			result, err := handler.ParseInput(tc.input)

			// assert the results
			mockLogger.AssertExpectations(t)
			assert.Equal(t, tc.expectedError, err)
			assert.Equal(t, tc.expectedTrip, result)
		})
	}
}

func TestSummarizeTrip(t *testing.T) {
	mockLogger := new(mocks.Logger)
	mockTripUsecase := new(mocks.TripUsecase)

	// Create a tripHandler instance with the mock logger and tripUsecase
	handler := _triphandler.NewTripHandler(mockLogger, mockTripUsecase)

	testTrips := []domain.Trip{
		{
			Location: "Location1",
			Duration: 1 * time.Hour,
			Mileage:  10,
		},
		{
			Location: "Location2",
			Duration: 2 * time.Hour,
			Mileage:  20,
		},
	}

	// set up mock expectations
	for _, trip := range testTrips {
		mockTripUsecase.On("CalculateFare", trip.Mileage).Return(uint64(0), nil)
	}

	// execute the function
	err := handler.SummarizeTrip(testTrips)

	// assert the error result
	assert.NoError(t, err)
	mockTripUsecase.AssertNumberOfCalls(t, "CalculateFare", len(testTrips))
}
