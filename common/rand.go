package common

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var letterRunesAlphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStr generate combination of number and alphabet
func RandStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RandStrAlphabet generate uppercase and lowercase alphabet  string
func RandStrAlphabet(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunesAlphabet[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RandNumber generate random integer number in specific range
func RandNumber(min int, max int) int {
	return rand.Intn(max-min) + min
}
