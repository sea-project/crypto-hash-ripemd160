package ripemd

import (
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)

// Hash160 returns the RIPEMD160 hash of the SHA-256 HASH of the given data.
func Hash160(data []byte) []byte {
	h := sha256.Sum256(data)
	return Ripemd160h(h[:])
}

// Ripemd160h returns the RIPEMD160 hash of the given data.
func Ripemd160h(data []byte) []byte {
	h := ripemd160.New()
	h.Write(data)
	return h.Sum(nil)
}
