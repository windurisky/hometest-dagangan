package cmd_test

import (
	"errors"
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
		envLowerLimit         string
		envUpperLimit         string
		expectedTrip          domain.Trip
		expectedError         error
		mockLoggerExpectation func(mockLogger *mocks.Logger)
	}{
		{
			name:          "Valid Input",
			input:         "Location 00:02:10.100 10",
			envLowerLimit: "2",
			envUpperLimit: "10",
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
		{
			name:          "Outside Allowed Duration Limit",
			input:         "Location 00:10:10.100 10",
			envLowerLimit: "2",
			envUpperLimit: "10",
			expectedTrip:  domain.Trip{},
			expectedError: domain.ErrInvalidDurationRange,
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", domain.ErrInvalidDurationRange.Error()).Return()
			},
		},
		{
			name:          "Invalid Parameter count",
			input:         "Location 00:10:10.100 10 11",
			envLowerLimit: "2",
			envUpperLimit: "10",
			expectedTrip:  domain.Trip{},
			expectedError: domain.ErrInvalidTripParameterCount,
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", domain.ErrInvalidTripParameterCount.Error()).Return()
			},
		},
		{
			name:          "Invalid Duration Format",
			input:         "Location 00:00:10:10.100 10",
			envLowerLimit: "2",
			envUpperLimit: "10",
			expectedTrip:  domain.Trip{},
			expectedError: domain.ErrInvalidDurationFormat,
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", domain.ErrInvalidDurationFormat.Error()).Return()
			},
		},
		{
			name:          "Invalid Duration Format (2)",
			input:         "Location 00:10:10 10",
			envLowerLimit: "2",
			envUpperLimit: "10",
			expectedTrip:  domain.Trip{},
			expectedError: domain.ErrInvalidDurationFormat,
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", domain.ErrInvalidDurationFormat.Error()).Return()
			},
		},
		{
			name:          "Invalid Lower Limit",
			input:         "Location 00:10:10.100 10",
			envLowerLimit: "ab",
			envUpperLimit: "10",
			expectedTrip:  domain.Trip{},
			expectedError: errors.New("strconv.Atoi: parsing \"ab\": invalid syntax"),
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", "strconv.Atoi: parsing \"ab\": invalid syntax").Return()
			},
		},
		{
			name:          "Invalid Upper Limit",
			input:         "Location 00:10:10.100 10",
			envLowerLimit: "2",
			envUpperLimit: "ab",
			expectedTrip:  domain.Trip{},
			expectedError: errors.New("strconv.Atoi: parsing \"ab\": invalid syntax"),
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", "strconv.Atoi: parsing \"ab\": invalid syntax").Return()
			},
		},
		{
			name:          "Invalid Hours",
			input:         "Location ab:10:10.100 10",
			envLowerLimit: "2",
			envUpperLimit: "10",
			expectedTrip:  domain.Trip{},
			expectedError: errors.New("strconv.Atoi: parsing \"ab\": invalid syntax"),
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", "strconv.Atoi: parsing \"ab\": invalid syntax").Return()
			},
		},
		{
			name:          "Invalid Minutes",
			input:         "Location 00:ab:10.100 10",
			envLowerLimit: "2",
			envUpperLimit: "10",
			expectedTrip:  domain.Trip{},
			expectedError: errors.New("strconv.Atoi: parsing \"ab\": invalid syntax"),
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", "strconv.Atoi: parsing \"ab\": invalid syntax").Return()
			},
		},
		{
			name:          "Invalid Seconds",
			input:         "Location 00:10:ab.100 10",
			envLowerLimit: "2",
			envUpperLimit: "10",
			expectedTrip:  domain.Trip{},
			expectedError: errors.New("strconv.Atoi: parsing \"ab\": invalid syntax"),
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", "strconv.Atoi: parsing \"ab\": invalid syntax").Return()
			},
		},
		{
			name:          "Invalid Milliseconds",
			input:         "Location 00:10:10.abc 10",
			envLowerLimit: "2",
			envUpperLimit: "10",
			expectedTrip:  domain.Trip{},
			expectedError: errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", "strconv.Atoi: parsing \"abc\": invalid syntax").Return()
			},
		},
		{
			name:          "Invalid Minutes Limit",
			input:         "Location 00:60:10.100 10",
			envLowerLimit: "2",
			envUpperLimit: "10",
			expectedTrip:  domain.Trip{},
			expectedError: domain.ErrInvalidMinutes,
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", domain.ErrInvalidMinutes.Error()).Return()
			},
		},
		{
			name:          "Invalid Seconds Limit",
			input:         "Location 00:02:60.100 10",
			envLowerLimit: "2",
			envUpperLimit: "10",
			expectedTrip:  domain.Trip{},
			expectedError: domain.ErrInvalidSeconds,
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", domain.ErrInvalidSeconds.Error()).Return()
			},
		},
		{
			name:          "Invalid Milliseconds Limit",
			input:         "Location 00:02:10.1000 10",
			envLowerLimit: "2",
			envUpperLimit: "10",
			expectedTrip:  domain.Trip{},
			expectedError: domain.ErrInvalidMilliseconds,
			mockLoggerExpectation: func(mockLogger *mocks.Logger) {
				mockLogger.On("Error", domain.ErrInvalidMilliseconds.Error()).Return()
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// hardcode the related env variables in test environment
			// to ensure any changes on these will not break the test
			t.Setenv("DURATION_UPPER_LIMIT_MINUTES", tc.envUpperLimit)
			t.Setenv("DURATION_LOWER_LIMIT_MINUTES", tc.envLowerLimit)

			mockLogger := new(mocks.Logger)
			handler := _triphandler.NewTripHandler(mockLogger, nil)

			// set up mock expectations
			tc.mockLoggerExpectation(mockLogger)

			// execute the function
			result, err := handler.ParseInput(tc.input)

			// assert the results
			mockLogger.AssertExpectations(t)
			if tc.expectedError != nil {
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			}
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
