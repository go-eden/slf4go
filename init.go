package slf4go

// global log factory, could be replaced
var defaultFactory LoggerFactory
var definedFactory LoggerFactory

func getLoggerFactory() LoggerFactory {
    if definedFactory != nil {
        return definedFactory
    }
    if defaultFactory == nil {
        defaultFactory = newNativeLoggerFactory()
    }
    return defaultFactory
}

// support any log framework by LoggerFactory
func SetLoggerFactory(newFactory LoggerFactory) {
    if newFactory == nil {
        panic("LoggerFactory can't be nil")
    }
    definedFactory = newFactory
}

// get logger from the LoggerFactory provided, if not, use the native log.
func GetLogger(name string) Logger {
    return getLoggerFactory().GetLogger(name)
}
