package slog

import (
	"github.com/go-eden/common/etime"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
)

// the cached id of current process
var pid = os.Getpid()

// the start time of current process
var startTime = etime.NowMicrosecond() - 1

var context string       // the process name
var globalDriver Driver  // the log driver
var globalLevel Level    // global lowest level, will cover all driver's configuration
var globalLogger *Logger // global default logger
var globalHook = newHooks()

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
	globalLogger = newLogger(nil)
	// setup default level
	globalLevel = TraceLevel
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
func SetLevel(l Level) {
	globalLevel = l
}

// RegisterHook register a hook, all log will inform it
func RegisterHook(f func(*Log)) {
	globalHook.addHook(f)
}

// GetLogger create new Logger by caller's package name
func GetLogger() *Logger {
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	s := ParseStack(pc[0])
	return newLogger(&s.Package)
}

// NewLogger create new Logger by the specified name
func NewLogger(name string) *Logger {
	return newLogger(&name)
}

// Trace record trace level's log
func Trace(v ...interface{}) {
	if globalLevel > TraceLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(TraceLevel, pc[0], nil, v...)
}

// Tracef record trace level's log with custom format.
func Tracef(format string, v ...interface{}) {
	if globalLevel > TraceLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(TraceLevel, pc[0], nil, format, v...)
}

// Debug record debug level's log
func Debug(v ...interface{}) {
	if globalLevel > DebugLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(DebugLevel, pc[0], nil, v...)
}

// Debugf record debug level's log with custom format.
func Debugf(format string, v ...interface{}) {
	if globalLevel > DebugLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(DebugLevel, pc[0], nil, format, v...)
}

// Info record info level's log
func Info(v ...interface{}) {
	if globalLevel > InfoLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(InfoLevel, pc[0], nil, v...)
}

// Infof record info level's log with custom format.
func Infof(format string, v ...interface{}) {
	if globalLevel > InfoLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(InfoLevel, pc[0], nil, format, v...)
}

// Warn record warn level's log
func Warn(v ...interface{}) {
	if globalLevel > WarnLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(WarnLevel, pc[0], nil, v...)
}

// Warnf record warn level's log with custom format.
func Warnf(format string, v ...interface{}) {
	if globalLevel > WarnLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(WarnLevel, pc[0], nil, format, v...)
}

// Error record error level's log
func Error(v ...interface{}) {
	if globalLevel > ErrorLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(ErrorLevel, pc[0], nil, v...)
}

// Errorf record error level's log with custom format.
func Errorf(format string, v ...interface{}) {
	if globalLevel > ErrorLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(ErrorLevel, pc[0], nil, format, v...)
}

// Panic record panic level's log
func Panic(v ...interface{}) {
	if globalLevel > PanicLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	globalLogger.print(PanicLevel, pc[0], &stack, v...)
}

// Panicf record panic level's log with custom format
func Panicf(format string, v ...interface{}) {
	if globalLevel > PanicLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	globalLogger.printf(PanicLevel, pc[0], &stack, format, v...)
}

// Fatal record fatal level's log
func Fatal(v ...interface{}) {
	if globalLevel > FatalLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	globalLogger.print(FatalLevel, pc[0], &stack, v...)
}

// Fatalf record fatal level's log with custom format.
func Fatalf(format string, v ...interface{}) {
	if globalLevel > FatalLevel {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	stack := string(debug.Stack())
	globalLogger.printf(FatalLevel, pc[0], &stack, format, v...)
}
