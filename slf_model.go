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
	Uptime  int64  `json:"uptime"`  // time(ms) elapsed since started
	Time    int64  `json:"date"`    // log's time(ms)
	Context string `json:"context"` // log's context name, like application name
	Logger  string `json:"logger"`  // log's name, default is package

	Pid      int    `json:"pid"`      // the process id which generated this log
	Gid      int    `json:"gid"`      // the goroutine id which generated this log
	Package  string `json:"package"`  // the package-name which generated this log
	Filename string `json:"filename"` // the file-name which generated this log
	Function string `json:"function"` // the function-name which generated this log
	Line     int    `json:"line"`     // the line-number which generated this log

	Level  Level         `json:"level"`  // log's level
	Format *string       `json:"format"` // log's format
	Args   []interface{} `json:"args"`   // log's format args
	Fields Fields        `json:"fields"` // additional custom fields
}

// Create an new Log instance
// for better performance, caller should be provided by upper
func NewLog(level Level, pc uintptr, format *string, args []interface{}) *Log {
	stack := ParseStack(pc)
	nowUs := now() // cost 40ns
	return &Log{
		Time:    nowUs,
		Uptime:  nowUs - startTime,
		Context: context,
		Logger:  stack.pkgName,

		Pid:      pid,
		Gid:      int(tls.ID()),
		Package:  stack.pkgName,
		Filename: stack.fileName,
		Function: stack.funcName,
		Line:     stack.line,

		Level:  level,
		Format: format,
		Args:   args,
	}
}

// Obtain current microsecond, use syscall for better performance
func now() int64 {
	var tv syscall.Timeval
	_ = syscall.Gettimeofday(&tv)
	return int64(tv.Sec)*1e6 + int64(tv.Usec)
}
