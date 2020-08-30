package rand

import "testing"

func TestSource(t *testing.T) {

	src := newSource()

	src.Seed(0) // no-op, for coverage purpose

	// validate int63
	i := src.Int63()
	if int64(uint64(i)&uint64(1<<63)) != 0 {
		t.Error("Value returned by Int63() is higher than 63-bit integer.")
	}
}
