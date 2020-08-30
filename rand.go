package rand

import (
	mrand "math/rand"
)

// Rand is a controller of the true random number generator (TRND).
type Rand struct {
	Mrand *mrand.Rand
}

// New generates a new Rand with the source of the RNG.
func New() *Rand {
	return &Rand{
		Mrand: mrand.New(newSource()),
	}
}

// Float32 returns a random float32 in the range [0;1).
func (r *Rand) Float32() float32 {
	return r.Mrand.Float32()
}

// Float64 returns a random float64 in the range [0;1).
func (r *Rand) Float64() float64 {
	return r.Mrand.Float64()
}

// Int returns a non-negative random int.
func (r *Rand) Int() int {
	return r.Mrand.Int()
}

// Int31 returns, as an int32, a non-negative random 31-bit integer.
func (r *Rand) Int31() int32 {
	return r.Mrand.Int31()
}

// Int31n returns, as an int32, a non-negative random 31-bit integer in range [0,n).It panics if n<=0.
func (r *Rand) Int31n(n int32) int32 {
	return r.Mrand.Int31n(n)
}

// Int63 returns, as an int63, a non-negative random 63-bit integer.
func (r *Rand) Int63() int64 {
	return r.Mrand.Int63()
}

// Int63n returns, as an int64, a non-negative random 63-bit integer in range [0,n).It panics if n<=0.
func (r *Rand) Int63n() int64 {
	return r.Mrand.Int63()
}

// Intn returns, as an int, a non-negative random number in range [0,n).It panics if n<=0.
func (r *Rand) Intn(n int) int {
	return r.Mrand.Intn(n)
}

// Perm returns, as a slice of n ints, a random permutation of the integers in range [0,n).
func (r *Rand) Perm(n int) []int {
	return r.Mrand.Perm(n)
}

// Read writes len(p) random bytes into p. It returns len(p) and a nil error. Read should not be called concurrently with any other Rand method.
func (r *Rand) Read(p []byte) (int, error) {
	return r.Mrand.Read(p)
}

// Shuffle randomizes the order of elements. n is the number of elements. Shuffle panics if n < 0. swap swaps the elements with indexes i and j.
func (r *Rand) Shuffle(n int, swap func(i, j int)) {
	r.Mrand.Shuffle(n, swap)
}

// Uint32 returns, as an uint32, a -random 32-bit integer.
func (r *Rand) Uint32() uint32 {
	return r.Mrand.Uint32()
}

// Uint64 returns, as an uint64, a -random 64-bit integer.
func (r *Rand) Uint64() uint64 {
	return r.Mrand.Uint64()
}
