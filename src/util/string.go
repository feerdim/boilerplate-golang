package util

import (
	"crypto/rand"
	"strings"
)

func GenerateRandomString(n int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	ll := len(chars)
	b := make([]byte, n)

	rand.Read(b)

	for i := range b {
		b[i] = chars[int(b[i])%ll]
	}

	return string(b)
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
