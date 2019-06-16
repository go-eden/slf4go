package xlog

import (
	"os"
	"runtime"
	"strings"
	"time"
)

var pid = os.Getpid()                 // the cached id of current process
var startTime = time.Now().UnixNano() // the start time of current process

var context string // the process name
var driver Driver  // the log driver

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
}

// SetContext update the global context name
func SetContext(name string) {
	context = name
}

// SetDriver update the global log driver
func SetDriver(d Driver) {
	driver = d
}

// GetLogger create new Logger by caller's package name
func GetLogger() *Logger {
	var pc [1]uintptr
	_ = runtime.Callers(2, pc[:])
	s := ParseStack(pc[0])
	return newLogger(s.pkgName)
}

// NewLogger create new Logger by the specified name
func NewLogger(name string) *Logger {
	return newLogger(name)
}
