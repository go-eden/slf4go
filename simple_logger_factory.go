package slf4go

type SimpleLoggerFactory struct {
    defaultLogger Logger
}

func newSimpleLoggerFactory() *SimpleLoggerFactory {
    factory := &SimpleLoggerFactory{}
    return factory
}

// get logger
func (factory *SimpleLoggerFactory) GetLogger(name string) Logger {
    return newSimpleLogger(name)
}
