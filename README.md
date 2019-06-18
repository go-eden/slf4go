# Slf4go [![Build Status](https://travis-ci.org/go-eden/slf4go.svg?branch=master)](https://travis-ci.org/go-eden/slf4go)

Simple logger facade for Golang, inspired by `slf4j`, which focused on performance and scalability.

# Introduction

Before introducing this library, let's walk through the composition of logging library.

1. **Provide API**: like `Trade` `Debug` `Info` `Warn` `Error` etc.
2. **Collect Info**: like timestamp, stacktrace, and other context fields etc. 
3. **Format & Store**: print log into `stdout` or store it directly etc.

For most logging library, `1` and `2` is quite similar, 
but different libraries may use different logging libraries, 
if your project dependents multi libraries, the final log could be very messy.

In the `java` language, most libraries use `slf4j` as its logging facade,
you can decide to use `logback` or `log4j` etc as real logging implementation, and switch it easily.   

I believe there should have similar "facade" in golang, and i hope this library could be golang's `slf4j`.

`slf4go` focus on `1` and `2`, and will collect all information to build an integrated `Log` instance,
it expect other library provide `3` implementation, for more details, check `Driver` section.

The structure of standard `Log` is:

```go
type Log struct {
	Time   int64  `json:"date"`   // log's time(us)
	Logger string `json:"logger"` // log's name, default is package

	Pid        int     `json:"pid"`         // the process id which generated this log
	Gid        int     `json:"gid"`         // the goroutine id which generated this log
	DebugStack *string `json:"debug_stack"` // the debug stack of this log. Only for Panic and Fatal
	Stack      *Stack  `json:"stack"`       // the stack info of this log. Contains {Package, Filename, Function, Line}

	Level  Level         `json:"level"`  // log's level
	Format *string       `json:"format"` // log's format
	Args   []interface{} `json:"args"`   // log's format args
	Fields Fields        `json:"fields"` // additional custom fields
}
``` 

What need special explanation is, `slf4go` has very high performance, for more details, check `Performance` section. 

# Components

`slf4go` have several components:

+ `Log`: Log record's structure, contains `Time`, `Logger`, `Pid`, `Gid`, `Stack`, `Fields`, etc.
+ `Logger`: Provide api for `Trace`, `Debug`, `Info`, `Warn`, `Error`, `Panic`, `Fatal` etc.
+ `Driver`: It's an interface, used for decoupling `Api` and `Implementation`.
+ `Hook`: Provide a hook feature, can be used for log's synchronous callback.

For better understanding, check this chart:

<img src="./doc/structure.jpg" width="480">

# Usage

This section provides complete instructions on how to install and use `slf4go`.

## Install

Could use this command to install `slf4go`:

```bash
go get github.com/go-eden/slf4go
```

Could import `slf4go` like this:

```go
import (
	log "github.com/go-eden/slf4go"
)
```

## Use Global Logger

By default, `Slf4go` provided a global `Logger`, in most case, you can use it directly by static function, don't need any other operation.

```go
log.Debugf("debug time: %v", time.Now())
log.Warn("warn log")
log.Error("error log")
log.Panicf("panic time: %v", time.Now())
``` 

The final log is like this:

```
2019-06-16 19:35:05.167 [0] [TRACE] [main] default_example.go:12 debug time: 2019-06-16 19:35:05.167783 +0800 CST m=+0.000355435
2019-06-16 19:35:05.168 [0] [ WARN] [main] default_example.go:15 warn log
2019-06-16 19:35:05.168 [0] [ERROR] [main] default_example.go:17 error log
2019-06-16 19:35:05.168 [0] [PANIC] [main] default_example.go:20 panic time: 2019-06-16 19:35:05.168064 +0800 CST m=+0.000636402
goroutine 1 [running]:
runtime/debug.Stack(0x10aab40, 0xc0000b4100, 0x1)
	/usr/local/Cellar/go/1.12.6/libexec/src/runtime/debug/stack.go:24 +0x9d
github.com/go-eden/slf4go.Panicf(0x10cfd89, 0xe, 0xc0000b40f0, 0x1, 0x1)
	/Users/sulin/workspace/go-eden/slf4go/slf_core.go:191 +0x80
main.main()
	/Users/sulin/workspace/go-eden/slf4go/example/default_example.go:20 +0x213
```

What needs additional explanation is that `panic` and `fatal` will print `goroutine` stack automatically.

## Use Your Own Logger

You can create your own `Logger` for other purposes:

```go
log1 := log.GetLogger() // Logger's name will be package name, like "main" or "github.com/go-eden/slf4go" etc
log1.Info("hello")
log2 := log.NewLogger("test") // Logger's name will be the specified "test"
log2.Info("world")
```

The `name` of `log1` will be caller's package name, like `main` or `github.com/go-eden/slf4go` etc, it depends on where you call it.
The `name` of `log2` will be the specified `test`.

Those `name` are important:

+ It would be fill into the final log directly.
+ It would be used to check if logging is enabled.
+ It would be used to decide whether and where to record the log.

## Use Fields

You could use `BindFields` to add fields into the specified `Logger`, and use `WithFields` to create new `Logger` with the specified fields.

```go
log1 := log.GetLogger()
log1.BindFields(log.Fields{"age": 18})
log1.Debug("hell1")

log1.WithFields(log.Fields{"fav": "basketball"}).Warn("hello2")

log2 := log1.WithFields(log.Fields{"fav": "basketball"})
log2.Info("hello2")
```

The `fields` will be attached to `Log`, and finally passed to `Driver`, 
the `Driver` decided how to print or where to store them. 

## Use Level

You can setup global level by `SetLevel`, which means the lower level log will be ignored.

```go
log.SetLevel(log.LEVEL_WARN)
log.Info("no log") // will be ignored
log.Error("error")
```

Above code setup global level to be `WARN`, so `INFO` log will be ignored, 
there should have other way to config different loggers' level, 
it based on which `Driver` you use. 

You can check the specified level was enabled or not like this:

```go
l := GetLogger()
if l.IsDebugEnabled() {
    l.Debug("debug....")
}
```

In this example, `slf4go` will ask `Driver` whether accept `DEBUG` log of current logger, this process should cost few nanoseconds.

In fact, `Logger` will call `IsDebugEnabled()` in the `Debug()` function to filter unnecessary log, 
but this can't avoid the performance loss of preparing `Debug()`'s arguments, like string's `concat`. 

As a comparison, preparing data and building `Log` would cost hundreds nanoseconds.

## Use Hook

In `slf4go`, it's very easy to register log hook:

```go
log.RegisterHook(func(l *Log) {
    println(l) // you better not log again, it could be infinite loop 
})
log.Trace("are you prety?", true)
log.Debugf("are you prety? %t", true)
```  

The `RegisterHook` accept `func(*Log)` argument, and `slf4go` will broadcast all `Log` to it asynchronously.

# Driver

## Default StdDriver

By default, `Slf4go` provide a `StdDriver` as fallback, it will format `Log` and print it into `stdout` directly.

If you don't need other features, could use it directly.

## Use `slf4go-classic`

TODO

## Use `slf4go-logrus`

TODO

## Provide your own driver

TODO

# Performance

TODO

# Benefit

As we can see, golang changes very quickly, and the logger-tech isn't very mature.

Separate the logger implement from modules maybe a good idea.

if oneday you need to use `logxxx` replace `logrus`, 
do you want to change all code contains `log.info(...)`?
   
or only change `logger_init.go`?