package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/windurisky/hometest-dagangan/util"
)

func TestGetEnvWithDefault(t *testing.T) {
	testCases := []struct {
		name           string
		key            string
		defaultValue   string
		expectedResult string
	}{
		{
			name:           "Environment Variable Exists",
			key:            "DURATION_LOWER_LIMIT_MINUTES",
			defaultValue:   "100",
			expectedResult: "2",
		},
		{
			name:           "Environment Variable Does Not Exist",
			key:            "ABCDEFGHIJ",
			defaultValue:   "100",
			expectedResult: "100",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// hardcode the related env variables in test environment
			// to ensure any changes on these will not break the test
			t.Setenv("DURATION_LOWER_LIMIT_MINUTES", tc.expectedResult)

			result := util.GetEnvWithDefault(tc.key, tc.defaultValue)

			// Assert the result matches the expected value
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
