package slog

import (
	"runtime"
	"runtime/debug"
	"sync/atomic"
)

type logger struct {
	name   string
	fields atomic.Value // Fields in logger, could be shared in all goroutines.
}

func newLogger(s string) Logger {
	l := &logger{name: s}
	l.fields.Store(Fields{})
	return l
}

// Name obtain logger's name
func (l *logger) Name() string {
	return l.name
}

func (l *logger) Level() Level {
	result := getDriver().GetLevel(l.name)
	lv := globalLevelSetting.getLoggerLevel(l.name)
	if result < lv {
		result = lv
	}
	return result
}

func (l *logger) BindFields(fields Fields) {
	oldFields := l.fields.Load().(Fields)
	l.fields.Store(NewFields(oldFields, fields))
}

func (l *logger) WithFields(fields Fields) Logger {
	oldFields := l.fields.Load().(Fields)
	result := newLogger(l.name)
	result.BindFields(NewFields(oldFields, fields))
	return result
}

func (l *logger) IsTraceEnabled() bool {
	return l.Level() <= TraceLevel
}

func (l *logger) IsDebugEnabled() bool {
	return l.Level() <= DebugLevel
}

func (l *logger) IsInfoEnabled() bool {
	return l.Level() <= InfoLevel
}

func (l *logger) IsWarnEnabled() bool {
	return l.Level() <= WarnLevel
}

func (l *logger) IsErrorEnabled() bool {
	return l.Level() <= ErrorLevel
}

func (l *logger) IsPanicEnabled() bool {
	return l.Level() <= PanicLevel
}

func (l *logger) IsFatalEnabled() bool {
	return l.Level() <= FatalLevel
}

func (l *logger) Trace(v ...interface{}) {
	if !l.IsTraceEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(TraceLevel, pc[0], nil, v...)
}

func (l *logger) Tracef(format string, v ...interface{}) {
	if !l.IsTraceEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(TraceLevel, pc[0], nil, format, v...)
}

func (l *logger) Debug(v ...interface{}) {
	if !l.IsDebugEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(DebugLevel, pc[0], nil, v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	if !l.IsDebugEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(DebugLevel, pc[0], nil, format, v...)
}

func (l *logger) Info(v ...interface{}) {
	if !l.IsInfoEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(InfoLevel, pc[0], nil, v...)
}

func (l *logger) Infof(format string, v ...interface{}) {
	if !l.IsInfoEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(InfoLevel, pc[0], nil, format, v...)
}

func (l *logger) Warn(v ...interface{}) {
	if !l.IsWarnEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(WarnLevel, pc[0], nil, v...)
}

func (l *logger) Warnf(format string, v ...interface{}) {
	if !l.IsWarnEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(WarnLevel, pc[0], nil, format, v...)
}

func (l *logger) Error(v ...interface{}) {
	if !l.IsErrorEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.print(ErrorLevel, pc[0], nil, v...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	if !l.IsErrorEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	l.printf(ErrorLevel, pc[0], nil, format, v...)
}

func (l *logger) Panic(v ...interface{}) {
	if !l.IsPanicEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	l.print(PanicLevel, pc[0], &stack, v...)
}

func (l *logger) Panicf(format string, v ...interface{}) {
	if !l.IsPanicEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	l.printf(PanicLevel, pc[0], &stack, format, v...)
}

func (l *logger) Fatal(v ...interface{}) {
	if !l.IsFatalEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	l.print(FatalLevel, pc[0], &stack, v...)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	if !l.IsFatalEnabled() {
		return // don't need log
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	l.printf(FatalLevel, pc[0], &stack, format, v...)
}

func (l *logger) print(level Level, pc uintptr, stack *string, v ...interface{}) {
	var cxtFields Fields
	if v := globalCxtFields.Get(); v != nil {
		cxtFields = v.(Fields)
	}
	log := NewLog(level, pc, stack, nil, v, l.fields.Load().(Fields), cxtFields)
	log.Logger = l.name
	globalHook.broadcast(log)
	getDriver().Print(log)
}

func (l *logger) printf(level Level, pc uintptr, stack *string, format string, v ...interface{}) {
	var cxtFields Fields
	if v := globalCxtFields.Get(); v != nil {
		cxtFields = v.(Fields)
	}
	log := NewLog(level, pc, stack, &format, v, l.fields.Load().(Fields), cxtFields)
	log.Logger = l.name
	globalHook.broadcast(log)
	getDriver().Print(log)
}
