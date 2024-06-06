package util

import "math/rand"

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

func GenerateRandomString(l int) []byte {
	b := make([]byte, l)

	for i := 0; i < l; i++ {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return b
}
