package rpl

import "fmt"

// Log levels
const (
	LevelSilent  int8 = 0
	LevelError   int8 = 1
	LevelSuccess int8 = 1
	LevelInfo    int8 = 2
	LevelWarn    int8 = 2
	LevelDebug   int8 = 3
)

// Logger is the most common Source,
// produce Log on remote.
type Logger struct {
	targets []Target

	// [Ch] is the [Log] channel of [Logger].
	Ch uint16
}

func NewLogger(ch uint16) *Logger {
	return &Logger{
		Ch: ch,
	}
}

func (logger *Logger) Register(target Target) {
	logger.targets = append(logger.targets, target)
}

func (logger *Logger) Logs(level int8, value string) {
	log := Log{
		Ch:    logger.Ch,
		Level: level,
		Value: value,
	}

	for _, target := range logger.targets {
		go func(t Target, l Log) {
			t.Writer() <- l
		}(target, log)
	}
}

func (logger *Logger) Log(level int8, args ...any) {
	logger.Logs(level, fmt.Sprint(args...))
}

func (logger *Logger) Logf(level int8, format string, args ...any) {
	logger.Log(level, fmt.Sprintf(format, args...))
}

func (logger *Logger) Success(args ...any) {
	logger.Log(LevelSuccess, args...)
}

func (logger *Logger) Error(args ...any) {
	logger.Log(LevelError, args...)
}

func (logger *Logger) Info(args ...any) {
	logger.Log(LevelInfo, args...)
}

func (logger *Logger) Warn(args ...any) {
	logger.Log(LevelWarn, args...)
}

func (logger *Logger) Debug(args ...any) {
	logger.Log(LevelDebug, args...)
}

func (logger *Logger) Successf(format string, args ...any) {
	logger.Logf(LevelSuccess, format, args...)
}

func (logger *Logger) Errorf(format string, args ...any) {
	logger.Logf(LevelError, format, args...)
}

func (logger *Logger) Infof(format string, args ...any) {
	logger.Logf(LevelInfo, format, args...)
}

func (logger *Logger) Warnf(format string, args ...any) {
	logger.Logf(LevelWarn, format, args...)
}

func (logger *Logger) Debugf(format string, args ...any) {
	logger.Logf(LevelDebug, format, args...)
}
