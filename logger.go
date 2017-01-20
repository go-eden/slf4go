package slf4go

// log interface
type Logger interface {
    // get the name of logger, which was used for `GetLogger`
    GetName() string
    
    Trace(args ...interface{})
    
    TraceF(format string, args ...interface{})
    
    Debug(args ...interface{})
    
    DebugF(format string, args ...interface{})
    
    Info(args ...interface{})
    
    InfoF(format string, args ...interface{})
    
    Warn(args ...interface{})
    
    WarnF(format string, args ...interface{})
    
    Error(args ...interface{})
    
    ErrorF(format string, args ...interface{})
    
    Fatal(args ...interface{})
    
    FatalF(format string, args ...interface{})
}

// log FACTORY interface
type LoggerFactory interface {
    GetLogger(name string) Logger
}
