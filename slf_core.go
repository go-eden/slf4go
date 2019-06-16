package xlog

import (
	"os"
	"runtime"
	"strings"
)

var pid = os.Getpid() // the cached id of current process
var startTime = now() // the start time of current process

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
	SetDriver(new(StdDriver))
	// setup default logger
	globalLogger = newLogger(nil)
	// setup default level
	globalLevel = LEVEL_TRACE
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
	if globalLevel > LEVEL_TRACE {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(LEVEL_TRACE, pc[0], v...)
}

// Tracef record trace level's log with custom format.
func Tracef(format string, v ...interface{}) {
	if globalLevel > LEVEL_TRACE {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(LEVEL_TRACE, pc[0], format, v...)
}

// Debug record debug level's log
func Debug(v ...interface{}) {
	if globalLevel > LEVEL_DEBUG {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(LEVEL_DEBUG, pc[0], v...)
}

// Debugf record debug level's log with custom format.
func Debugf(format string, v ...interface{}) {
	if globalLevel > LEVEL_DEBUG {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(LEVEL_TRACE, pc[0], format, v...)
}

// Info record info level's log
func Info(v ...interface{}) {
	if globalLevel > LEVEL_INFO {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(LEVEL_INFO, pc[0], v...)
}

// Infof record info level's log with custom format.
func Infof(format string, v ...interface{}) {
	if globalLevel > LEVEL_INFO {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(LEVEL_INFO, pc[0], format, v...)
}

// Warn record warn level's log
func Warn(v ...interface{}) {
	if globalLevel > LEVEL_WARN {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(LEVEL_WARN, pc[0], v...)
}

// Warnf record warn level's log with custom format.
func Warnf(format string, v ...interface{}) {
	if globalLevel > LEVEL_WARN {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(LEVEL_WARN, pc[0], format, v...)
}

// Error record error level's log
func Error(v ...interface{}) {
	if globalLevel > LEVEL_ERROR {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(LEVEL_ERROR, pc[0], v...)
}

// Errorf record error level's log with custom format.
func Errorf(format string, v ...interface{}) {
	if globalLevel > LEVEL_ERROR {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(LEVEL_ERROR, pc[0], format, v...)
}

// Fatal record fatal level's log
func Fatal(v ...interface{}) {
	if globalLevel > LEVEL_FATAL {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.print(LEVEL_FATAL, pc[0], v...)
}

// Fatalf record fatal level's log with custom format.
func Fatalf(format string, v ...interface{}) {
	if globalLevel > LEVEL_FATAL {
		return
	}
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	globalLogger.printf(LEVEL_FATAL, pc[0], format, v...)
}
