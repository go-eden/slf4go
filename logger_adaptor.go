package slf4go

type LoggerAdaptor struct {
	name  string
	level Level
}

func (a *LoggerAdaptor) SetName(name string) {
	a.name = name
}

func (a *LoggerAdaptor) GetName() string {
	return a.name
}

func (a *LoggerAdaptor) GetLevel() Level {
	return a.level
}

func (a *LoggerAdaptor) SetLevel(l Level) {
	a.level = l
}

func (a *LoggerAdaptor) IsEnableTrace() bool {
	return a.level <= LEVEL_TRACE
}

func (a *LoggerAdaptor) IsEnableDebug() bool {
	return a.level <= LEVEL_DEBUG
}

func (a *LoggerAdaptor) IsEnableInfo() bool {
	return a.level <= LEVEL_INFO
}

func (a *LoggerAdaptor) IsEnableWarn() bool {
	return a.level <= LEVEL_WARN
}

func (a *LoggerAdaptor) IsEnableError() bool {
	return a.level <= LEVEL_ERROR
}

func (a *LoggerAdaptor) IsEnableFatal() bool {
	return a.level <= LEVEL_FATAL
}
