package example_test

import (
	"time"

	slog "github.com/go-eden/slf4go"
)

func Example_custom() {
	testLog := slog.NewLogger("test")
	testLog.Info("hello world")

	time.Sleep(time.Millisecond)
}
