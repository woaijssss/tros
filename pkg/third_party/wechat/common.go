package wechat

import "math/rand"

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateNonceStr generate wechat nonce string
func GenerateNonceStr() string {
	nonceChars := make([]byte, 32)
	for index := 0; index < len(nonceChars); index++ {
		nonceChars[index] = symbols[rand.Intn(len(symbols))]
	}
	return string(nonceChars)
}
