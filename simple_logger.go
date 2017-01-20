package slf4go

import (
    "log"
    "os"
)

// simple logger that use native log
type SimpleLogger struct {
    name    string
    outImpl *log.Logger
    errImpl *log.Logger
}

// it should be private
func newSimpleLogger(name string) *SimpleLogger {
    logger := &SimpleLogger{}
    logger.name = name
    logger.outImpl = log.New(os.Stdout, name, log.Ldate|log.Ltime|log.Lshortfile)
    logger.errImpl = log.New(os.Stderr, name, log.Ldate|log.Ltime|log.Lshortfile)
    return logger
}

func (logger *SimpleLogger) GetName() string {
    return logger.name
}

func (logger *SimpleLogger) Trace(args ...interface{}) {
    logger.outImpl.Println(args)
}

func (logger *SimpleLogger) TraceF(format string, args ...interface{}) {
    logger.outImpl.Printf(format, args)
}

func (logger *SimpleLogger) Debug(args ...interface{}) {
    logger.outImpl.Println(args)
}

func (logger *SimpleLogger) DebugF(format string, args ...interface{}) {
    logger.outImpl.Printf(format, args)
}

func (logger *SimpleLogger) Info(args ...interface{}) {
    logger.outImpl.Println(args)
}

func (logger *SimpleLogger) InfoF(format string, args ...interface{}) {
    logger.outImpl.Printf(format, args)
}

func (logger *SimpleLogger) Warn(args ...interface{}) {
    logger.outImpl.Println(args)
}

func (logger *SimpleLogger) WarnF(format string, args ...interface{}) {
    logger.outImpl.Printf(format, args)
}

func (logger *SimpleLogger) Error(args ...interface{}) {
    logger.errImpl.Println(args)
}

func (logger *SimpleLogger) ErrorF(format string, args ...interface{}) {
    logger.errImpl.Printf(format, args)
}

func (logger *SimpleLogger) Fatal(args ...interface{}) {
    logger.errImpl.Println(args)
}

func (logger *SimpleLogger) FatalF(format string, args ...interface{}) {
    logger.errImpl.Printf(format, args)
}
