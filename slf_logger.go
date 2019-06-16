package xlog

import (
	"fmt"
	"runtime"
)

type Logger struct {
	name   string
	fields Fields
	driver Driver
}

func newLogger(s string) *Logger {
	return &Logger{
		name:   s,
		driver: driver,
	}
}

func (l *Logger) GetName() string {
	return l.name
}

func (l *Logger) IsEnableTrace() bool {
	pc := make([]uintptr, 1, 1)
	if l.name != "" {
		_ = runtime.Callers(2, pc[:])
	}
	return l.level(pc[0]) >= LEVEL_TRACE
}

func (l *Logger) IsEnableDebug() bool {
	pc := make([]uintptr, 1, 1)
	if l.name != "" {
		_ = runtime.Callers(2, pc[:])
	}
	return l.level(pc[0]) >= LEVEL_DEBUG
}

func (l *Logger) IsEnableInfo() bool {
	pc := make([]uintptr, 1, 1)
	if l.name != "" {
		_ = runtime.Callers(2, pc[:])
	}
	return l.level(pc[0]) >= LEVEL_INFO
}

func (l *Logger) IsEnableWarn() bool {
	pc := make([]uintptr, 1, 1)
	if l.name != "" {
		_ = runtime.Callers(2, pc[:])
	}
	return l.level(pc[0]) >= LEVEL_WARN
}

func (l *Logger) IsEnableError() bool {
	pc := make([]uintptr, 1, 1)
	if l.name != "" {
		_ = runtime.Callers(2, pc[:])
	}
	return l.level(pc[0]) >= LEVEL_ERROR
}

func (l *Logger) IsEnableFatal() bool {
	pc := make([]uintptr, 1, 1)
	if l.name != "" {
		_ = runtime.Callers(2, pc[:])
	}
	return l.level(pc[0]) >= LEVEL_FATAL
}

func (l *Logger) Trace(v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_TRACE {
		return // don't need log
	}
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(2, pc[:])
	msg := fmt.Sprint(v...)
	l.print(NewLog(LEVEL_TRACE, pc[0], msg))
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_TRACE {
		return // don't need log
	}
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(2, pc[:])
	msg := fmt.Sprintf(format, v...)
	l.print(NewLog(LEVEL_TRACE, pc[0], msg))
}

func (l *Logger) Debug(v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_DEBUG {
		return // don't need log
	}
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(2, pc[:])
	msg := fmt.Sprint(v...)
	l.print(NewLog(LEVEL_DEBUG, pc[0], msg))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_DEBUG {
		return // don't need log
	}
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(2, pc[:])
	msg := fmt.Sprintf(format, v...)
	l.print(NewLog(LEVEL_DEBUG, pc[0], msg))
}

func (l *Logger) Info(v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_INFO {
		return // don't need log
	}
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(2, pc[:])
	msg := fmt.Sprint(v...)
	l.print(NewLog(LEVEL_INFO, pc[0], msg))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_INFO {
		return // don't need log
	}
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(2, pc[:])
	msg := fmt.Sprintf(format, v...)
	l.print(NewLog(LEVEL_INFO, pc[0], msg))
}

func (l *Logger) Warn(v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_WARN {
		return // don't need log
	}
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(2, pc[:])
	msg := fmt.Sprint(v...)
	l.print(NewLog(LEVEL_WARN, pc[0], msg))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_WARN {
		return // don't need log
	}
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(2, pc[:])
	msg := fmt.Sprintf(format, v...)
	l.print(NewLog(LEVEL_WARN, pc[0], msg))
}

func (l *Logger) Error(v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_ERROR {
		return // don't need log
	}
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(2, pc[:])
	msg := fmt.Sprint(v...)
	l.print(NewLog(LEVEL_ERROR, pc[0], msg))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_ERROR {
		return // don't need log
	}
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(2, pc[:])
	msg := fmt.Sprintf(format, v...)
	l.print(NewLog(LEVEL_ERROR, pc[0], msg))
}

func (l *Logger) Fatal(v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_FATAL {
		return // don't need log
	}
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(2, pc[:])
	msg := fmt.Sprint(v...)
	l.print(NewLog(LEVEL_FATAL, pc[0], msg))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_FATAL {
		return // don't need log
	}
	pc := make([]uintptr, 1, 1)
	_ = runtime.Callers(2, pc[:])
	msg := fmt.Sprintf(format, v...)
	l.print(NewLog(LEVEL_FATAL, pc[0], msg))
}

// BindFields add the specified fields into the current Logger.
func (l *Logger) BindFields(fields Fields) {
	l.fields = NewFields(l.fields, fields)
}

// WithFields derive an new Logger by the specified fields from the current Logger.
func (l *Logger) WithFields(fields Fields) *Logger {
	result := newLogger(l.name)
	result.BindFields(NewFields(l.fields, fields))
	return result
}

// do print
func (l *Logger) print(log *Log) {
	l.driver.Print(log)
}

// retrieve current logger's lowest level
func (l *Logger) level(pc uintptr) Level {
	logger := l.name
	if logger == "" {
		logger = ParseStack(pc).pkgName
	}
	return l.driver.GetLevel(logger)
}
