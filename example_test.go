package rand

import "fmt"

func ExampleFloat64() {

	r := New()

	fmt.Printf("My favorite floating number is %g.", r.Float64())
}

func ExampleIntn() {

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

func ExampleShuffle() {

	r := New()

	vegetables := []string{
		"cucumber",
		"carrot",
		"cabbage",
		"beetroot",
		"spinach",
		"radish",
	}
	swapVeges := func(i, j int) {
		vegetables[i], vegetables[j] = vegetables[j], vegetables[i]
	}

	r.Shuffle(len(vegetables), swapVeges)
	fmt.Println(vegetables)
}
