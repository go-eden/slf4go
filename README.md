# SLF4GO
Simple Logger Facade for Golang, inspired by SLF4J

# What is SLF4GO

SLF4GO is not a logger framework like logrus, it doesn't have better logger implement.

But SLF4GO could be used for separating your business code from logger framework.

# How SLF4GO does

SLF4GO provides two interface, named `Logger` and `LoggerFactory`.

`LoggerFactory` used for adapting your logger framework.

`Logger` used as log operation standard API, like `Trace`, `Debug`, `Info`, `Warn`, `Error`,
all log methods of your logger framework need be wrapped by `Logger`.

SLF4GO support logrus/log by default, you can use them directly.

After above steps,
You can customize any logger framework as your what,
and then you need adapt it as a `LoggerFactory`, make it as the global LoggerFactory by `slf4go.SetLoggerFactory`.

# Usage

If you use native log or logrus, the code below shows you how it works.

If you use other logger frameworks, you need implement `LoggerFactory` by yourself.

## Install

```bash
go get github.com/sisyphsu/slf4go
```

## Use native log as logger

```go
package main

import (
    "github.com/sisyphsu/slf4go"
    "github.com/sisyphsu/slf4go/example/modules"
)

// doesn't need initialize

// use slf4j everywhere
func main() {
    logger := slf4go.GetLogger("main")
    logger.DebugF("I want %s", "Cycle Import")
    logger.ErrorF("please support it, in %02d second!", 1)
    modules.Login()
}

// just use slf4go everywhere, doesn't care aboud the implement.
func Login() {
    logger := slf4go.GetLogger("login")
    logger.Info("do login")
    logger.ErrorF("login result %s", "failed")
}
```

## Use logrus as logger


```go
package main

import (
    log "github.com/sirupsen/logrus"
    "os"
    "github.com/sisyphsu/slf4go"
    "github.com/sisyphsu/slf4go/adapter/logrus"
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
    slf4go.SetLoggerFactory(logrus.NewLoggerFactory(logger))
}

// use slf4go everywhere
func main() {
    logger := slf4go.GetLogger("main")
    logger.DebugF("I want %s", "Cycle Import")
    logger.ErrorF("please support it, in %02d second!", 1)
}
```

# Benefit

As we can see, golang changes very quickly, and the logger-tech isn't very mature.

Separate the logger implement from modules maybe a good idea.

if oneday you need to use `logxxx` replace `logrus`,
do you want to change all code contains `log.info(...)`?

or only change `logger_init.go`?
