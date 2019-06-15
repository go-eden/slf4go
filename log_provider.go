package slf4go

// LogProvider define the standard log print interface
type LogProvider interface {
	// Retrieve the name of current provider, like 'default', 'logrus'...
	Name() string

	// Print responsible of printing the standard Log
	Print(l *Log)

	// Retrieve log level of the specified logger,
	// it should return the lowest Level that could be print,
	// which can help invoker to decide whether prepare print or not.
	GetLevel(logger string) Level
}
