package rand

import "testing"

func TestRand(t *testing.T) {

	r := New()

	// try to run math/crypto Rand's method
	_ = r.Int()
}
