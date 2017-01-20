package slf4go

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

//------------------------------------------------------------------------------------------------------------
// simple logger that use log package
type NativeLogger struct {
    name        string
    traceLogger *log.Logger
    debugLogger *log.Logger
    infoLogger  *log.Logger
    warnLogger  *log.Logger
    errorLogger *log.Logger
    fatalLogger *log.Logger
}

// it should be private
func newNativeLogger(name string) *NativeLogger {
    flag := log.Ldate | log.Ltime | log.Lshortfile | log.Lmicroseconds
    logger := &NativeLogger{}
    logger.name = name
    logger.traceLogger = log.New(os.Stdout, fmt.Sprintf(PREFIX_FORMAT, name, LEVEL_TRACE), flag)
    logger.debugLogger = log.New(os.Stdout, fmt.Sprintf(PREFIX_FORMAT, name, LEVEL_DEBUG), flag)
    logger.infoLogger = log.New(os.Stdout, fmt.Sprintf(PREFIX_FORMAT, name, LEVEL_INFO), flag)
    logger.warnLogger = log.New(os.Stdout, fmt.Sprintf(PREFIX_FORMAT, name, LEVEL_WARN), flag)
    logger.errorLogger = log.New(os.Stderr, fmt.Sprintf(PREFIX_FORMAT, name, LEVEL_ERROR), flag)
    logger.fatalLogger = log.New(os.Stderr, fmt.Sprintf(PREFIX_FORMAT, name, LEVEL_FATAL), flag)
    return logger
}

func (logger *NativeLogger) GetName() string {
    return logger.name
}

func (logger *NativeLogger) Trace(args ...interface{}) {
    logger.traceLogger.Println(args)
}

func (logger *NativeLogger) TraceF(format string, args ...interface{}) {
    logger.traceLogger.Printf(format, args)
}

func (logger *NativeLogger) Debug(args ...interface{}) {
    logger.debugLogger.Println(args)
}

func (logger *NativeLogger) DebugF(format string, args ...interface{}) {
    logger.debugLogger.Printf(format, args)
}

func (logger *NativeLogger) Info(args ...interface{}) {
    logger.infoLogger.Println(args)
}

func (logger *NativeLogger) InfoF(format string, args ...interface{}) {
    logger.infoLogger.Printf(format, args)
}

func (logger *NativeLogger) Warn(args ...interface{}) {
    logger.warnLogger.Println(args)
}

func (logger *NativeLogger) WarnF(format string, args ...interface{}) {
    logger.warnLogger.Printf(format, args)
}

func (logger *NativeLogger) Error(args ...interface{}) {
    logger.errorLogger.Println(args)
}

func (logger *NativeLogger) ErrorF(format string, args ...interface{}) {
    logger.errorLogger.Printf(format, args)
}

func (logger *NativeLogger) Fatal(args ...interface{}) {
    logger.fatalLogger.Fatalln(args)
}

func (logger *NativeLogger) FatalF(format string, args ...interface{}) {
    logger.fatalLogger.Fatalf(format, args)
}

//------------------------------------------------------------------------------------------------------------
// factory
type NativeLoggerFactory struct {
}

func newNativeLoggerFactory() *NativeLoggerFactory {
    factory := &NativeLoggerFactory{}
    return factory
}

func (factory *NativeLoggerFactory) GetLogger(name string) Logger {
    return newNativeLogger(name)
}
