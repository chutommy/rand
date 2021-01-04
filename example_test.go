package rand_test

import (
	"github.com/chutified/rand"
)

func ExampleRand_Float64() {
	r := rand.New()

	// generate random float in range [0,1]
	_ = r.Float64()
}

func ExampleRand_Intn() {
	r := rand.New()
	fruits := []string{
		"apple",
		"banana",
		"cherry",
		"kiwi",
		"lemon",
		"avocado",
		"mango",
	}

	// get random fruit
	_ = fruits[r.Intn(len(fruits))]
	// Output:
}

func ExampleRand_Shuffle() {
	r := rand.New()

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

	// shuffle the vegetables
	r.Shuffle(len(veges), swapVeges)
	_ = veges
	// Output:
}
