# Go Rand Package

Clean and easy to use `math/rand` library with a true random generator.

![Rand package logo](img/rand.svg)

## Why?

Go's standard library includes two rather distinct packages for generating
random numbers.

One of these two packages is `math/rand` which has an excellent interface. It is
well documented and contains a rich and diverse method set. But it has one
catch. It is not random enough (not at all actually). The randomness of the
package depends on a seed which can't be randomly generated within the package,
and it isn't easy to generate with external packages either. Therefore, it is
commonly used with a time seed when cryptographical level of security isn't
required. This can be enough for some use cases but keep in mind that all values
can be easily replicated with the same seed value.

On the other hand there is a second random package, the `crypto/rand`. It is
perfectly able to provide a cryptographical security randomness. However, it is
hard to use and has very weak interface with only one function to generate a
random value. On top of that the data type of the returned value is
`big.Int`.

As you may expect, this package takes the best of both worlds and merges them
into a one beautiful and easy to use package with a true random generator.

#### Why isn't `rand.Seed(time.Now().UnixNano())` with `math/rand` safe enough?

If the same time seed value is provided, everything that is "randomly" generated
can be easily replicated. This implementation cannot be used if a
cryptographically secure random generation is required.

## Comparison

### `math/rand`

* excellent interface
* rich method set
* easy to use
* not cryptographically random (results can be easily regenerated)

### `crypto/rand`

* cryptographically secure random number generator
* hard to use
* limited interface

### `chutommy/rand`

* beautiful math's rand interface
* nice and rich method set
* simple and easy to use
* true RNG

## Installation

1. Download the package:

```shell
$ go get -u github.com/chutommy/rand
```

2. Import it in your code:

```go
import "github.com/chutommy/rand"
```

3. (Optional) Import `math/rand` and/or `crypto/rand`. Use aliases to
   distinguish the packages:

```go
import (
   crand "crypto/rand"
   mrand "math/rand"

   "github.com/chutommy/rand"
)
```

## Usage

```go
package main

import "github.com/chutommy/rand"

func main() {
	// get the *math.Rand with a cryptographically generated seed
	r := rand.NewRand()
	
	// random floats
	_ = r.Float32()
	_ = r.Float64()
	
	// random integers
	_ = r.Int()
	_ = r.Intn(10)
	
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
