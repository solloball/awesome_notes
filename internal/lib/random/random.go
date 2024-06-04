package random

import (
    "math/rand"
)

func NewRandomString(size int, seed int64) string {
    rnd := rand.New(rand.NewSource(seed)) 

    charSet := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	buf := make([]rune, size)
	for i := range buf {
		buf[i] = charSet[rnd.Intn(len(charSet))]
	}

	return string(buf)
}
