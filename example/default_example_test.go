package example_test

import (
	"math/rand"
	"time"

	slog "github.com/go-eden/slf4go"
)

func Example_default() {
	log := slog.NewLogger("example")
	log.Trace("trace log")
	log.Tracef("trace time: %v", time.Now())
	log.Debug("debug log")

	slog.Debugf("debug time: %v", time.Now())

	log.Info("info log")
	log.Infof("info log: %v", time.Now())

	slog.Warn("warn log")

	log.Warnf("warn log: %v", time.Now())

	slog.Error("error log")

	log.Errorf("error time: %v", time.Now())
	log.Panic("panic log")

	slog.Panicf("panic time: %v", time.Now())

	slog.SetContextField("RequestID", rand.Uint64())
	slog.Info("<- RequestID is here")

	log.Warn("<- RequestID is here")

	time.Sleep(time.Millisecond)
}
