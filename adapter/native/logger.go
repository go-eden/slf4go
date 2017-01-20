package native

import (
    "log"
    "os"
    "fmt"
)

const (
    LEVEL_TRACE = "trace"
    LEVEL_DEBUG = "debug"
    LEVEL_INFO  = "info"
    LEVEL_WARN  = "warn"
    LEVEL_ERROR = "error"
    LEVEL_FATAL = "fatal"
)

const PREFIX_FORMAT = "[%s] [%s] "

// simple logger that use native log
type Logger struct {
    name        string
    traceLogger *log.Logger
    debugLogger *log.Logger
    infoLogger  *log.Logger
    warnLogger  *log.Logger
    errorLogger *log.Logger
    fatalLogger *log.Logger
}

// it should be private
func NewSimpleLogger(name string) Logger {
    flag := log.Ldate | log.Ltime | log.Lshortfile
    logger := Logger{}
    logger.name = name
    logger.traceLogger = log.New(os.Stdout, fmt.Sprintf(PREFIX_FORMAT, name, LEVEL_TRACE), flag)
    logger.debugLogger = log.New(os.Stdout, fmt.Sprintf(PREFIX_FORMAT, name, LEVEL_DEBUG), flag)
    logger.infoLogger = log.New(os.Stdout, fmt.Sprintf(PREFIX_FORMAT, name, LEVEL_INFO), flag)
    logger.warnLogger = log.New(os.Stdout, fmt.Sprintf(PREFIX_FORMAT, name, LEVEL_WARN), flag)
    logger.errorLogger = log.New(os.Stderr, fmt.Sprintf(PREFIX_FORMAT, name, LEVEL_ERROR), flag)
    logger.fatalLogger = log.New(os.Stderr, fmt.Sprintf(PREFIX_FORMAT, name, LEVEL_FATAL), flag)
    return logger
}

func (logger Logger) GetName() string {
    return logger.name
}

func (logger Logger) Trace(args ...interface{}) {
    logger.traceLogger.Println(args)
}

func (logger Logger) TraceF(format string, args ...interface{}) {
    logger.traceLogger.Printf(format, args)
}

func (logger Logger) Debug(args ...interface{}) {
    logger.debugLogger.Println(args)
}

func (logger Logger) DebugF(format string, args ...interface{}) {
    logger.debugLogger.Printf(format, args)
}

func (logger Logger) Info(args ...interface{}) {
    logger.infoLogger.Println(args)
}

func (logger Logger) InfoF(format string, args ...interface{}) {
    logger.infoLogger.Printf(format, args)
}

func (logger Logger) Warn(args ...interface{}) {
    logger.warnLogger.Println(args)
}

func (logger Logger) WarnF(format string, args ...interface{}) {
    logger.warnLogger.Printf(format, args)
}

func (logger Logger) Error(args ...interface{}) {
    logger.errorLogger.Println(args)
}

func (logger Logger) ErrorF(format string, args ...interface{}) {
    logger.errorLogger.Printf(format, args)
}

func (logger Logger) Fatal(args ...interface{}) {
    logger.fatalLogger.Fatalln(args)
}

func (logger Logger) FatalF(format string, args ...interface{}) {
    logger.fatalLogger.Fatalf(format, args)
}
