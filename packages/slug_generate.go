package services

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateRandomSlug(lengthSlug int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	slug := make([]byte, lengthSlug)
	for i := range slug {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			fmt.Println("Error generating random slug:", err)
			return "", err
		}
		slug[i] = charset[num.Int64()]
	}
	return string(slug), nil
}
