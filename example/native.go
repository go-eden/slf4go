package main

import (
    "github.com/aellwein/slf4go"
    "github.com/aellwein/slf4go/example/modules"
)

// doesn't need initialize

// use slf4go everywhere
func main() {
    logger := slf4go.GetLogger("main")
    logger.DebugF("I want %s", "Cycle Import")
    logger.ErrorF("please support it, in %02d second!", 1)
    modules.Login()
}
