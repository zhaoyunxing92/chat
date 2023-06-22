package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.SugaredLogger

// Info is info level
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warn is warning level
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Error is error level
func Error(args ...interface{}) {
	log.Error(args...)
}

// Debug is debug level
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Infof is format info level
func Infof(fmt string, args ...interface{}) {
	log.Infof(fmt, args...)
}

// Warnf is format warning level
func Warnf(fmt string, args ...interface{}) {
	log.Warnf(fmt, args...)
}

// Errorf is format error level
func Errorf(fmt string, args ...interface{}) {
	log.Errorf(fmt, args...)
}

// Debugf is format debug level
func Debugf(fmt string, args ...interface{}) {
	log.Debugf(fmt, args...)
}

// Fatal logs a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Fatalf logs a templated message, then calls os.Exit.
func Fatalf(fmt string, args ...interface{}) {
	log.Fatalf(fmt, args...)
}

func NewDefault() {
	New("info")
}

func New(level string) {
	var (
		lv  zapcore.Level
		err error
	)
	if lv, err = zapcore.ParseLevel(level); err != nil {
		lv = zapcore.InfoLevel
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig())

	log = zap.New(zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), lv),
		zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()

}

func encoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		CallerKey:      "line",
		NameKey:        "Logger",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
