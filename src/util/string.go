package util

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func GenerateRandomString(n int) (string, error) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, n)
	for i := range b {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(len(letterBytes))))
		if err != nil {
			return "", err
		}

		b[i] = letterBytes[r.Int64()]
	}

	return string(b), nil
}

func FormatPhoneNumber(phoneNumber string) string {
	prefix := phoneNumber[0:3]
	if prefix != "+62" {
		phoneNumber = strings.Replace(phoneNumber, "0", "+62", 1)
	}

	return phoneNumber
}

func CapitalFirstLetter(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
