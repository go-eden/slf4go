package slog

import (
	"github.com/go-eden/common/etime"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
)

const RootLoggerName = "root"

var (
	context   string                       // the process name, useless
	pid       = os.Getpid()                // the cached id of current process
	startTime = etime.NowMicrosecond() - 1 // the start time of current process

	globalHook         = newHooks()
	globalDriver       Driver       // the log driver
	globalLogger       *Logger      // global default logger
	globalLevelSetting LevelSetting // global setting
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
	globalLogger = newLogger(RootLoggerName)
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
func GetLogger() (l *Logger) {
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
func NewLogger(name string) *Logger {
	return newLogger(name)
}

// Trace record trace level's log
func Trace(v ...interface{}) {
	if !globalLogger.IsTraceEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(TraceLevel, pc[0], nil, v...)
}

// Tracef record trace level's log with custom format.
func Tracef(format string, v ...interface{}) {
	if !globalLogger.IsTraceEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(TraceLevel, pc[0], nil, format, v...)
}

// Debug record debug level's log
func Debug(v ...interface{}) {
	if !globalLogger.IsDebugEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(DebugLevel, pc[0], nil, v...)
}

// Debugf record debug level's log with custom format.
func Debugf(format string, v ...interface{}) {
	if !globalLogger.IsDebugEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(DebugLevel, pc[0], nil, format, v...)
}

// Info record info level's log
func Info(v ...interface{}) {
	if !globalLogger.IsInfoEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(InfoLevel, pc[0], nil, v...)
}

// Infof record info level's log with custom format.
func Infof(format string, v ...interface{}) {
	if !globalLogger.IsInfoEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(InfoLevel, pc[0], nil, format, v...)
}

// Warn record warn level's log
func Warn(v ...interface{}) {
	if !globalLogger.IsWarnEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(WarnLevel, pc[0], nil, v...)
}

// Warnf record warn level's log with custom format.
func Warnf(format string, v ...interface{}) {
	if !globalLogger.IsWarnEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(WarnLevel, pc[0], nil, format, v...)
}

// Error record error level's log
func Error(v ...interface{}) {
	if !globalLogger.IsErrorEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(ErrorLevel, pc[0], nil, v...)
}

// Errorf record error level's log with custom format.
func Errorf(format string, v ...interface{}) {
	if !globalLogger.IsErrorEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(ErrorLevel, pc[0], nil, format, v...)
}

// Panic record panic level's log
func Panic(v ...interface{}) {
	if !globalLogger.IsPanicEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	globalLogger.print(PanicLevel, pc[0], &stack, v...)
}

// Panicf record panic level's log with custom format
func Panicf(format string, v ...interface{}) {
	if !globalLogger.IsPanicEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	globalLogger.printf(PanicLevel, pc[0], &stack, format, v...)
}

// Fatal record fatal level's log
func Fatal(v ...interface{}) {
	if !globalLogger.IsFatalEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	globalLogger.print(FatalLevel, pc[0], &stack, v...)
}

// Fatalf record fatal level's log with custom format.
func Fatalf(format string, v ...interface{}) {
	if !globalLogger.IsFatalEnabled() {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	globalLogger.printf(FatalLevel, pc[0], &stack, format, v...)
}
