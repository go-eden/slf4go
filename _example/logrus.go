package main

import (
    log "github.com/Sirupsen/logrus"
    "os"
    "github.com/aellwein/slf4go"
    "github.com/aellwein/slf4go/_example/modules"
    logrus2 "github.com/aellwein/slf4go/adaptor/logrus"
)

// initialize logger, just like `log4j.properties` or `logback.xml`
func init() {
    // Log as JSON instead of the default ASCII formatter.
    log.SetFormatter(&log.JSONFormatter{})
    // Output to stdout instead of the default stderr, could also be a file.
    log.SetOutput(os.Stdout)
    // Only log the warning severity or above.
    log.SetLevel(log.WarnLevel)
    logger := log.New()
    // customize your root logger
    slf4go.SetLoggerFactory(logrus2.NewLogrusLoggerFactory(logger))
}

// use slf4go everywhere
func main() {
    logger := slf4go.GetLogger("main")
    logger.DebugF("I want %s", "Cycle Import")
    logger.ErrorF("please support it, in %02d second!", 1)
    modules.Login()
}
