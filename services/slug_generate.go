package services

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandomSlug() string {
	lengthSlug := 8
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	slug := make([]byte, lengthSlug)
	for i := range slug {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		slug[i] = charset[num.Int64()]
	}
	return string(slug)
}
