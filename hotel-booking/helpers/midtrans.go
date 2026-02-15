package helpers

import (
	"fmt"
	"time"
)

type PaymentChargeResponse struct {
	PaymentID   string
	RedirectURL string
	Status      string
}

func CreateMidtransCharge(total int) PaymentChargeResponse {
	trxID := fmt.Sprintf("TRX-%d", time.Now().UnixNano())
	return PaymentChargeResponse{
		PaymentID:   trxID,
		RedirectURL: "/payment?trx=" + trxID,
		Status:      fmt.Sprintf("pending (simulasi) total Rp%d", total),
	}
}
