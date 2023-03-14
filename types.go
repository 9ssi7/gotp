package gotp

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

// Digits is the number of digits in the code.
// It can be 6 or 8.
// The default is 6.
type Digits int

const (
	// DigitsSix is the number of digits in the code.
	DigitsSix Digits = 6

	// DigitsEight is the number of digits in the code.
	DigitsEight Digits = 8
)

// String returns the string representation of the number of digits.
func (d Digits) String() string {
	return fmt.Sprintf("%d", d)
}

// Length returns the number of digits.
func (d Digits) Length() int {
	return int(d)
}

// Algorithm is the hash function used to generate the code.
type Algorithm int

const (
	// AlgorithmSHA1 is the hash function used to generate the code.
	AlgorithmSHA1 Algorithm = 1

	// AlgorithmSHA256 is the hash function used to generate the code.
	AlgorithmSHA256 Algorithm = 2

	// AlgorithmSHA512 is the hash function used to generate the code.
	AlgorithmSHA512 Algorithm = 3
)

// String returns the string representation of the hash function.
func (a Algorithm) String() string {
	switch a {
	case AlgorithmSHA1:
		return "SHA1"
	case AlgorithmSHA256:
		return "SHA256"
	case AlgorithmSHA512:
		return "SHA512"
	}
	return ""
}

// Hash returns the hash function.
func (a Algorithm) Hash() func() hash.Hash {
	switch a {
	case AlgorithmSHA1:
		return sha1.New
	case AlgorithmSHA256:
		return sha256.New
	case AlgorithmSHA512:
		return sha512.New
	}
	return nil
}
