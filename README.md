# Go Rand Package

Clean and easy to use `math/rand` library clone with a true random generator.

![Rand package logo](img/rand.svg)

## Table of contents

  - [Go Rand Package](#go-rand-package)
    - [Contents](#table-of-contents)
    - [Why another rand?](#wait-why-another-rand-package)
    - [Comparison](#comparison-of-the-existing-rand-packages)
      - [crypto/rand](#cryptorand)
      - [math/rand](#mathrand)
      - [chutommy/rand](#chutommyrand)
    - [Installation](#installation)
    - [Examples](#examples)

## Why another rand package?

Well, to put it simply, in the Go's standard library the `math/rand` is not random
enough (not at all actually) and the `crypto/rand` has a terrifying and really weak
user interface (compared with the `math/rand`). This package takes the best of both worlds
and merges them into a one beautiful and easy to use package with a truly random generator.

## Comparison of the existing rand packages

### `crypto/rand`

  * cryptographically secure random number generator
  + hard to use
  + ugly and weak interface

### `math/rand`

  * an excellent interface 
  * rich method set
  * easy to use
  * great documentation 
  + not cryptographically random (results can be easily regenerated)

#### Why does not `rand.Seed(time.Now().UnixNano())` make the package `math/rand` safe enough?

If the same seed value is provided, everything that was randomly generated can be replicated
and therefore this implementation can not be used when a cryptographically secure random
generation is required.

### `chutommy/rand`

  * beautiful math's rand interface 
  * nice and rich method set
  * simple and easy to use
  * true random generator

## Installation

To install `chutommy/rand` package, you need to install Go and set your Go workspace and an environment first.

1. Then you can use the Go command bellow to install `chutommy/rand`:

```shell
$ go get -u github.com/chutommy/rand
```

2. Import the package in your code.

```go
import "github.com/chutommy/rand"
```

3. (Optional) Import `math/rand` and/or `crypto/rand`. Aliases are needed to distinguish the packages.

```go
import (
    crand "crypto/rand"
    mrand "math/rand"
  
    rand "github.com/chutommy/rand"
)
```
## Examples

```go
package main

import "github.com/chutommy/rand"

func main() {
	r := rand.NewRand()
	
	// generate a random float
	_ = r.Float32()
	_ = r.Float64()
	
	// generate a random integer
	_ = r.Int()
	_ = r.Intn(10) // (0~9)
	
	// shuffle a slice
	fruits := []string{
		"apple",
		"banana",
		"cherry",
		"kiwi",
		"lemon",
		"avocado",
		"mango",
	}
	r.Shuffle(len(fruits), func(i, j int) {
		fruits[i], fruits[j] = fruits[j], fruits[i]
	})
}
```
