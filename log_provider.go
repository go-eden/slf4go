package xlog

import (
	"fmt"
	"io"
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
type WriterProvider struct {
	sync.Mutex
	writer io.Writer
}

func (p *WriterProvider) Name() string {
	return "default"
}

func (p *WriterProvider) Print(l *Log) {
	p.Lock()
	defer p.Unlock()
	var ts = time.Unix(0, l.Time).Format("2006-01-02 15:04:05.999")
	result := fmt.Sprintf("%-29s [%-5s] [%s] %s:%d %s\n", ts, l.Level.String(), l.Logger, l.Filename, l.Line, l.Msg)
	_, _ = p.writer.Write([]byte(result))
}

func (p *WriterProvider) GetLevel(logger string) Level {
	return LEVEL_TRACE
}
