package main

import slog "github.com/go-eden/slf4go"

func main() {
	testLog := slog.NewLogger("test")
	testLog.Info("hello world")
}
