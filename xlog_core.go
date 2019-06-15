package xlog

import (
	"os"
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

// get l from the LoggerFactory provided, if not, use the native log.
func GetLogger(name string) *Logger {
	return NewLogger(name)
}

// LoggerFactory is Logger's provider
type LoggerFactory interface {
	GetLogger(name string) Logger
}
