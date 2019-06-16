package main

import log "github.com/go-eden/slf4go"

func main() {
	log := log.NewLogger("test")
	log.Info("hello world")
}
