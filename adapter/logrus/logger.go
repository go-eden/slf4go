package logrus

import (
    log "github.com/Sirupsen/logrus"
    "github.com/sisyphsu/slf4go"
)

//---------------------------------------------------------------------
// facade for logrus
type Logger struct {
    name  string
    entry *log.Entry
}

func newLogger(name string, logger *log.Logger) *Logger {
    result := &Logger{}
    result.name = name
    result.entry = log.NewEntry(logger).WithField("name", name)
    return result
}

func (logger *Logger) GetName() string {
    return logger.name
}

func (logger *Logger) Trace(args ...interface{}) {
    // forward to Debug
    logger.Debug(args)
}

func (logger *Logger) TraceF(format string, args ...interface{}) {
    // forward to Debug
    logger.DebugF(format, args)
}

func (logger *Logger) Debug(args ...interface{}) {
    logger.entry.Debugln(args)
}

func (logger *Logger) DebugF(format string, args ...interface{}) {
    logger.entry.Debugf(format, args)
}

func (logger *Logger) Info(args ...interface{}) {
    logger.entry.Infoln(args)
}

func (logger *Logger) InfoF(format string, args ...interface{}) {
    logger.entry.Infof(format, args)
}

func (logger *Logger) Warn(args ...interface{}) {
    logger.entry.Warnln(args)
}

func (logger *Logger) WarnF(format string, args ...interface{}) {
    logger.entry.Warnf(format, args)
}

func (logger *Logger) Error(args ...interface{}) {
    logger.entry.Errorln(args)
}

func (logger *Logger) ErrorF(format string, args ...interface{}) {
    logger.entry.Errorf(format, args)
}

func (logger *Logger) Fatal(args ...interface{}) {
    logger.entry.Fatalln(args)
}

func (logger *Logger) FatalF(format string, args ...interface{}) {
    logger.entry.Fatalf(format, args)
}

//------------------------------------------------------------------------------
// LoggerFactory for logrus
type LoggerFactory struct {
    logger *log.Logger
}

func (factory *LoggerFactory) GetLogger(name string) slf4go.Logger {
    return newLogger(name, factory.logger)
}

//-------------------------------------------------------------------------------
func NewLoggerFactory(logger *log.Logger) *LoggerFactory {
    factory := &LoggerFactory{}
    factory.logger = logger
    return factory
}
