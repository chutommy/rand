package rand_test

import (
	"testing"

	"github.com/chutified/rand"
)

func TestRand(t *testing.T) {
	r1 := rand.NewRand()
	r2 := rand.NewRand()

	got1 := r1.Uint64()
	got2 := r2.Uint64()

	if got1 == got2 {
		t.Errorf("highly unlikely to generate same integers: %d", got1)
	}
}

func BenchmarkRand(b *testing.B) {
	r := rand.NewRand()

	for n := 0; n < b.N; n++ {
		r.Float64()
	}
}

func ExampleNewRand_float64() {
	r := rand.NewRand()

	// generate random float in range [0,1]
	_ = r.Float64()
	// Output:
}

func ExampleNewRand_intn() {
	r := rand.NewRand()

	fruits := []string{
		"apple",
		"banana",
		"cherry",
		"kiwi",
		"lemon",
		"avocado",
		"mango",
	}

	// select random fruit from the slice
	_ = fruits[r.Intn(len(fruits))]
	// Output:
}

func ExampleNewRand_shuffle() {
	r := rand.NewRand()

	veges := []string{
		"cucumber",
		"carrot",
		"cabbage",
		"beetroot",
		"spinach",
		"radish",
	}
	swapVeges := func(i, j int) {
		veges[i], veges[j] = veges[j], veges[i]
	}

	// shuffle the slice of vegetables
	r.Shuffle(len(veges), swapVeges)
	_ = veges
	// Output:
}
