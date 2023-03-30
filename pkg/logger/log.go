package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level string

func NewLogger(service string, logLevel Level, path []string) (*Logger, error) {
	zapLevel := getZapLevel(logLevel)

	cfg := newZapConfig(service, zapLevel)
	if len(path) != 0 {
		cfg.OutputPaths = path
	}
	zl, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return newLogger(zl), nil
}

func getZapLevel(l Level) zapcore.Level {
	out, ok := levelToZap[l]
	if !ok {
		return defaultLevel
	}
	return out
}

func newZapConfig(service string, level zapcore.Level) zap.Config {
	return zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "lvl",
			TimeKey:        "ts",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochNanosTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields: map[string]interface{}{
			"source": service,
		},
	}
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.zapLogger.Infow(msg, fieldsToInterface(fields)...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.zapLogger.Fatalw(msg, fieldsToInterface(fields)...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.zapLogger.Errorw(msg, fieldsToInterface(fields)...)
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.zapLogger.Debugw(msg, fieldsToInterface(fields)...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.zapLogger.Warnw(msg, fieldsToInterface(fields)...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.zapLogger.Infof(template, args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.zapLogger.Fatalf(template, args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.zapLogger.Errorf(template, args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.zapLogger.Debugf(template, args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.zapLogger.Warnf(template, args...)
}

func (l *Logger) AddField(name, value string) *Logger {
	return &Logger{
		zapLogger: l.zapLogger.With(name, value),
	}
}
