package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmopqtuvsxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func randomNumber(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int64) string {
	var sb strings.Builder
	k := len(alphabet)
	for range n {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomMoney() int64 {
	return randomNumber(800, 10000)
}

func RandomName() string {
	return randomString(6)
}

func RandomProveedor() int64 {
	return randomNumber(1, 20)
}
