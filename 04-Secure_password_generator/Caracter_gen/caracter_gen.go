package caractergen

import (
	"crypto/rand"
)

func GenRandomPassword(length int, special, numbers, up, marks bool) (string, error) {
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nums := "0123456789"
	dotComma := ".,"
	specialChars := "!@#$%^&*()-_+=<>?{}[]~"

	var chars string
	chars += lower

	if up {
		chars += upper
	}
	if numbers {
		chars += nums
	}
	if marks {
		chars += dotComma
	}
	if special {
		chars += specialChars
	}

	password := make([]byte, length)

	for i := range password {
		num, err := randomIndex(len(chars))
		if err != nil {
			return "", err
		}
		password[i] = chars[num]
	}

	return string(password), nil
}

func randomIndex(max int) (int, error) {
	b := make([]byte, 1)
	if _, err := rand.Read(b); err != nil {
		return 0, err
	}
	return int(b[0]) % max, nil
}
