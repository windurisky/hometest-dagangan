package logger

import (
	"strconv"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() (*ZapLogger, error) {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:      "timestamp",
			LevelKey:     "severity",
			NameKey:      "logger",
			CallerKey:    "caller",
			MessageKey:   "message",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	zapLogger, _ := cfg.Build()
	return &ZapLogger{logger: zapLogger}, nil
}

func (z *ZapLogger) Info(msg string, args ...interface{}) {
	fields := z.genericArgsToFields(args...)
	z.logger.Info(msg, fields...)
}

func (z *ZapLogger) Error(msg string, args ...interface{}) {
	fields := z.genericArgsToFields(args...)
	fields = append(fields, zap.Stack("stacktrace"))
	z.logger.Error(msg, fields...)
}

func (z *ZapLogger) Close() error {
	return z.logger.Sync()
}

func (z *ZapLogger) genericArgsToFields(args ...interface{}) (fields []zap.Field) {
	fields = make([]zap.Field, 0, len(args))
	for i, arg := range args {
		// Check if arg is a map[string]interface{}
		if keyValue, ok := arg.(map[string]interface{}); ok {
			for key, value := range keyValue {
				fields = append(fields, zap.Any(key, value))
			}
		} else {
			// Otherwise generate the key by index, e.g.: arg0, arg1, and so on.
			fields = append(fields, zap.Any("arg"+strconv.Itoa(i), arg))
		}
	}
	return
}
