package xlog

import (
	"fmt"
	"runtime"
)

type Logger struct {
	name   string
	fields Fields
}

func newLogger(s string) *Logger {
	return &Logger{name: s}
}

// GetName obtain logger's name
func (l *Logger) GetName() string {
	return l.name
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

// Whether trace of current logger enabled or not
func (l *Logger) IsEnableTrace() bool {
	var pc [1]uintptr
	if l.name != "" {
		_ = runtime.Callers(2, pc[:])
	}
	return l.level(pc[0]) >= LEVEL_TRACE
}

// Whether debug of current logger enabled or not
func (l *Logger) IsEnableDebug() bool {
	var pc [1]uintptr
	if l.name != "" {
		_ = runtime.Callers(2, pc[:])
	}
	return l.level(pc[0]) >= LEVEL_DEBUG
}

// Whether info of current logger enabled or not
func (l *Logger) IsEnableInfo() bool {
	var pc [1]uintptr
	if l.name != "" {
		_ = runtime.Callers(2, pc[:])
	}
	return l.level(pc[0]) >= LEVEL_INFO
}

// Whether warn of current logger enabled or not
func (l *Logger) IsEnableWarn() bool {
	var pc [1]uintptr
	if l.name != "" {
		_ = runtime.Callers(2, pc[:])
	}
	return l.level(pc[0]) >= LEVEL_WARN
}

// Whether error of current logger enabled or not
func (l *Logger) IsEnableError() bool {
	var pc [1]uintptr
	if l.name != "" {
		_ = runtime.Callers(2, pc[:])
	}
	return l.level(pc[0]) >= LEVEL_ERROR
}

// Whether fatal of current logger enabled or not
func (l *Logger) IsEnableFatal() bool {
	var pc [1]uintptr
	if l.name != "" {
		_ = runtime.Callers(2, pc[:])
	}
	return l.level(pc[0]) >= LEVEL_FATAL
}

// Trace record trace level's log
func (l *Logger) Trace(v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_TRACE {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(LEVEL_TRACE, pc[0], v...)
}

// Tracef record trace level's log with custom format.
func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_TRACE {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(LEVEL_TRACE, pc[0], format, v...)
}

// Debug record debug level's log
func (l *Logger) Debug(v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_DEBUG {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(LEVEL_DEBUG, pc[0], v...)
}

// Debugf record debug level's log with custom format.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_DEBUG {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(LEVEL_DEBUG, pc[0], format, v...)
}

// Info record info level's log
func (l *Logger) Info(v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_INFO {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(LEVEL_INFO, pc[0], v...)
}

// Infof record info level's log with custom format.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_INFO {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(LEVEL_INFO, pc[0], format, v...)
}

// Warn record warn level's log
func (l *Logger) Warn(v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_WARN {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(LEVEL_WARN, pc[0], v...)
}

// Warnf record warn level's log with custom format.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_WARN {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(LEVEL_WARN, pc[0], format, v...)
}

// Error record error level's log
func (l *Logger) Error(v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_ERROR {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(LEVEL_ERROR, pc[0], v...)
}

// Errorf record error level's log with custom format.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_ERROR {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(LEVEL_ERROR, pc[0], format, v...)
}

// Fatal record fatal level's log
func (l *Logger) Fatal(v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_FATAL {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(LEVEL_FATAL, pc[0], v...)
}

// Fatalf record fatal level's log with custom format.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.name != "" && l.level(0) > LEVEL_FATAL {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(LEVEL_FATAL, pc[0], format, v...)
}

// do print
func (l *Logger) print(level Level, pc uintptr, v ...interface{}) {
	msg := fmt.Sprint(v...)
	log := NewLog(level, pc, msg)
	driver.Print(log)
}

// do printf
func (l *Logger) printf(level Level, pc uintptr, format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log := NewLog(level, pc, msg)
	driver.Print(log)
}

// retrieve current logger's lowest level
func (l *Logger) level(pc uintptr) Level {
	logger := l.name
	if logger == "" {
		logger = ParseStack(pc).pkgName
	}
	return driver.GetLevel(logger)
}
