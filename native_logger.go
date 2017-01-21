package slf4go

import (
    "log"
    "os"
    "fmt"
)

const (
    l_TRACE = "trace"
    l_DEBUG = "debug"
    l_INFO  = "info"
    l_WARN  = "warn"
    l_ERROR = "error"
    l_FATAL = "fatal"
)

const prefix_format = "[%s] [%s] "

//------------------------------------------------------------------------------------------------------------
// simple logger that use log package
type nativeLogger struct {
    name        string
    traceLogger *log.Logger
    debugLogger *log.Logger
    infoLogger  *log.Logger
    warnLogger  *log.Logger
    errorLogger *log.Logger
    fatalLogger *log.Logger
}

// it should be private
func newNativeLogger(name string) *nativeLogger {
    flag := log.Ldate | log.Ltime | log.Lmicroseconds
    logger := &nativeLogger{}
    logger.name = name
    logger.traceLogger = log.New(os.Stdout, fmt.Sprintf(prefix_format, name, l_TRACE), flag)
    logger.debugLogger = log.New(os.Stdout, fmt.Sprintf(prefix_format, name, l_DEBUG), flag)
    logger.infoLogger = log.New(os.Stdout, fmt.Sprintf(prefix_format, name, l_INFO), flag)
    logger.warnLogger = log.New(os.Stdout, fmt.Sprintf(prefix_format, name, l_WARN), flag)
    logger.errorLogger = log.New(os.Stderr, fmt.Sprintf(prefix_format, name, l_ERROR), flag)
    logger.fatalLogger = log.New(os.Stderr, fmt.Sprintf(prefix_format, name, l_FATAL), flag)
    return logger
}

func (logger *nativeLogger) GetName() string {
    return logger.name
}

func (logger *nativeLogger) Trace(args ...interface{}) {
    logger.traceLogger.Println(args)
}

func (logger *nativeLogger) TraceF(format string, args ...interface{}) {
    logger.traceLogger.Printf(format, args)
}

func (logger *nativeLogger) Debug(args ...interface{}) {
    logger.debugLogger.Println(args)
}

func (logger *nativeLogger) DebugF(format string, args ...interface{}) {
    logger.debugLogger.Printf(format, args)
}

func (logger *nativeLogger) Info(args ...interface{}) {
    logger.infoLogger.Println(args)
}

func (logger *nativeLogger) InfoF(format string, args ...interface{}) {
    logger.infoLogger.Printf(format, args)
}

func (logger *nativeLogger) Warn(args ...interface{}) {
    logger.warnLogger.Println(args)
}

func (logger *nativeLogger) WarnF(format string, args ...interface{}) {
    logger.warnLogger.Printf(format, args)
}

func (logger *nativeLogger) Error(args ...interface{}) {
    logger.errorLogger.Println(args)
}

func (logger *nativeLogger) ErrorF(format string, args ...interface{}) {
    logger.errorLogger.Printf(format, args)
}

func (logger *nativeLogger) Fatal(args ...interface{}) {
    logger.fatalLogger.Fatalln(args)
}

func (logger *nativeLogger) FatalF(format string, args ...interface{}) {
    logger.fatalLogger.Fatalf(format, args)
}

//------------------------------------------------------------------------------------------------------------
// factory
type nativeLoggerFactory struct {
}

func newNativeLoggerFactory() *nativeLoggerFactory {
    factory := &nativeLoggerFactory{}
    return factory
}

func (factory *nativeLoggerFactory) GetLogger(name string) Logger {
    return newNativeLogger(name)
}
