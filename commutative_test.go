package rapid

import (
	"testing"

	asrt "github.com/stretchr/testify/assert"
)

func Test_Idempotency(t *testing.T) {
	Check(t, func(rt *T) {
		var (
			a2 = Int().Draw(rt, "a")
			b2 = Int().Draw(rt, "b")
		)
		asrt.Equal(rt, a2-b2, b2-a2)
	})
}
