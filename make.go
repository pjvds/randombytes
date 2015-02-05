package randombytes

import "math/rand"

// Makes a new bytes slice with the length and capacity set
// to the given length value and the content set to random bytes.
func Make(length int) []byte {
	value := make([]byte, length, length)
	for i := 0; i < length; i++ {
		value[i] = byte(rand.Intn(255))
	}

	return value
}
