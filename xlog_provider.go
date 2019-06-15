package xlog

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Provider define the standard log print interface
type Provider interface {
	// Retrieve the name of current provider, like 'default', 'logrus'...
	Name() string

	// Print responsible of printing the standard Log
	Print(l *Log)

	// Retrieve log level of the specified logger,
	// it should return the lowest Level that could be print,
	// which can help invoker to decide whether prepare print or not.
	GetLevel(logger string) Level
}

// The default provider, print stdout directly
type StdProvider struct {
	sync.Mutex
}

func (p *StdProvider) Name() string {
	return "default"
}

func (p *StdProvider) Print(l *Log) {
	p.Lock()
	defer p.Unlock()
	var ts = time.Unix(0, l.Time).Format("2006-01-02 15:04:05.999")
	result := fmt.Sprintf("%-29s [%-5s] [%s] %s:%d %s\n", ts, l.Level.String(), l.Logger, l.Filename, l.Line, l.Msg)
	_, _ = os.Stdout.Write([]byte(result))
}

func (p *StdProvider) GetLevel(logger string) Level {
	return LEVEL_TRACE
}
