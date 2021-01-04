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

// uint64 uses the crypto/rand package to generate a random uint64.
func (s *source) uint64() uint64 {
	var value uint64
	_ = binary.Read(crand.Reader, binary.BigEndian, &value)

	return value
}

// Int63 returns a random int63 by generating
// a random uint64 and shifting it by one binary
// non-negative to narrow the range.
func (s *source) Int63() int64 {
	// ^uint64(1<<63) fills the first 63 digits with 1 (and leaves the last digit with 0)
	// so the first digit of the returning value is never equal to 1.
	return int64(s.uint64() & ^uint64(1<<63))
}

// Seed has no operations. Function is implemented
// only to satisfy the rand.Source interface from the math package.
// It is unnecessarily to generate new Seed since the function
// Int63() uses cryptographically secure random number generator
// from the crypto/rand package and does not depend on any Seed.
func (*source) Seed(int64) { /* noop */ }

// NewSource generates a new source which implements the math's rand.Source interface.
func NewSource() mrand.Source {
	return &source{}
}
