package main

import log "github.com/go-eden/slf4go"

func main() {
	testLog := log.NewLogger("test")
	testLog.Info("hello world")
}
