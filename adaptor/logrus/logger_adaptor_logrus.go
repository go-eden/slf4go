package logrus

import (
    log "github.com/sirupsen/logrus"
    "github.com/sisyphsu/slf4go"
)

//---------------------------------------------------------------------
// facade for logrus
type LoggerAdaptorLogrus struct {
    slf4go.LoggerAdaptor
    entry *log.Entry
}

func newLogrusLogger(name string, logger *log.Logger) *LoggerAdaptorLogrus {
    result := new(LoggerAdaptorLogrus)
    result.SetName(name)
    result.entry = log.NewEntry(logger).WithField("name", name)
    return result
}

func (logger *LoggerAdaptorLogrus) SetLevel(l slf4go.LEVEL) {
    logger.LoggerAdaptor.SetLevel(l)
    switch l {
    case slf4go.LEVEL_TRACE:
        logger.entry.Level = log.DebugLevel
    case slf4go.LEVEL_DEBUG:
        logger.entry.Level = log.DebugLevel
    case slf4go.LEVEL_INFO:
        logger.entry.Level = log.InfoLevel
    case slf4go.LEVEL_WARN:
        logger.entry.Level = log.WarnLevel
    case slf4go.LEVEL_ERROR:
        logger.entry.Level = log.ErrorLevel
    case slf4go.LEVEL_FATAL:
        logger.entry.Level = log.FatalLevel
    }
}
func (logger *LoggerAdaptorLogrus) Trace(args ...interface{}) {
    // forward to Debug
    logger.Debug(args)
}

func (logger *LoggerAdaptorLogrus) TraceF(format string, args ...interface{}) {
    // forward to Debug
    logger.DebugF(format, args)
}

func (logger *LoggerAdaptorLogrus) Debug(args ...interface{}) {
    logger.entry.Debugln(args)
}

func (logger *LoggerAdaptorLogrus) DebugF(format string, args ...interface{}) {
    logger.entry.Debugf(format, args)
}

func (logger *LoggerAdaptorLogrus) Info(args ...interface{}) {
    logger.entry.Infoln(args)
}

func (logger *LoggerAdaptorLogrus) InfoF(format string, args ...interface{}) {
    logger.entry.Infof(format, args)
}

func (logger *LoggerAdaptorLogrus) Warn(args ...interface{}) {
    logger.entry.Warnln(args)
}

func (logger *LoggerAdaptorLogrus) WarnF(format string, args ...interface{}) {
    logger.entry.Warnf(format, args)
}

func (logger *LoggerAdaptorLogrus) Error(args ...interface{}) {
    logger.entry.Errorln(args)
}

func (logger *LoggerAdaptorLogrus) ErrorF(format string, args ...interface{}) {
    logger.entry.Errorf(format, args)
}

func (logger *LoggerAdaptorLogrus) Fatal(args ...interface{}) {
    logger.entry.Fatalln(args)
}

func (logger *LoggerAdaptorLogrus) FatalF(format string, args ...interface{}) {
    logger.entry.Fatalf(format, args)
}

//------------------------------------------------------------------------------
// LoggerFactory for logrus
type LogrusLoggerFactory struct {
    logger *log.Logger
}

func NewLogrusLoggerFactory(logger *log.Logger) slf4go.LoggerFactory {
    factory := &LogrusLoggerFactory{}
    factory.logger = logger
    return factory
}

func (factory *LogrusLoggerFactory) GetLogger(name string) slf4go.Logger {
    return newLogrusLogger(name, factory.logger)
}
