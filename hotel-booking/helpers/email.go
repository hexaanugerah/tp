package helpers

import "log"

func SendEmail(to, subject, body string) {
	log.Printf("[email-simulated] to=%s subject=%s body=%s", to, subject, body)
}
