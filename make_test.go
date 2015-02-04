package randombytes_test

import (
	"testing"

	"github.com/pjvds/randombytes"
)

func TestNoAllTheSame(t *testing.T) {
	result := randombytes.Make(50)

	for i := 2; i < len(result); i++ {
		if result[i-2] == result[i-1] && result[i-1] == result[i] {
			t.Fatalf("3 bytes the same in a row, from index %v to %v", i-2, i)
		}
	}
}

func TestMakeCap(t *testing.T) {
	result := randombytes.Make(50)

	if cap(result) != 50 {
		t.Fatalf("unexpected result capacity")
	}
}

func TestMakeLen(t *testing.T) {
	result := randombytes.Make(50)

	if len(result) != 50 {
		t.Fatalf("unexpected result length")
	}
}
