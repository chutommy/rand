package rand_test

import (
	"testing"

	"github.com/chutified/rand"
)

func TestSource(t *testing.T) {
	src := rand.NewSource()

	src.Seed(0) // no-op, for coverage purpose

	// validate int63
	i := src.Int63()
	if int64(uint64(i)&uint64(1<<63)) != 0 {
		t.Error("Value returned by Int63() does not fit into the 63-bit integer.")
	}
}

func BenchmarkSource(b *testing.B) {
	for n := 0; n < b.N; n++ {
		src := rand.NewSource()
		src.Int63()
	}
}
