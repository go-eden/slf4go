package slog

import (
	"fmt"
	"github.com/go-eden/common/efmt"
	"io"
	"os"
	"runtime/debug"
	"time"
)

type immutableLog struct {
	Time       int64 // -1 means EOF
	Logger     string
	Pid        int
	Gid        int
	Stack      *Stack
	Fields     Fields
	Level      Level
	Msg        string
	DebugStack *string
}

// StdDriver The default driver, just print stdout directly
type StdDriver struct {
	stdout  io.Writer
	errout  io.Writer
	channel chan immutableLog
}

func newStdDriver(bufsize int32) *StdDriver {
	t := &StdDriver{
		stdout:  os.Stdout,
		errout:  os.Stderr,
		channel: make(chan immutableLog, bufsize),
	}

	go t.asyncPrint()

	return t
}

func (t *StdDriver) Name() string {
	return "default"
}

func (t *StdDriver) Print(l *Log) {
	iLog := immutableLog{
		Time:       l.Time,
		Logger:     l.Logger,
		Pid:        l.Pid,
		Gid:        l.Gid,
		Stack:      l.Stack,
		Fields:     l.Fields,
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

func (t *StdDriver) GetLevel(_ string) Level {
	return TraceLevel
}

// print log asynchronously
func (t *StdDriver) asyncPrint() {
	defer func() {
		if err := recover(); err != nil {
			_, _ = t.errout.Write([]byte(fmt.Sprintf("StdDriver panic: %v\n%s", err, string(debug.Stack()))))
		}
	}()
	var p efmt.Printer
	for {
		l, ok := <-t.channel
		if !ok || l.Time == -1 {
			return
		}

		// format and print log to stdout
		if w := t.stdout; w != nil {
			var body []byte
			ts := time.Unix(0, l.Time*1000).Format("2006-01-02 15:04:05.999999")
			if l.DebugStack != nil {
				body = p.Sprintf("%-26s [%d] [%-5s] [%s] %s:%d %s\n%s\n", ts, l.Gid, l.Level.String(), l.Logger, l.Stack.Filename, l.Stack.Line, l.Msg, *l.DebugStack)
			} else {
				body = p.Sprintf("%-26s [%d] [%-5s] [%s] %s:%d %s\n", ts, l.Gid, l.Level.String(), l.Logger, l.Stack.Filename, l.Stack.Line, l.Msg)
			}
			_, _ = w.Write(body)
		}
	}
}

// close this driver
func (t *StdDriver) close() {
	t.channel <- immutableLog{Time: -1}
}
