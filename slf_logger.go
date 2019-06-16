package xlog

import (
	"runtime"
)

// Logger wrap independent logger
type Logger struct {
	name   *string
	fields Fields
}

func newLogger(s *string) *Logger {
	return &Logger{name: s}
}

// Name obtain logger's name
func (l *Logger) Name() string {
	return *l.name
}

// Level obtain logger's level, lower will not be print
func (l *Logger) Level() Level {
	r := globalDriver.GetLevel(*l.name)
	if r < globalLevel {
		r = globalLevel
	}
	return r
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
func (l *Logger) IsTraceEnabled() bool {
	return l.Level() <= LEVEL_TRACE
}

// Whether debug of current logger enabled or not
func (l *Logger) IsDebugEnabled() bool {
	return l.Level() <= LEVEL_DEBUG
}

// Whether info of current logger enabled or not
func (l *Logger) IsInfoEnabled() bool {
	return l.Level() <= LEVEL_INFO
}

// Whether warn of current logger enabled or not
func (l *Logger) IsWarnEnabled() bool {
	return l.Level() <= LEVEL_WARN
}

// Whether error of current logger enabled or not
func (l *Logger) IsErrorEnabled() bool {
	return l.Level() <= LEVEL_ERROR
}

// Whether fatal of current logger enabled or not
func (l *Logger) IsFatalEnabled() bool {
	return l.Level() <= LEVEL_FATAL
}

// Trace record trace level's log
func (l *Logger) Trace(v ...interface{}) {
	if !l.IsTraceEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(LEVEL_TRACE, pc[0], v...)
}

// Tracef record trace level's log with custom format.
func (l *Logger) Tracef(format string, v ...interface{}) {
	if !l.IsTraceEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(LEVEL_TRACE, pc[0], format, v...)
}

// Debug record debug level's log
func (l *Logger) Debug(v ...interface{}) {
	if !l.IsDebugEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(LEVEL_DEBUG, pc[0], v...)
}

// Debugf record debug level's log with custom format.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if !l.IsDebugEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(LEVEL_DEBUG, pc[0], format, v...)
}

// Info record info level's log
func (l *Logger) Info(v ...interface{}) {
	if !l.IsInfoEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(LEVEL_INFO, pc[0], v...)
}

// Infof record info level's log with custom format.
func (l *Logger) Infof(format string, v ...interface{}) {
	if !l.IsInfoEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(LEVEL_INFO, pc[0], format, v...)
}

// Warn record warn level's log
func (l *Logger) Warn(v ...interface{}) {
	if !l.IsWarnEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(LEVEL_WARN, pc[0], v...)
}

// Warnf record warn level's log with custom format.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if !l.IsWarnEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(LEVEL_WARN, pc[0], format, v...)
}

// Error record error level's log
func (l *Logger) Error(v ...interface{}) {
	if !l.IsErrorEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(LEVEL_ERROR, pc[0], v...)
}

// Errorf record error level's log with custom format.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if !l.IsErrorEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(LEVEL_ERROR, pc[0], format, v...)
}

// Fatal record fatal level's log
func (l *Logger) Fatal(v ...interface{}) {
	if !l.IsFatalEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(LEVEL_FATAL, pc[0], v...)
}

// Fatalf record fatal level's log with custom format.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if !l.IsFatalEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(LEVEL_FATAL, pc[0], format, v...)
}

// do print
func (l *Logger) print(level Level, pc uintptr, v ...interface{}) {
	log := NewLog(level, pc, nil, v, l.fields)
	if l.name != nil {
		log.Logger = *l.name
	}
	globalHook.broadcast(log)
	globalDriver.Print(log)
}

// do printf
func (l *Logger) printf(level Level, pc uintptr, format string, v ...interface{}) {
	log := NewLog(level, pc, &format, v, l.fields)
	if l.name != nil {
		log.Logger = *l.name
	}
	globalHook.broadcast(log)
	globalDriver.Print(log)
}
