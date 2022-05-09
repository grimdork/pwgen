// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// RandNonce generates a random nonce. It's not likely to be printable
func RandNonce(size int) []byte {
	buf := []byte{}
	for i := 0; i < size; i++ {
		c, err := rand.Int(rand.Reader, big.NewInt(256))
		if err != nil {
			return buf
		}

		buf = append(buf, byte(c.Int64()))
	}
	return buf
}

// validChars contains a safe subset of printable symbols.
const validChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-.,"

// RandString generates a random string suitable for passwords.
func RandString(size int) string {
	var s strings.Builder
	for i := 0; i < size; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(validChars))))
		if err != nil {
			return s.String()
		}

		s.WriteByte(validChars[n.Int64()])
	}
	return s.String()
}

// RandWords returns count words separated by hypens.
func RandWords(count int) string {
	var s strings.Builder
	upper := false
	for i := 0; i < count; i++ {
		switch i {
		case count - 1:
			if !upper {
				writeWord(&s, true)
			} else {
				writeWord(&s, false)
			}
		default:
			if d100() < 16 {
				upper = true
				writeWord(&s, true)
			} else {
				writeWord(&s, false)
			}
			s.WriteByte('-')
		}
	}
	return s.String()
}

func writeWord(s *strings.Builder, upper bool) {
	var list, upperlist []string

	if num(2) == 0 {
		list = adjectives
		upperlist = upperAdjectives
	} else {
		list = nouns
		upperlist = upperNouns
	}

	if upper || d100() < 16 {
		s.WriteString(upperlist[num(len(upperlist))])
	} else {
		s.WriteString(list[num(len(list)-1)])
	}
}

func d100() int {
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		return 99
	}

	return int(n.Int64())
}

func num(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return max
	}

	return int(n.Int64())
}
