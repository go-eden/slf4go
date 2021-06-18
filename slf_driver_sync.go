package slog

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// StdDriver The default driver, just print stdout directly
type StdDriver struct {
	sync.Mutex
}

func (p *StdDriver) Name() string {
	return "default"
}

func (p *StdDriver) Print(l *Log) {
	p.Lock()
	defer p.Unlock()
	var ts = time.Unix(0, l.Time*1000).Format("2006-01-02 15:04:05.999999")
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
	var msg string
	if l.Format != nil {
		msg = fmt.Sprintf(*l.Format, l.Args...)
	} else {
		msg = fmt.Sprint(l.Args...)
	}
	if len(fields) > 0 {
		msg = "[" + strings.Join(fields, ", ") + "] " + msg
	}
	var result string
	if l.DebugStack != nil {
		result = fmt.Sprintf("%-26s [%d] [%-5s] [%s] %s:%d %s\n%s\n", ts, l.Gid, l.Level.String(), l.Logger, l.Stack.Filename, l.Stack.Line, msg, *l.DebugStack)
	} else {
		result = fmt.Sprintf("%-26s [%d] [%-5s] [%s] %s:%d %s\n", ts, l.Gid, l.Level.String(), l.Logger, l.Stack.Filename, l.Stack.Line, msg)
	}
	_, _ = os.Stdout.Write([]byte(result))
}

func (p *StdDriver) GetLevel(_ string) Level {
	return TraceLevel
}
