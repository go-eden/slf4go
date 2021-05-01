package slog

import (
	"runtime"
	"runtime/debug"
	"sync/atomic"
)

// Logger wrap independent logger
type Logger struct {
	name   *string
	fields atomic.Value // Fields
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
	dv := globalDriver.GetLevel(*l.name)
	lv := globalLevelSetting.getLoggerLevel(*l.name)
	if dv < lv {
		dv = lv
	}
	return dv
}

// BindFields add the specified fields into the current Logger.
func (l *Logger) BindFields(fields Fields) {
	oldFields := l.fields.Load().(Fields)
	l.fields.Store(NewFields(oldFields, fields))
}

// WithFields derive an new Logger by the specified fields from the current Logger.
func (l *Logger) WithFields(fields Fields) *Logger {
	oldFields := l.fields.Load().(Fields)
	result := newLogger(l.name)
	result.BindFields(NewFields(oldFields, fields))
	return result
}

// IsTraceEnabled Whether trace of current logger enabled or not
func (l *Logger) IsTraceEnabled() bool {
	return l.Level() <= TraceLevel
}

// IsDebugEnabled Whether debug of current logger enabled or not
func (l *Logger) IsDebugEnabled() bool {
	return l.Level() <= DebugLevel
}

// IsInfoEnabled Whether info of current logger enabled or not
func (l *Logger) IsInfoEnabled() bool {
	return l.Level() <= InfoLevel
}

// IsWarnEnabled Whether warn of current logger enabled or not
func (l *Logger) IsWarnEnabled() bool {
	return l.Level() <= WarnLevel
}

// IsErrorEnabled Whether error of current logger enabled or not
func (l *Logger) IsErrorEnabled() bool {
	return l.Level() <= ErrorLevel
}

// IsPanicEnabled Whether panic of current logger enabled or not
func (l *Logger) IsPanicEnabled() bool {
	return l.Level() <= PanicLevel
}

// IsFatalEnabled Whether fatal of current logger enabled or not
func (l *Logger) IsFatalEnabled() bool {
	return l.Level() <= FatalLevel
}

// Trace record trace level's log
func (l *Logger) Trace(v ...interface{}) {
	if !l.IsTraceEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(TraceLevel, pc[0], nil, v...)
}

// Tracef record trace level's log with custom format.
func (l *Logger) Tracef(format string, v ...interface{}) {
	if !l.IsTraceEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(TraceLevel, pc[0], nil, format, v...)
}

// Debug record debug level's log
func (l *Logger) Debug(v ...interface{}) {
	if !l.IsDebugEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(DebugLevel, pc[0], nil, v...)
}

// Debugf record debug level's log with custom format.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if !l.IsDebugEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(DebugLevel, pc[0], nil, format, v...)
}

// Info record info level's log
func (l *Logger) Info(v ...interface{}) {
	if !l.IsInfoEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(InfoLevel, pc[0], nil, v...)
}

// Infof record info level's log with custom format.
func (l *Logger) Infof(format string, v ...interface{}) {
	if !l.IsInfoEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(InfoLevel, pc[0], nil, format, v...)
}

// Warn record warn level's log
func (l *Logger) Warn(v ...interface{}) {
	if !l.IsWarnEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(WarnLevel, pc[0], nil, v...)
}

// Warnf record warn level's log with custom format.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if !l.IsWarnEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(WarnLevel, pc[0], nil, format, v...)
}

// Error record error level's log
func (l *Logger) Error(v ...interface{}) {
	if !l.IsErrorEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(ErrorLevel, pc[0], nil, v...)
}

// Errorf record error level's log with custom format.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if !l.IsErrorEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(ErrorLevel, pc[0], nil, format, v...)
}

// Panic record panic level's log
func (l *Logger) Panic(v ...interface{}) {
	if !l.IsPanicEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	l.print(FatalLevel, pc[0], &stack, v...)
}

// Panicf record panic level's log with custom format
func (l *Logger) Panicf(format string, v ...interface{}) {
	if !l.IsPanicEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	l.printf(ErrorLevel, pc[0], &stack, format, v...)
}

// Fatal record fatal level's log
func (l *Logger) Fatal(v ...interface{}) {
	if !l.IsFatalEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	l.print(FatalLevel, pc[0], &stack, v...)
}

// Fatalf record fatal level's log with custom format.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if !l.IsFatalEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	l.printf(FatalLevel, pc[0], &stack, format, v...)
}

func (l *Logger) print(level Level, pc uintptr, stack *string, v ...interface{}) {
	log := NewLog(level, pc, stack, nil, v, l.fields.Load().(Fields))
	if l.name != nil {
		log.Logger = *l.name
	}
	globalHook.broadcast(log)
	globalDriver.Print(log)
}

func (l *Logger) printf(level Level, pc uintptr, stack *string, format string, v ...interface{}) {
	log := NewLog(level, pc, stack, &format, v, l.fields.Load().(Fields))
	if l.name != nil {
		log.Logger = *l.name
	}
	globalHook.broadcast(log)
	globalDriver.Print(log)
}
