package modules

import "github.com/aellwein/slf4go"

// just use slf4go everywhere, doesn't care aboud the implement.
func Login() {
    logger := slf4go.GetLogger("login")
    logger.Info("do login")
    logger.ErrorF("login result %s", "failed")
}
