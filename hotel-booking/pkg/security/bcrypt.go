package security

import "crypto/sha256"

func HashPassword(password string) string {
	sum := sha256.Sum256([]byte(password))
	return string(sum[:])
}
