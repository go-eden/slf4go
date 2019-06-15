package slf4go

import (
	"io"
	"os"
)

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
			defaultFactory = newNativeLoggerFactory()
		}
		factory = defaultFactory
	}
	return factory.GetLogger(name)
}
