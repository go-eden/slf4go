package slf4go

type LEVEL int

const (
	LEVEL_TRACE LEVEL = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
)

// Logger interface
type Logger interface {
	// Get the name of l, which was used for `GetLogger`
	GetName() string
	// Setup l's level.
	SetLevel(l LEVEL)

	IsEnableTrace() bool

	IsEnableDebug() bool

	IsEnableInfo() bool

	IsEnableWarn() bool

	IsEnableError() bool

	IsEnableFatal() bool

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
