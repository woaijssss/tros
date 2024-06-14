package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// 加密函数
func Encrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext = PKCS7Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, key)
	mode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

// 解密函数
func Decrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, key)
	mode.CryptBlocks(plaintext, ciphertext)

	return PKCS7UnPadding(plaintext), nil
}

// PKCS7填充
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// PKCS7去填充
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
