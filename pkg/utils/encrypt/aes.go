package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func Padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

func UnPadding(src []byte) []byte {
	n := len(src)
	unpadNum := int(src[n-1])
	return src[:n-unpadNum]
}

func Encrypt(src []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	src = Padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src
}

func Decrypt(src []byte, key []byte) []byte {
	if len(src) == 0 {
		return nil
	}
	if len(src)%16 != 0 {
		return nil
	}
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = UnPadding(src)
	return src
}
