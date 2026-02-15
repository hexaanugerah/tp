package helpers

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GenerateToken(userID int, role string) string {
	payload := fmt.Sprintf("%d|%s|%d", userID, role, time.Now().Unix())
	return base64.StdEncoding.EncodeToString([]byte(payload))
}

func ParseToken(token string) (int, string, bool) {
	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0, "", false
	}
	parts := strings.Split(string(decoded), "|")
	if len(parts) < 2 {
		return 0, "", false
	}
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", false
	}
	return id, parts[1], true
}
