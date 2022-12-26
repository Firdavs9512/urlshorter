package test

import (
	"testing"

	"github.com/firdavs9512/urlshorter/services/random"
)

// Random file testing for create random url
func TestRandomURL(t *testing.T) {

	text := random.RandomURL(10)

	if len(text) != 10 {
		t.Error("Random files error!")
		return
	}
}
