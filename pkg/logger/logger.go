package logger

import (
	"os"
	"strings"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DefaultMaxSize  = 100
	DefaultBackups  = 10
	DefaultMaxAge   = 30
	DefaultCompress = true
)

type Logger interface {
	Debug(format string)
	Debugf(format string, args ...interface{})
	Info(format string)
	Infof(format string, args ...interface{})
	Warn(format string)
	Warnf(format string, args ...interface{})
	Error(format string)
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	WithFields(fields map[string]interface{}) Logger
}

type logger struct {
	*zap.SugaredLogger
}

type Config struct {
	Level    string
	FilePath string
}

func New(cfg Config) Logger {
	level := getZapLevel(cfg.Level)
	writer := getWirteSyncer(cfg.FilePath)
	core := zapcore.NewCore(getEncoder(), writer, level)

	zapLogger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.WarnLevel),
	)

	defer func() {
		if err := zapLogger.Sync(); err != nil {
			panic("logger sync failed")
		}
	}()

	return &logger{zapLogger.Sugar()}
}

func getWirteSyncer(filePath string) zapcore.WriteSyncer {
	if filePath == "" {
		return zapcore.Lock(os.Stdout)
	}

	return getLumberjackWriteSyncer(filePath)
}

func getLumberjackWriteSyncer(filePath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    DefaultMaxSize,
		MaxBackups: DefaultBackups,
		MaxAge:     DefaultMaxAge,
		Compress:   DefaultCompress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getZapLevel(loggerLevel string) zapcore.Level {
	level := strings.ToLower(loggerLevel)

	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func (l *logger) Debug(format string) {
	l.SugaredLogger.Desugar().Debug(format)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.SugaredLogger.Debugf(format, args...)
}

func (l *logger) Info(format string) {
	l.SugaredLogger.Desugar().Info(format)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.SugaredLogger.Infof(format, args...)
}

func (l *logger) Warn(format string) {
	l.SugaredLogger.Desugar().Warn(format)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.SugaredLogger.Warnf(format, args...)
}

func (l *logger) Error(format string) {
	l.SugaredLogger.Desugar().Error(format)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.SugaredLogger.Errorf(format, args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.SugaredLogger.Fatalf(format, args...)
}

func (l *logger) Panicf(format string, args ...interface{}) {
	l.SugaredLogger.Panicf(format, args...)
}

func (l *logger) WithFields(fields map[string]interface{}) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}

	return &logger{l.SugaredLogger.With(f...)}
}
