package rand

import (
	rand "math/rand"
	"reflect"
	"testing"
)

func TestRand(t *testing.T) {

	r := New()

	got := reflect.TypeOf(r).String()
	exp := reflect.TypeOf(&rand.Rand{}).String()
	if got != exp {
		t.Errorf("Function New returns %s, expected %s", got, exp)
	}
}
