package encrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func Sha1Encode(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	signa := hash.Sum(nil)
	return hex.EncodeToString(signa)
}

func Sha256Encode(data string) (string, error) {
	var d []byte
	h := hmac.New(sha256.New, []byte(data))
	_, err := h.Write(d)
	if err != nil {
		return "", err
	}

	signature := Base64Encode(string(h.Sum(nil)))
	return signature, nil
}
