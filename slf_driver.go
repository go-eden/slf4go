package slog

import (
	"github.com/go-eden/common/efmt"
	"io"
	"math"
	"os"
	"runtime/debug"
	"sync/atomic"
	"time"
)

type immutableLog struct {
	printer efmt.Printer

	Time       int64
	Logger     string
	Pid        int
	Gid        int
	Stack      *Stack
	Fields     Fields
	Level      Level
	Msg        []byte
	DebugStack *string
}

// StdDriver The default driver, just print stdout directly
type StdDriver struct {
	stdout io.Writer
	errout io.Writer
	count  int32
	offset int64
	border int64
	signal chan bool
	buffer []immutableLog
}

func newStdDriver(bufsize int32) Driver {
	t := &StdDriver{
		stdout: os.Stdout,
		errout: os.Stderr,
		count:  0,
		offset: -1,
		border: int64(bufsize - 1),
		signal: make(chan bool, 2),
		buffer: make([]immutableLog, bufsize),
	}

	go t.asyncPrint()

	return t
}

func (t *StdDriver) Name() string {
	return "default"
}

func (t *StdDriver) Print(l *Log) {
	if atomic.LoadInt32(&t.count) < 0 {
		return // if closed
	}

	// allocate next offset
	var off int64
	for {
		tmp := atomic.LoadInt64(&t.offset)
		if tmp >= atomic.LoadInt64(&t.border) {
			_, _ = t.errout.Write([]byte("WARNING: log lost, channel is full...\n"))
			return
		}
		if atomic.CompareAndSwapInt64(&t.offset, tmp, tmp+1) {
			off = tmp + 1
			break
		}
	}

	// write to buffer
	iLog := &t.buffer[off%int64(len(t.buffer))]
	iLog.Time = l.Time
	iLog.Logger = l.Logger
	iLog.Pid = l.Pid
	iLog.Gid = l.Gid
	iLog.Stack = l.Stack
	iLog.Fields = l.Fields
	iLog.Level = l.Level
	iLog.DebugStack = l.DebugStack

	// prepare message
	if l.Format != nil {
		iLog.Msg = iLog.printer.Sprintf(*l.Format, l.Args...)
	} else {
		iLog.Msg = iLog.printer.Sprint(l.Args...)
	}

	// send signal when buffer is empty
	if atomic.AddInt32(&t.count, 1) == 1 {
		t.signal <- true
	}
}

func (t *StdDriver) GetLevel(_ string) Level {
	return TraceLevel
}

// print log asynchronously
func (t *StdDriver) asyncPrint() {
	defer func() {
		if err := recover(); err != nil {
			_, _ = t.errout.Write([]byte("StdDriver panic: \n" + string(debug.Stack())))
		}
	}()
	var p efmt.Printer
	for off := int64(0); ; off++ {
		for {
			if c := atomic.LoadInt32(&t.count); c == 0 {
				<-t.signal
			} else if c > 0 {
				break
			} else {
				return // break if closed
			}
		}
		l := &t.buffer[off%int64(len(t.buffer))]

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
		// clear memory
		l.Stack = nil
		l.Fields = nil
		l.DebugStack = nil

		// upgrade count and border
		atomic.AddInt32(&t.count, -1)
		atomic.AddInt64(&t.border, 1)
	}
}

func (t *StdDriver) close() {
	atomic.StoreInt32(&t.count, math.MinInt16)
	t.signal <- true
}
