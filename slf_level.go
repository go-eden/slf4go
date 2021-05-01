package slog

import (
	"strconv"
	"sync/atomic"
)

type Level int

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

// String Retrieve Level's name
func (l Level) String() string {
	switch l {
	case TraceLevel:
		return "TRACE"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case PanicLevel:
		return "PANIC"
	case FatalLevel:
		return "FATAL"
	default:
		return strconv.Itoa(int(l))
	}
}

type LevelSetting struct {
	rootLevel Level
	loggerMap atomic.Value // map[string]Level
}

func (t *LevelSetting) setRootLevel(l Level) {
	t.rootLevel = l
}

func (t *LevelSetting) setLoggerLevel(levelMap map[string]Level) {
	newSettings := map[string]Level{}
	if tmp := t.loggerMap.Load(); tmp != nil {
		for k, v := range tmp.(map[string]Level) {
			newSettings[k] = v
		}
	}
	for k, v := range levelMap {
		newSettings[k] = v
	}
	t.loggerMap.Store(newSettings)
}

func (t *LevelSetting) getLoggerLevel(loggerName string) Level {
	if m := t.loggerMap.Load(); m != nil {
		v, ok := m.(map[string]Level)[loggerName]
		if ok {
			return v
		}
	}
	return t.rootLevel
}
