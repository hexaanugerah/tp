package helpers

import (
	"fmt"
	"time"
)

func CreatePaymentToken(bookingCode string, total int) string {
	return fmt.Sprintf("MID-%s-%d-%d", bookingCode, total, time.Now().Unix())
}
