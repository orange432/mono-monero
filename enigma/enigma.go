package enigma

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/orange432/mono-monero/config"
)

var initializationVector = "1010101010101010"

func RandomString(n uint) string {
	rand.Seed(time.Now().UnixMicro())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func Encrypt(text string) string {
	block, err := aes.NewCipher([]byte(config.SECRET[0:32]))
	if err != nil {
		fmt.Println("Invalid key size!")
	}

	ecb := cipher.NewCBCEncrypter(block, []byte(initializationVector))

	content := PKCS5Padding([]byte(text), block.BlockSize())

	encrypted := make([]byte, len(content))

	ecb.CryptBlocks(encrypted, content)

	return base64.StdEncoding.EncodeToString(encrypted)
}

func Decrypt(text string) string {
	block, err := aes.NewCipher([]byte(config.SECRET[0:32]))
	if err != nil {
		fmt.Println("Invalid key size!")
	}

	ecb := cipher.NewCBCDecrypter(block, []byte(initializationVector))

	content, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		fmt.Println(err.Error())
	}

	decrypted := make([]byte, len(content))

	ecb.CryptBlocks(decrypted, content)

	return string(PKCS5Trimming(decrypted))

}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
