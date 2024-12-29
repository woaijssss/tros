package encrypt

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1Encode(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	signa := hash.Sum(nil)
	return hex.EncodeToString(signa)
}
