package slf4go

import (
	"io"
	"os"
	"strings"
	"time"
)

var pid int         // the cached id of current process
var context string  // the process name
var startTime int64 // the start time of current process
var provider Provider

func init() {
	pid = os.Getpid()
	startTime = time.Now().UnixNano()
	// parse context
	exec := os.Args[0]
	sp := uint8(os.PathSeparator)
	if off := strings.LastIndexByte(exec, sp); off > 0 {
		exec = exec[off+1:]
	}
	SetContext(exec)
}

// SetContext update the global context name
func SetContext(name string) {
	context = name
}

// SetProvider update the global provider
func SetProvider(p Provider) {
	provider = p
}

// global log factory, could be replaced
var defaultFactory LoggerFactory
var definedFactory LoggerFactory
var Writer io.Writer = os.Stdout

// support any log framework by LoggerFactory
func SetLoggerFactory(factory LoggerFactory) {
	if factory == nil {
		panic("LoggerFactory can't be nil")
	}
	definedFactory = factory
}

// get l from the LoggerFactory provided, if not, use the native log.
func GetLogger(name string) Logger {
	var factory LoggerFactory
	if definedFactory != nil {
		factory = definedFactory
	} else {
		if defaultFactory == nil {
			//defaultFactory = newNativeLoggerFactory()
		}
		factory = defaultFactory
	}
	return factory.GetLogger(name)
}

// LoggerFactory is Logger's provider
type LoggerFactory interface {
	GetLogger(name string) Logger
}
