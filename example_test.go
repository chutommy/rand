package rand

import "fmt"

func ExampleRand_Float64() {

	r := New()

	// generate random float in range [0,1]
	_ = r.Float64()
}

func ExampleRand_Intn() {

	r := New()

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
	f := fruits[r.Intn(len(fruits))]
	fmt.Println(f)
}

func ExampleRand_Shuffle() {

	r := New()

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
	fmt.Println(veges)
}
