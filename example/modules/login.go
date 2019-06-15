package modules

import "github.com/sisyphsu/slf4go"

// just use slf4go everywhere, doesn't care aboud the implement.
func Login() {
	logger := slf4go.GetLogger("login")
	logger.Info("do login")
	logger.Errorf("login result %s", "failed")
}
