package slog

import (
	"fmt"
	"github.com/go-eden/common/efmt"
	"io"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

type immutableLog struct {
	Time       int64 // -1 means EOF
	Logger     string
	Pid        int
	Gid        int
	Level      Level
	Stack      *Stack
	Fields     Fields
	CxtFields  Fields
	Msg        string
	DebugStack *string
}

// AsyncDriver The async driver, it should have better performance
type AsyncDriver struct {
	stdout  io.Writer
	errout  io.Writer
	channel chan immutableLog
}

func newAsyncDriver(bufsize int32) *AsyncDriver {
	t := &AsyncDriver{
		stdout:  os.Stdout,
		errout:  os.Stderr,
		channel: make(chan immutableLog, bufsize),
	}

	go t.asyncPrint()

	return t
}

func (t *AsyncDriver) Name() string {
	return "default"
}

func (t *AsyncDriver) Print(l *Log) {
	iLog := immutableLog{
		Time:       l.Time,
		Logger:     l.Logger,
		Pid:        l.Pid,
		Gid:        l.Gid,
		Stack:      l.Stack,
		Fields:     l.Fields,
		CxtFields:  l.CxtFields,
		Level:      l.Level,
		DebugStack: l.DebugStack,
	}

	// prepare message
	if l.Format != nil {
		iLog.Msg = fmt.Sprintf(*l.Format, l.Args...)
	} else {
		iLog.Msg = fmt.Sprint(l.Args...)
	}

	// send signal when buffer is empty
	select {
	case t.channel <- iLog:
	default:
		_, _ = t.errout.Write([]byte("WARNING: log lost, channel is full...\n"))
	}
}

func (t *AsyncDriver) GetLevel(_ string) Level {
	return TraceLevel
}

// print log asynchronously
func (t *AsyncDriver) asyncPrint() {
	defer func() {
		if err := recover(); err != nil {
			_, _ = t.errout.Write([]byte(fmt.Sprintf("AsyncDriver panic: %v\n%s", err, string(debug.Stack()))))
		}
	}()
	var p efmt.Printer
	for {
		l, ok := <-t.channel
		if !ok || l.Time == -1 {
			return
		}
		var fields []string
		if len(l.CxtFields) > 0 {
			for k, v := range l.CxtFields {
				fields = append(fields, fmt.Sprintf("%s=%v", k, v))
			}
		}
		if len(l.Fields) > 0 {
			for k, v := range l.Fields {
				fields = append(fields, fmt.Sprintf("%s=%v", k, v))
			}
		}
		var msg = l.Msg
		if len(fields) > 0 {
			msg = "[" + strings.Join(fields, ", ") + "] " + msg
		}
		// format and print log to stdout
		if w := t.stdout; w != nil {
			var body []byte
			ts := time.Unix(0, l.Time*1000).Format("2006-01-02 15:04:05.999999")
			if l.DebugStack != nil {
				body = p.Sprintf("%-26s [%d] [%-5s] [%s] %s:%d %s\n%s\n", ts, l.Gid, l.Level.String(), l.Logger, l.Stack.Filename, l.Stack.Line, msg, *l.DebugStack)
			} else {
				body = p.Sprintf("%-26s [%d] [%-5s] [%s] %s:%d %s\n", ts, l.Gid, l.Level.String(), l.Logger, l.Stack.Filename, l.Stack.Line, msg)
			}
			_, _ = w.Write(body)
		}
	}
}

// close this driver
func (t *AsyncDriver) close() {
	t.channel <- immutableLog{Time: -1}
}
