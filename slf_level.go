package slog

import (
	"strconv"
	"sync/atomic"
)

type Level int32

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
	rootLevel int32
	loggerMap atomic.Value // map[string]Level
}

func (t *LevelSetting) setRootLevel(l Level) {
	atomic.StoreInt32(&t.rootLevel, int32(l))
}

func (t *LevelSetting) setLoggerLevel(levelMap map[string]Level) {
	if rv, ok := levelMap[rootLoggerName]; ok {
		t.setRootLevel(rv)
	}
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
	rl := Level(atomic.LoadInt32(&t.rootLevel))
	if loggerName == rootLoggerName {
		return rl
	}
	if m := t.loggerMap.Load(); m != nil {
		if v, ok := m.(map[string]Level)[loggerName]; ok {
			return v
		}
	}
	return rl
}
