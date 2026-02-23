package security

import (
	"encoding/base64"
	"fmt"
	"time"
)

func GenerateToken(email, role string) string {
	payload := fmt.Sprintf("%s|%s|%d", email, role, time.Now().Unix())
	return base64.StdEncoding.EncodeToString([]byte(payload))
}
