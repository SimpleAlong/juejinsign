package juejinsign_test

import (
	"deme/juejinsign"
	"testing"
)

func TestSign(t *testing.T) {
	t.Run("掘金签到", func(t *testing.T) {
		sign := juejinsign.New()
		sign.AddUuid(1)
		sign.AddAid(1)
		sign.AddCookie("")
		sign.AddMsToken("")
		sign.AddBogus("")
		sign.AddToken("")
		sign.Do()
	})
}
