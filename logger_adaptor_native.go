package slf4go

import (
    "log"
    "os"
    "fmt"
)

const (
    l_TRACE    = "TRACE"
    l_DEBUG    = "DEBUG"
    l_INFO     = "INFO"
    l_WARN     = "WARN"
    l_ERROR    = "ERROR"
    l_FATAL    = "FATAL"
    call_depth = 2
)

const prefix_format = "[%s] [%s] %s"

//------------------------------------------------------------------------------------------------------------
// simple l that use log package
type logger_adaptor_native struct {
    LoggerAdaptor
    logger *log.Logger
}

// it should be private
func newNativeLogger(name string) *logger_adaptor_native {
    flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
    logger := new(logger_adaptor_native)
    logger.name = name
    logger.logger = log.New(os.Stdout, "", flag)
    return logger
}

func (l *logger_adaptor_native) Trace(args ...interface{}) {
    if l.level <= LEVEL_TRACE {
        str := fmt.Sprintln(args)
        l.logger.Output(call_depth, fmt.Sprintf(prefix_format, l_TRACE, l.name, str))
    }
}

func (l *logger_adaptor_native) TraceF(format string, args ...interface{}) {
    if l.level <= LEVEL_TRACE {
        str := fmt.Sprintf(format, args)
        l.logger.Output(call_depth, fmt.Sprintf(prefix_format, l_TRACE, l.name, str))
    }
}

func (l *logger_adaptor_native) Debug(args ...interface{}) {
    if l.level <= LEVEL_DEBUG {
        str := fmt.Sprintln(args)
        l.logger.Output(call_depth, fmt.Sprintf(prefix_format, l_DEBUG, l.name, str))
    }
}

func (l *logger_adaptor_native) DebugF(format string, args ...interface{}) {
    if l.level <= LEVEL_DEBUG {
        str := fmt.Sprintf(format, args)
        l.logger.Output(call_depth, fmt.Sprintf(prefix_format, l_DEBUG, l.name, str))
    }
}

func (l *logger_adaptor_native) Info(args ...interface{}) {
    if l.level <= LEVEL_INFO {
        str := fmt.Sprintln(args)
        l.logger.Output(call_depth, fmt.Sprintf(prefix_format, l_INFO, l.name, str))
    }
}

func (l *logger_adaptor_native) InfoF(format string, args ...interface{}) {
    if l.level <= LEVEL_INFO {
        str := fmt.Sprintf(format, args)
        l.logger.Output(call_depth, fmt.Sprintf(prefix_format, l_INFO, l.name, str))
    }
}

func (l *logger_adaptor_native) Warn(args ...interface{}) {
    if l.level <= LEVEL_WARN {
        str := fmt.Sprintln(args)
        l.logger.Output(call_depth, fmt.Sprintf(prefix_format, l_WARN, l.name, str))
    }
}

func (l *logger_adaptor_native) WarnF(format string, args ...interface{}) {
    if l.level <= LEVEL_WARN {
        str := fmt.Sprintf(format, args)
        l.logger.Output(call_depth, fmt.Sprintf(prefix_format, l_WARN, l.name, str))
    }
}

func (l *logger_adaptor_native) Error(args ...interface{}) {
    if l.level <= LEVEL_ERROR {
        str := fmt.Sprintln(args)
        l.logger.Output(call_depth, fmt.Sprintf(prefix_format, l_ERROR, l.name, str))
    }
}

func (l *logger_adaptor_native) ErrorF(format string, args ...interface{}) {
    if l.level <= LEVEL_ERROR {
        str := fmt.Sprintf(format, args)
        l.logger.Output(call_depth, fmt.Sprintf(prefix_format, l_ERROR, l.name, str))
    }
}

func (l *logger_adaptor_native) Fatal(args ...interface{}) {
    str := fmt.Sprintln(args)
    l.logger.Output(call_depth, fmt.Sprintf(prefix_format, l_FATAL, l.name, str))
    os.Exit(1)
}

func (l *logger_adaptor_native) FatalF(format string, args ...interface{}) {
    str := fmt.Sprintf(format, args)
    l.logger.Output(call_depth, fmt.Sprintf(prefix_format, l_FATAL, l.name, str))
    os.Exit(1)
}

//------------------------------------------------------------------------------------------------------------
// factory
type nativeLoggerFactory struct {
}

func newNativeLoggerFactory() LoggerFactory {
    factory := &nativeLoggerFactory{}
    return factory
}

func (factory *nativeLoggerFactory) GetLogger(name string) Logger {
    return newNativeLogger(name)
}
