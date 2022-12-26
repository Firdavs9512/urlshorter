package random

import (
	"crypto/rand"
	"math/big"
)

func RandomURL(a int) string {
	text := "qwertyuiopasdfghjklzxcvbnmQWERYIOPASDFGHJKLZXCVBNM1234567890"
	randtext := make([]byte, a)
	for i := 0; i < a; i++ {
		son, err := rand.Int(rand.Reader, big.NewInt(int64(len(text))))
		if err != nil {
			panic(err)
		}
		randtext[i] = text[son.Uint64()]
	}
	return string(randtext)
}
