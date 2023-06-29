// BSD 3-Clause License
//
// Copyright (c) 2022, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version    :  0.1
//

package random

import (
	"math/rand"
	"time"
)

type (
	Random struct {
		Length         int
		UseSpecialChar bool
	}
)

const (
	numeric      = "1234567890"
	alpha        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	specialChars = "@#$%^*(){}[]<>/\\"
)

func New(useSpecialChar bool, length int) *Random {
	return &Random{
		Length:         length,
		UseSpecialChar: useSpecialChar,
	}
}

func (r *Random) Generate() string {
	charSet := alphaNumeric
	if r.UseSpecialChar {
		charSet = charSet + specialChars
	}

	random := make([]byte, r.Length)
	set := rand.New(rand.NewSource(time.Now().UnixNano()))
	for cnt, _ := range random {
		random[cnt] = charSet[set.Intn(len(charSet))]
	}
	return string(random)
}

func (r *Random) Alpha() string {
	charSet := alpha
	random := make([]byte, r.Length)
	set := rand.New(rand.NewSource(time.Now().UnixNano()))
	for cnt, _ := range random {
		random[cnt] = charSet[set.Intn(len(charSet))]
	}
	return string(random)
}

func (r *Random) Numeric() string {
	charSet := numeric
	random := make([]byte, r.Length)
	set := rand.New(rand.NewSource(time.Now().UnixNano()))
	for cnt, _ := range random {
		random[cnt] = charSet[set.Intn(len(charSet))]
	}
	return string(random)
}

func (r *Random) SetLength(length int) {
	r.Length = length
}
