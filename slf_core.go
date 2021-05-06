package slog

import (
	"github.com/go-eden/common/etime"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
)

var (
	context   string                       // the process name, useless
	pid       = os.Getpid()                // the cached id of current process
	startTime = etime.NowMicrosecond() - 1 // the start time of current process

	globalHook         = newHooks()
	globalDriver       Driver       // the log driver
	globalLevelSetting LevelSetting // global setting
	globalLogger       *logger      // global default logger
)

func init() {
	exec := os.Args[0]
	sp := uint8(os.PathSeparator)
	if off := strings.LastIndexByte(exec, sp); off > 0 {
		exec = exec[off+1:]
	}
	// setup default context
	SetContext(exec)
	// setup default driver
	SetDriver(&StdDriver{})
	// setup default logger
	globalLogger = newLogger(RootLoggerName).(*logger)
	// setup default level
	globalLevelSetting.setRootLevel(TraceLevel)
}

// SetContext update the global context name
func SetContext(name string) {
	context = name
}

// GetContext obtain the global context name
func GetContext() string {
	return context
}

// SetDriver update the global log driver
func SetDriver(d Driver) {
	globalDriver = d
}

// SetLevel update the global level, all lower level will not be send to driver to print
func SetLevel(lv Level) {
	globalLevelSetting.setRootLevel(lv)
}

// SetLoggerLevel setup log-level of the specified logger by name
func SetLoggerLevel(loggerName string, level Level) {
	globalLevelSetting.setLoggerLevel(map[string]Level{loggerName: level})
}

// SetLoggerLevelMap batch set logger's level
func SetLoggerLevelMap(levelMap map[string]Level) {
	globalLevelSetting.setLoggerLevel(levelMap)
}

// RegisterHook register a hook, all log will inform it
func RegisterHook(f func(*Log)) {
	globalHook.addHook(f)
}

// GetLogger create new Logger by caller's package name
func GetLogger() (l Logger) {
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	s := ParseStack(pc[0])
	if name := s.Package; len(name) > 0 {
		return newLogger(name)
	}
	Warnf("cannot parse package, use global logger")
	return globalLogger
}

// NewLogger create new Logger by the specified name
func NewLogger(name string) Logger {
	return newLogger(name)
}

// Trace reference Logger.Trace
func Trace(v ...interface{}) {
	if !globalLogger.IsTraceEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(TraceLevel, pc[0], nil, v...)
}

// Tracef reference Logger.Tracef
func Tracef(format string, v ...interface{}) {
	if !globalLogger.IsTraceEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(TraceLevel, pc[0], nil, format, v...)
}

// Debug reference Logger.Debug
func Debug(v ...interface{}) {
	if !globalLogger.IsDebugEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(DebugLevel, pc[0], nil, v...)
}

// Debugf reference Logger.Debugf
func Debugf(format string, v ...interface{}) {
	if !globalLogger.IsDebugEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(DebugLevel, pc[0], nil, format, v...)
}

// Info reference Logger.Info
func Info(v ...interface{}) {
	if !globalLogger.IsInfoEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(InfoLevel, pc[0], nil, v...)
}

// Infof reference Logger.Infof
func Infof(format string, v ...interface{}) {
	if !globalLogger.IsInfoEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(InfoLevel, pc[0], nil, format, v...)
}

// Warn reference Logger.Warn
func Warn(v ...interface{}) {
	if !globalLogger.IsWarnEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(WarnLevel, pc[0], nil, v...)
}

// Warnf reference Logger.Warnf
func Warnf(format string, v ...interface{}) {
	if !globalLogger.IsWarnEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(WarnLevel, pc[0], nil, format, v...)
}

// Error reference Logger.Error
func Error(v ...interface{}) {
	if !globalLogger.IsErrorEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(ErrorLevel, pc[0], nil, v...)
}

// Errorf reference Logger.Errorf
func Errorf(format string, v ...interface{}) {
	if !globalLogger.IsErrorEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(ErrorLevel, pc[0], nil, format, v...)
}

// Panic reference Logger.Panic
func Panic(v ...interface{}) {
	if !globalLogger.IsPanicEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	globalLogger.print(PanicLevel, pc[0], &stack, v...)
}

// Panicf reference Logger.Panicf
func Panicf(format string, v ...interface{}) {
	if !globalLogger.IsPanicEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	globalLogger.printf(PanicLevel, pc[0], &stack, format, v...)
}

// Fatal reference Logger.Fatal
func Fatal(v ...interface{}) {
	if !globalLogger.IsFatalEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	globalLogger.print(FatalLevel, pc[0], &stack, v...)
}

// Fatalf reference Logger.Fatalf
func Fatalf(format string, v ...interface{}) {
	if !globalLogger.IsFatalEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	globalLogger.printf(FatalLevel, pc[0], &stack, format, v...)
}
