package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func GenerateToken(userID int, secret string) string {
	return fmt.Sprintf("%s:%d", secret, userID)
}

func ParseToken(token, secret string) (int, error) {
	parts := strings.Split(token, ":")
	if len(parts) != 2 || parts[0] != secret {
		return 0, fmt.Errorf("invalid token")
	}
	uid, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("invalid user id")
	}
	return uid, nil
}
