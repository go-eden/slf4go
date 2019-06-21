package slog

import (
	"github.com/go-eden/etime"
	"github.com/huandu/go-tls"
	"strconv"
)

// Fields represents attached fileds of log
type Fields map[string]interface{}

// Merge multi fileds into new Fields instance
func NewFields(fields ...Fields) Fields {
	result := Fields{}
	for _, item := range fields {
		if item == nil {
			continue
		}
		for k, v := range item {
			result[k] = v
		}
	}
	return result
}

// Log level
type Level int

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FataLevel
)

// Retrieve Level's name
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
	case FataLevel:
		return "FATAL"
	default:
		return strconv.Itoa(int(l))
	}
}

// Log represent an log, contains all properties.
type Log struct {
	Time   int64  `json:"date"`   // log's time(us)
	Logger string `json:"logger"` // log's name, default is package

	Pid        int     `json:"pid"`         // the process id which generated this log
	Gid        int     `json:"gid"`         // the goroutine id which generated this log
	Stack      *Stack  `json:"stack"`       // the stack info of this log
	DebugStack *string `json:"debug_stack"` // the debug stack of this log

	Level  Level         `json:"level"`  // log's level
	Format *string       `json:"format"` // log's format
	Args   []interface{} `json:"args"`   // log's format args
	Fields Fields        `json:"fields"` // additional custom fields
}

// Create an new Log instance
// for better performance, caller should be provided by upper
func NewLog(level Level, pc uintptr, debugStack *string, format *string, args []interface{}, fields Fields) *Log {
	var stack *Stack
	// support first args as custom stack
	if format == nil && len(args) > 1 {
		if s, ok := args[0].(*Stack); ok {
			stack = s
			args = args[1:]
		}
	}
	// default stack
	if stack == nil {
		stack = ParseStack(pc)
	}
	return &Log{
		Time:   etime.CurrentMicrosecond(),
		Logger: stack.Package,

		Pid:        pid,
		Gid:        int(tls.ID()),
		Stack:      stack,
		DebugStack: debugStack,

		Level:  level,
		Format: format,
		Args:   args,
		Fields: fields,
	}
}

// Uptime obtain log's createTime relative to application's startTime
func (l *Log) Uptime() int64 {
	return l.Time - startTime
}
