package helpers

import "fmt"

func SendEmail(to, subject, body string) error {
	fmt.Printf("[Email] to=%s subject=%s\n%s\n", to, subject, body)
	return nil
}
