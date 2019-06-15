package slf4go

// Logger interface
type Logger interface {
	// GetName return the name of current Logger
	GetName() string

	IsEnableTrace() bool

	IsEnableDebug() bool

	IsEnableInfo() bool

	IsEnableWarn() bool

	IsEnableError() bool

	IsEnableFatal() bool

	Trace(v ...interface{})

	Tracef(format string, v ...interface{})

	Debug(v ...interface{})

	Debugf(format string, v ...interface{})

	Info(v ...interface{})

	Infof(format string, v ...interface{})

	Warn(v ...interface{})

	Warnf(format string, v ...interface{})

	Error(v ...interface{})

	Errorf(format string, v ...interface{})

	Fatal(v ...interface{})

	Fatalf(format string, v ...interface{})
}

// LoggerFactory is Logger's provider
type LoggerFactory interface {
	GetLogger(name string) Logger
}
