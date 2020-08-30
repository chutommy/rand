package rand

import (
	crand "crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

// source is an object which satisfies the Source
// interface from the math/rand package and
// provides a random source on each function call.
type source struct{}

// uint64 uses the crypto/rand package to generate
// and return a random uint64.
func (s *source) uint64() uint64 {
	var value uint64
	binary.Read(crand.Reader, binary.BigEndian, &value)
	return value
}

// Int63 returns a random int63 by generating
// a random uint64 and shifting it by one binary
// non-negative to narrow the range.
func (s *source) Int63() int64 {
	// using bitwise operation 'AND' so
	// 1<<63 is technically going to fill
	// the first 63 digits with 1 which basically means
	// that the first digit (64th in the left to right direction)
	// is never equal to 1.
	return int64(s.uint64() & ^uint64(1<<63))
}

// Seed has no operations. Function is implemented
// only to satisfy the rand.Source interface from the math package.
// It is unnecessarily to generate new Seed since the function
// Int63() uses cryptographically secure random number generator
// from the crypto/rand package and does not depend on any Seed.
func (*source) Seed(int64) { /* noop */ }

// newSource generates a new source which implements
// the math's rand.Source interface.
func newSource() mrand.Source {
	return &source{}
}
