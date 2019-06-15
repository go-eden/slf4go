package slf4go

import (
	"fmt"
	"runtime"
)

type Logger struct {
	name     string
	fields   Fields
	provider Provider
}

func NewLogger(s string) *Logger {
	return &Logger{
		name:     s,
		provider: provider,
	}
}

func (l *Logger) GetName() string {
	return l.name
}

func (l *Logger) IsEnableTrace() bool {
	pc, _, _, _ := runtime.Caller(1)
	return l.level(pc) >= LEVEL_TRACE
}

func (l *Logger) IsEnableDebug() bool {
	pc, _, _, _ := runtime.Caller(1)
	return l.level(pc) >= LEVEL_DEBUG
}

func (l *Logger) IsEnableInfo() bool {
	pc, _, _, _ := runtime.Caller(1)
	return l.level(pc) >= LEVEL_INFO
}

func (l *Logger) IsEnableWarn() bool {
	pc, _, _, _ := runtime.Caller(1)
	return l.level(pc) >= LEVEL_WARN
}

func (l *Logger) IsEnableError() bool {
	pc, _, _, _ := runtime.Caller(1)
	return l.level(pc) >= LEVEL_ERROR
}

func (l *Logger) IsEnableFatal() bool {
	pc, _, _, _ := runtime.Caller(1)
	return l.level(pc) >= LEVEL_FATAL
}

func (l *Logger) Trace(v ...interface{}) {
	pc, fileName, line, _ := runtime.Caller(1)
	msg := fmt.Sprint(v...)
	l.print(NewLog(LEVEL_TRACE, pc, fileName, line, msg))
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	pc, fileName, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf(format, v...)
	l.print(NewLog(LEVEL_TRACE, pc, fileName, line, msg))
}

func (l *Logger) Debug(v ...interface{}) {
	pc, fileName, line, _ := runtime.Caller(1)
	msg := fmt.Sprint(v...)
	l.print(NewLog(LEVEL_DEBUG, pc, fileName, line, msg))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	pc, fileName, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf(format, v...)
	l.print(NewLog(LEVEL_DEBUG, pc, fileName, line, msg))
}

func (l *Logger) Info(v ...interface{}) {
	pc, fileName, line, _ := runtime.Caller(1)
	msg := fmt.Sprint(v...)
	l.print(NewLog(LEVEL_INFO, pc, fileName, line, msg))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	pc, fileName, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf(format, v...)
	l.print(NewLog(LEVEL_INFO, pc, fileName, line, msg))
}

func (l *Logger) Warn(v ...interface{}) {
	pc, fileName, line, _ := runtime.Caller(1)
	msg := fmt.Sprint(v...)
	l.print(NewLog(LEVEL_WARN, pc, fileName, line, msg))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	pc, fileName, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf(format, v...)
	l.print(NewLog(LEVEL_WARN, pc, fileName, line, msg))
}

func (l *Logger) Error(v ...interface{}) {
	pc, fileName, line, _ := runtime.Caller(1)
	msg := fmt.Sprint(v...)
	l.print(NewLog(LEVEL_ERROR, pc, fileName, line, msg))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	pc, fileName, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf(format, v...)
	l.print(NewLog(LEVEL_ERROR, pc, fileName, line, msg))
}

func (l *Logger) Fatal(v ...interface{}) {
	pc, fileName, line, _ := runtime.Caller(1)
	msg := fmt.Sprint(v...)
	l.print(NewLog(LEVEL_FATAL, pc, fileName, line, msg))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	pc, fileName, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf(format, v...)
	l.print(NewLog(LEVEL_FATAL, pc, fileName, line, msg))
}

// BindFields add the specified fields into the current Logger.
func (l *Logger) BindFields(fields Fields) {
	l.fields = NewFields(l.fields, fields)
}

// WithFields derive an new Logger by the specified fields from the current Logger.
func (l *Logger) WithFields(fields Fields) *Logger {
	result := NewLogger(l.name)
	result.BindFields(NewFields(l.fields, fields))
	return result
}

// do print
func (l *Logger) print(log *Log) {
	l.provider.Print(log)
}

// retrieve current logger's lowest level
func (l *Logger) level(pc uintptr) Level {
	logger := l.name
	if logger == "" {
		logger, _ = parseFunc(pc)
	}
	return l.provider.GetLevel(logger)
}
