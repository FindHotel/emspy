package logger

import (
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

// newZapCustomConfig is a reasonable custom logging configuration.
// Logging is enabled at InfoLevel and above.
//
// It enables development mode (which makes DPanicLevel logs panic), uses a
// console encoder, writes to standard error, and disables sampling.
// Stacktraces are automatically included on logs of WarnLevel and above.
//
// Brutally stolen from:
// https://github.com/uber-go/zap/blob/master/config.go
func newZapCustomConfig() zap.Config {
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	// don't print the file and line number since it will clog up the output
	encoderCfg.EncodeCaller = nil
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    encoderCfg,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

// NewZapCustom builds a custom Logger that writes InfoLevel and above
// logs to standard error in a human-friendly format.
//
// It's a shortcut for newZapCustomConfig().Build(...Option).
func NewZapCustom(options ...zap.Option) (*zap.Logger, error) {
	return newZapCustomConfig().Build(options...)
}

func init() {
	// Zap logger (uber-go): https://github.com/uber-go/zap
	// https://godoc.org/go.uber.org/zap
	// logger, err := zap.NewDevelopment()
	logger, err := NewZapCustom()

	if err != nil {
		println("error in getting the ZAP for production, fallback to EXAMPLE")
		logger = zap.NewExample()
	}

	// In contexts where performance is nice, but not critical, use
	// the SugaredLogger. It's 4-10x faster than other structured
	// logging packages and includes both structured and printf-style
	// APIs.
	log = logger.Sugar()
}

func GlobalLogger() (logger *zap.SugaredLogger) {
	return log
}

// SetLogger supports setting external logger for the library.
func SetGlobalLogger(l *zap.SugaredLogger) {
	log = l
}

type zapWrapper struct {
	*zap.SugaredLogger
}

func (w *zapWrapper) Named(name string) Logger {
	newLogger := w.SugaredLogger.Named(name)
	return &zapWrapper{newLogger}
}

func (w *zapWrapper) With(args ...interface{}) Logger {
	newLogger := w.SugaredLogger.With(args...)
	return &zapWrapper{newLogger}
}

func Must(service string) Logger {
	logger, err := NewZapCustom()
	if err != nil {
		panic(err)
	}
	result := logger.Sugar().Named(service).With("service", service)
	return &zapWrapper{result}
}
