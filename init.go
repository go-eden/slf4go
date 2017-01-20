package slf4go

// global log factory, could be replaced
var factory LoggerFactory

func init() {
    factory = newSimpleLoggerFactory()
}

// support any log framework by LoggerFactory
func SetLoggerFactory(newFactory LoggerFactory) {
    if newFactory == nil {
        panic("LoggerFactory can't be nil")
    }
    factory = newFactory
}

// get logger from the LoggerFactory provided, if not, use the native log.
func GetLogger(name string) Logger {
    return factory.GetLogger(name)
}
