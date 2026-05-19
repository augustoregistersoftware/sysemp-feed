package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/argon2"
)

const saltSize = 16

func NewCryptConfig(password string) (string, error) {
	_ = godotenv.Load()

	passwordEnv := os.Getenv("APP_SECRET")
	if passwordEnv == "" {
		return "", fmt.Errorf("APP_SECRET não configurado no .env")
	}

	passwordOriginal := password
	criptografado, err := Encrypt(passwordOriginal, passwordEnv)
	if err != nil {
		return "", err
	}
	return criptografado, nil
}

func deriveKey(password string, salt []byte) []byte {
	return argon2.IDKey(
		[]byte(password),
		salt,
		1,
		64*1024,
		4,
		32,
	)
}

func Encrypt(plainText string, password string) (string, error) {
	salt := make([]byte, saltSize)

	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}

	key := deriveKey(password, salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	cipherText := gcm.Seal(nil, nonce, []byte(plainText), nil)

	finalData := append(salt, nonce...)
	finalData = append(finalData, cipherText...)

	return base64.StdEncoding.EncodeToString(finalData), nil
}
