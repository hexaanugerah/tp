package helpers

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(raw string) string {
	h := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(h[:])
}

func CheckPassword(raw, hashed string) bool {
	return HashPassword(raw) == hashed
}
