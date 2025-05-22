package utils_google

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomPassword() string {
	b := make([]byte, 12)
	_, err := rand.Read(b)
	if err != nil {
		return "defaultpassword"
	}
	return base64.URLEncoding.EncodeToString(b)
}
