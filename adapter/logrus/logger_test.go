package logrus

import (
    "testing"
    "github.com/sisyphsu/slf4go"
    "github.com/Sirupsen/logrus"
)

func TestGetLogger(t *testing.T) {
    // use defined logger factory
    slf4go.SetLoggerFactory(NewLoggerFactory(logrus.New()))
    // do test
    logger := slf4go.GetLogger("test")
    logger.Debug("are you prety?", true)
    logger.DebugF("are you prety? %t", true)
    logger.Info("how old are you? ", nil)
    logger.InfoF("i'm %010d", 18)
    logger.Warn("you aren't honest! ")
    logger.WarnF("haha%02d", 1000, nil)
    logger.Error("what?")
    logger.ErrorF("what?..$%s$", "XD")
    logger.FatalF("import cycle not allowed! %s", "shit...")
    logger.Fatal("never reach here")
}
