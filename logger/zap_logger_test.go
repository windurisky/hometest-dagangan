package logger_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/windurisky/hometest-dagangan/logger"
)

func TestLoggerFunctions(t *testing.T) {
	logger, err := logger.NewZapLogger()
	assert.NoError(t, err)

	keyValLogFields := map[string]interface{}{
		"sample_data": 123,
	}
	logger.Info("Test", keyValLogFields)

	valueLogFields := []interface{}{
		123,
		"abc",
	}
	logger.Error("Error", valueLogFields)
	logger.Close()
}
