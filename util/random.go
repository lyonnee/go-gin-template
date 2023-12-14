package util

import (
	cryptorand "crypto/rand"
	"math/rand"
	"strconv"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := cryptorand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GetRandByAddress(address string, n uint64) string {
	return RandSeq(AddressToInt64(address), 8)
}

func RandSeq(seed int64, n uint64) string {
	rand.NewSource(seed)

	letters := []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func AddressToInt64(address string) int64 {
	address = address[28:]

	value, err := strconv.ParseInt(address, 16, 64)
	if err != nil {
		return 0
	}

	return value
}
