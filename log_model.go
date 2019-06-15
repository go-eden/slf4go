package slf4go

import (
	"github.com/huandu/go-tls"
	"os"
	"runtime"
	"strings"
	"time"
)

var pid = os.Getpid()                 // the cached id of current process
var startTime = time.Now().UnixNano() // the start time of current process

// Log represent an log, contains all properties.
type Log struct {
	Uptime  int64  `json:"uptime"`  // time(ns) elapsed since started
	Time    int64  `json:"date"`    // log's time(ns)
	Context string `json:"context"` // log's context name, like application name
	Logger  string `json:"logger"`  // log's name, default is package

	Pid      int    `json:"pid"`      // the process id which generated this log
	Gid      int    `json:"gid"`      // the goroutine id which generated this log
	Package  string `json:"package"`  // the package-name which generated this log
	Filename string `json:"filename"` // the file-name which generated this log
	Function string `json:"function"` // the function-name which generated this log
	Line     int    `json:"line"`     // the line-number which generated this log

	Level  Level  `json:"level"`  // log's level
	Msg    string `json:"msg"`    // log's message
	Fields Fields `json:"fields"` // additional custom fields
}

// Create an new Log instance
// for better performance, caller should be provided by upper
func NewLog(level Level, pc uintptr, filename string, line int, msg string) *Log {
	if off := strings.LastIndexByte(filename, '/'); off > 0 && off < len(filename)-1 {
		filename = filename[off+1:]
	}
	pkgName, funcName := parseFunc(pc)
	now := time.Now().UnixNano()
	return &Log{
		Time:    now,
		Uptime:  now - startTime,
		Context: "",
		Logger:  pkgName,

		Pid:      pid,
		Gid:      int(tls.ID()),
		Package:  pkgName,
		Filename: filename,
		Function: funcName,
		Line:     line,

		Level:  level,
		Msg:    msg,
		Fields: Fields{},
	}
}

// Parse package and function by pc
func parseFunc(pc uintptr) (pkgName, funcName string) {
	f := runtime.FuncForPC(pc)
	if f == nil {
		return
	}
	var off int
	name := f.Name()
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' {
			break
		}
		if name[i] == '.' {
			off = i
		}
	}
	if off > 0 {
		pkgName = name[:off]
		if off < len(name)-1 {
			funcName = name[off+1:]
		}
	} else {
		pkgName = name
	}
	return
}
