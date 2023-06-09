
package util

import (
	"golang.org/x/crypto/blake2b"
)

// HashBlake2b calculates the hash blake2b(b).
func HashBlake2b(buf []byte) []byte {
	hashedBuf := blake2b.Sum256(buf)
	return hashedBuf[:]
}
