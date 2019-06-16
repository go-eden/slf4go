package xlog

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Driver define the standard log print specification
type Driver interface {
	// Retrieve the name of current driver, like 'default', 'logrus'...
	Name() string

	// Print responsible of printing the standard Log
	Print(l *Log)

	// Retrieve log level of the specified logger,
	// it should return the lowest Level that could be print,
	// which can help invoker to decide whether prepare print or not.
	GetLevel(logger string) Level
}

// The default driver, just print stdout directly
type StdDriver struct {
	sync.Mutex
}

func (p *StdDriver) Name() string {
	return "default"
}

func (p *StdDriver) Print(l *Log) {
	p.Lock()
	defer p.Unlock()
	var ts = time.Unix(0, l.Time).Format("2006-01-02 15:04:05.999")
	var msg string
	if l.Format != nil {
		msg = fmt.Sprintf(*l.Format, l.Args...)
	} else {
		msg = fmt.Sprint(l.Args...)
	}
	result := fmt.Sprintf("%s [%d] [%5s] [%s] %s:%d %s\n", ts, l.Gid, l.Level.String(), l.Logger, l.Filename, l.Line, msg)
	_, _ = os.Stdout.Write([]byte(result))
}

func (p *StdDriver) GetLevel(logger string) Level {
	return LEVEL_TRACE
}
