package slf4go

import (
    "testing"
    "time"
)

func TestGetLogger(t *testing.T) {
    logger := GetLogger("test")
    logger.Trace("are you prety?", true)
    logger.DebugF("are you prety? %t", true)
    logger.Info("how old are you? ", nil)
    logger.InfoF("i'm %010d", 18)
    logger.Warn("you aren't honest! ")
    logger.WarnF("haha%02d", 1000, nil)
    logger.Trace("set level to warn!!!!!")
    logger.SetLevel(LEVEL_WARN)
    logger.Trace("what?")
    logger.Info("what?")
    logger.Error("what?")
    logger.ErrorF("what?..$%s$", "XD")
    logger.FatalF("import cycle not allowed! %s", "shit...")
    logger.Fatal("never reach here")
    time.Sleep(time.Millisecond * 10)
}
