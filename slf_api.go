package slog

const (
	rootLoggerName = "root"
	stdBufSize     = 1 << 10
)

// Driver define the standard log print specification
type Driver interface {

	// Name return the name of current driver, like 'default', 'logrus'...
	Name() string

	// Print responsible of printing the standard Log
	Print(l *Log)

	// GetLevel return log level of the specified logger,
	// it should return the lowest Level that could be print,
	// which can help invoker to decide whether prepare print or not.
	GetLevel(logger string) Level
}

type Logger interface {

	// Name obtain logger's name
	Name() string

	// Level obtain logger's level, lower will not be print
	Level() Level

	// BindFields add the specified fields into the current Logger.
	BindFields(fields Fields)

	// WithFields derive an new Logger by the specified fields from the current Logger.
	WithFields(fields Fields) Logger

	// IsTraceEnabled Whether trace of current logger enabled or not
	IsTraceEnabled() bool

	// IsDebugEnabled Whether debug of current logger enabled or not
	IsDebugEnabled() bool

	// IsInfoEnabled Whether info of current logger enabled or not
	IsInfoEnabled() bool

	// IsWarnEnabled Whether warn of current logger enabled or not
	IsWarnEnabled() bool

	// IsErrorEnabled Whether error of current logger enabled or not
	IsErrorEnabled() bool

	// IsPanicEnabled Whether panic of current logger enabled or not
	IsPanicEnabled() bool

	// IsFatalEnabled Whether fatal of current logger enabled or not
	IsFatalEnabled() bool

	// Trace record trace level's log
	Trace(v ...interface{})

	// Tracef record trace level's log with custom format.
	Tracef(format string, v ...interface{})

	// Debug record debug level's log
	Debug(v ...interface{})

	// Debugf record debug level's log with custom format.
	Debugf(format string, v ...interface{})

	// Info record info level's log
	Info(v ...interface{})

	// Infof record info level's log with custom format.
	Infof(format string, v ...interface{})

	// Warn record warn level's log
	Warn(v ...interface{})

	// Warnf record warn level's log with custom format.
	Warnf(format string, v ...interface{})

	// Error record error level's log
	Error(v ...interface{})

	// Errorf record error level's log with custom format.
	Errorf(format string, v ...interface{})

	// Panic record panic level's log
	Panic(v ...interface{})

	// Panicf record panic level's log with custom format
	Panicf(format string, v ...interface{})

	// Fatal record fatal level's log
	Fatal(v ...interface{})

	// Fatalf record fatal level's log with custom format.
	Fatalf(format string, v ...interface{})
}
