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
	log := &Log{
		Ch:    logger.Ch,
		Level: level,
		Value: value,
	}

	for _, target := range logger.targets {
		go func(t Target, l *Log) {
			t.Writer() <- l
		}(target, log)
	}
}

func (logger *Logger) Log(level int8, args ...interface{}) {
	logger.Logs(level, fmt.Sprint(args...))
}

func (logger *Logger) Logf(level int8, format string, args ...interface{}) {
	logger.Log(level, fmt.Sprintf(format, args...))
}

func (logger *Logger) Success(args ...interface{}) {
	logger.Log(LevelSuccess, args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.Log(LevelError, args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.Log(LevelInfo, args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.Log(LevelWarn, args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.Log(LevelDebug, args...)
}

func (logger *Logger) Successf(format string, args ...interface{}) {
	logger.Logf(LevelSuccess, format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	logger.Logf(LevelError, format, args...)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	logger.Logf(LevelInfo, format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	logger.Logf(LevelWarn, format, args...)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	logger.Logf(LevelDebug, format, args...)
}
