package logrus

import (
    "testing"
    "github.com/sisyphsu/slf4go"
    log "github.com/Sirupsen/logrus"
    "os"
)

func TestGetLogrusLogger(t *testing.T) {
    // Log as JSON instead of the default ASCII formatter.
    log.SetFormatter(&log.JSONFormatter{})
    // Output to stdout instead of the default stderr, could also be a file.
    log.SetOutput(os.Stdout)
    // Only log the warning severity or above.
    log.SetLevel(log.DebugLevel)
    // use defined logger factory
    l := log.New()
    l.WriterLevel(log.DebugLevel)
    
    slf4go.SetLoggerFactory(NewLogrusLoggerFactory(l))
    
    // do test
    logger := slf4go.GetLogger("test")
    logger.SetLevel(slf4go.LEVEL_TRACE)
    logger.Debug("are you prety?", true)
    logger.DebugF("are you prety? %t", true)
    logger.Info("how old are you? ", nil)
    logger.InfoF("i'm %010d", 18)
    logger.Warn("you aren't honest! ")
    logger.WarnF("haha%02d", 1000, nil)
    logger.Trace("set level!!!!!!!")
    logger.SetLevel(slf4go.LEVEL_WARN)
    logger.Trace("set level???")
    logger.Info("this should net appear.")
    logger.Error("what?")
    logger.ErrorF("what?..$%s$", "XD")
    logger.FatalF("import cycle not allowed! %s", "shit...")
    logger.Fatal("never reach here")
}

func TestLogrusPanic(t *testing.T) {
    logger := slf4go.GetLogger("test")

    defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic as expected")
        }
    }()
    logger.Panic("this causes panic!")
}