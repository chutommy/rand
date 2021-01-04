package rand_test

import (
	"testing"

	"github.com/chutified/rand"
)

func TestRand(t *testing.T) {
	r1 := rand.New()
	r2 := rand.New()

	got1 := r1.Uint64()
	got2 := r2.Uint64()

	if got1 == got2 {
		t.Errorf("highly unlikely to generate same integers: %d", got1)
	}
}

func BenchmarkRand(b *testing.B) {
	r := rand.New()

	for n := 0; n < b.N; n++ {
		r.Float64()
	}
}
