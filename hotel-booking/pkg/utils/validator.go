package utils

func IsEmail(email string) bool {
	for i := 0; i < len(email); i++ {
		if email[i] == '@' {
			return true
		}
	}
	return false
}
