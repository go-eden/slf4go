package slf4go

import "testing"

func TestSetLoggerFactory(t *testing.T) {
    SetLoggerFactory(newSimpleLoggerFactory())
}

func TestGetLogger(t *testing.T) {
    logger := GetLogger("test")
    logger.Debug("are you prety?", true)
    logger.DebugF("are you prety? %t", true)
    logger.Info("how old are you? ", nil)
    logger.InfoF("i'm %010d", 18)
    logger.Warn("you aren't honest! ")
    logger.WarnF("haha%02d", 1000, nil)
    logger.Error("what?")
    logger.ErrorF("what?..$%s$", "XD")
}
