package slf4go

import (
    "log"
    "os"
    "fmt"
    "time"
    "runtime"
    "io"
    "strings"
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

//------------------------------------------------------------------------------------------------------------
// simple l that use log package
type logger_adaptor_native struct {
    LoggerAdaptor
    tf   string
    out  io.Writer
    flag int
}

// it should be private
func newNativeLogger(name string) *logger_adaptor_native {
    logger := new(logger_adaptor_native)
    logger.name = name
    logger.out = os.Stdout
    logger.tf = "2006-01-02 15:04:05.999999999"
    logger.flag = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
    return logger
}

func (l *logger_adaptor_native) Trace(args ...interface{}) {
    if l.level <= LEVEL_TRACE {
        str := fmt.Sprint(args...)
        l.output(call_depth, l_TRACE, str)
    }
}

func (l *logger_adaptor_native) TraceF(format string, args ...interface{}) {
    if l.level <= LEVEL_TRACE {
        str := fmt.Sprintf(format, args...)
        l.output(call_depth, l_TRACE, str)
    }
}

func (l *logger_adaptor_native) Debug(args ...interface{}) {
    if l.level <= LEVEL_DEBUG {
        str := fmt.Sprint(args...)
        l.output(call_depth, l_DEBUG, str)
    }
}

func (l *logger_adaptor_native) DebugF(format string, args ...interface{}) {
    if l.level <= LEVEL_DEBUG {
        str := fmt.Sprintf(format, args...)
        l.output(call_depth, l_DEBUG, str)
    }
}

func (l *logger_adaptor_native) Info(args ...interface{}) {
    if l.level <= LEVEL_INFO {
        str := fmt.Sprint(args...)
        l.output(call_depth, l_INFO, str)
    }
}

func (l *logger_adaptor_native) InfoF(format string, args ...interface{}) {
    if l.level <= LEVEL_INFO {
        str := fmt.Sprintf(format, args...)
        l.output(call_depth, l_INFO, str)
    }
}

func (l *logger_adaptor_native) Warn(args ...interface{}) {
    if l.level <= LEVEL_WARN {
        str := fmt.Sprint(args...)
        l.output(call_depth, l_WARN, str)
    }
}

func (l *logger_adaptor_native) WarnF(format string, args ...interface{}) {
    if l.level <= LEVEL_WARN {
        str := fmt.Sprintf(format, args...)
        l.output(call_depth, l_WARN, str)
    }
}

func (l *logger_adaptor_native) Error(args ...interface{}) {
    if l.level <= LEVEL_ERROR {
        str := fmt.Sprint(args...)
        l.output(call_depth, l_ERROR, str)
    }
}

func (l *logger_adaptor_native) ErrorF(format string, args ...interface{}) {
    if l.level <= LEVEL_ERROR {
        str := fmt.Sprintf(format, args...)
        l.output(call_depth, l_ERROR, str)
    }
}

func (l *logger_adaptor_native) Fatal(args ...interface{}) {
    str := fmt.Sprint(args...)
    l.output(call_depth, l_FATAL, str)
    os.Exit(1)
}

func (l *logger_adaptor_native) FatalF(format string, args ...interface{}) {
    str := fmt.Sprintf(format, args...)
    l.output(call_depth, l_FATAL, str)
    os.Exit(1)
}

func (l *logger_adaptor_native) output(calldepth int, level, s string) error {
    var file string
    var line int
    var ts string = time.Now().Format(l.tf)
    if l.flag&(log.Lshortfile|log.Llongfile) != 0 {
        var ok bool
        _, file, line, ok = runtime.Caller(calldepth)
        if !ok {
            file = "???"
            line = 0
        }
        lastIndex := strings.LastIndex(file, "/")
        if lastIndex > 0 {
            file = file[lastIndex+1:]
        }
    }
    result := fmt.Sprintf("%-29s [%-5s] %s:%d %s\n", ts, level, file, line, s)
    _, err := l.out.Write([]byte(result))
    return err
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
