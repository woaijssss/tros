package encrypt

import "encoding/base64"

func Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Decode(str string) string {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	}

	return string(decoded)
}
