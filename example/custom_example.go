package main

import (
	slog "github.com/go-eden/slf4go"
	"time"
)

func main() {
	testLog := slog.NewLogger("test")
	testLog.Info("hello world")

	time.Sleep(time.Millisecond)
}
