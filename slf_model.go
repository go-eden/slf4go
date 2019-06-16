package xlog

import (
	"github.com/huandu/go-tls"
	"strconv"
	"syscall"
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
	LEVEL_TRACE Level = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
)

// Retrieve Level's name
func (l Level) String() string {
	switch l {
	case LEVEL_TRACE:
		return "TRACE"
	case LEVEL_DEBUG:
		return "DEBUG"
	case LEVEL_INFO:
		return "INFO"
	case LEVEL_WARN:
		return "WARN"
	case LEVEL_ERROR:
		return "ERROR"
	case LEVEL_FATAL:
		return "FATAL"
	default:
		return strconv.Itoa(int(l))
	}
}

// Log represent an log, contains all properties.
type Log struct {
	Time   int64  `json:"date"`   // log's time(ms)
	Logger string `json:"logger"` // log's name, default is package

	Pid   int    `json:"pid"`   // the process id which generated this log
	Gid   int    `json:"gid"`   // the goroutine id which generated this log
	Stack *Stack `json:"stack"` // the stack info of this log

	Level  Level         `json:"level"`  // log's level
	Format *string       `json:"format"` // log's format
	Args   []interface{} `json:"args"`   // log's format args
	Fields Fields        `json:"fields"` // additional custom fields
}

// Create an new Log instance
// for better performance, caller should be provided by upper
func NewLog(level Level, pc uintptr, format *string, args []interface{}, fields Fields) *Log {
	stack := ParseStack(pc)
	return &Log{
		Time:   now(),
		Logger: stack.Package,

		Pid:   pid,
		Gid:   int(tls.ID()),
		Stack: stack,

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

// Obtain current microsecond, use syscall for better performance
func now() int64 {
	var tv syscall.Timeval
	_ = syscall.Gettimeofday(&tv)
	return int64(tv.Sec)*1e6 + int64(tv.Usec)
}
