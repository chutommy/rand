---
weight: 4
title: "Usage"
draft: false
---

## Usage

```
package main

import "github.com/chutommy/rand"

func main() {
	// get the *math.Rand with a cryptographically
	// generated seed
	r := rand.NewRand()
	
	// random floats
	_ = r.Float32()
	_ = r.Float64()
	
	// random integers
	_ = r.Int()
	_ = r.Intn(10)
	
	// shuffle a slice
	ff := []string{
		"apple",
		"banana",
		"cherry",
		"kiwi",
		"lemon",
		"avocado",
		"mango",
	}
	r.Shuffle(len(ff), func(i, j int) {
		ff[i], ff[j] = ff[j], ff[i]
	})
}
```
