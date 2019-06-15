package xlog

import (
	"os"
	"runtime"
	"strings"
	"time"
)

var pid = os.Getpid()                 // the cached id of current process
var startTime = time.Now().UnixNano() // the start time of current process

var context string    // the process name
var provider Provider // the log provider

func init() {
	exec := os.Args[0]
	sp := uint8(os.PathSeparator)
	if off := strings.LastIndexByte(exec, sp); off > 0 {
		exec = exec[off+1:]
	}
	// setup default context
	SetContext(exec)
	// setup default provider
	SetProvider(new(StdProvider))
}

// SetContext update the global context name
func SetContext(name string) {
	context = name
}

// SetProvider update the global provider
func SetProvider(p Provider) {
	provider = p
}

// NewLogger create new Logger by caller's package name
func NewLogger() *Logger {
	pc, _, _, _ := runtime.Caller(1)
	pkgName, _ := parseFunc(pc)
	return newLogger(pkgName)
}

// GetLogger create new Logger by the specified name
func GetLogger(name string) *Logger {
	return newLogger(name)
}

// LoggerFactory is Logger's provider
type LoggerFactory interface {
	GetLogger(name string) Logger
}
