package slf4go

import (
	"github.com/huandu/go-tls"
	"os"
	"time"
)

var pid = os.Getpid()
var startTime = time.Now()

type Log struct {
	Uptime  time.Duration `json:"uptime"`  // duration elapsed since started
	Time    time.Time     `json:"date"`    // log's time
	Context string        `json:"context"` // log's context name, like application name
	Logger  string        `json:"logger"`  // log's name, default is package

	Pid      int    `json:"pid"`      // the process id which generated this log
	Gid      int    `json:"gid"`      // the goroutine id which generated this log
	Package  string `json:"package"`  // the package-name which generated this log
	Filename string `json:"filename"` // the file-name which generated this log
	Function string `json:"function"` // the function-name which generated this log
	Line     int    `json:"line"`     // the line-number which generated this log

	Level LEVEL  `json:"level"` // log's level
	Msg   string `json:"msg"`   // log's message
}

func NewLog() *Log {
	return &Log{
		Time:   time.Now(),
		Uptime: time.Since(startTime),
		Pid:    pid,
		Gid:    int(tls.ID()),
	}
}
