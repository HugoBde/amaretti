package password

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"runtime"
	"unicode"

	"golang.org/x/crypto/argon2"
)

type Password struct {
	Hash []byte
	Salt []byte
}

// Argon2ID hashing parameters
const HashTime = 1
const HashMem = 64 * 1024
const HashOutputLen = 32

func ContainsLowerCaseChar(str string) bool {
	for _, c := range str {
		if unicode.IsLower(c) {
			return true
		}
	}
	return false
}

func ContainsUpperCaseChar(str string) bool {
	for _, c := range str {
		if unicode.IsUpper(c) {
			return true
		}
	}
	return false
}

func ContainsDigit(str string) bool {
	for _, c := range str {
		if unicode.IsDigit(c) {
			return true
		}
	}
	return false
}

func ContainsSymbol(str string) bool {
	for _, c := range str {
		if unicode.IsSymbol(c) {
			return true
		}
	}
	return false
}

func IsPasswordValid(password string) bool {
	if len(password) < 8 {
		return false
	}

	if !ContainsDigit(password) {
		return false
	}

	if !ContainsLowerCaseChar(password) {
		return false
	}

	if !ContainsUpperCaseChar(password) {
		return false
	}

	if !ContainsSymbol(password) {
		return false
	}

	return true
}

func CreatePassword(passwordStr string) (password Password, err error) {

	if password.Salt, err = createSalt(); err != nil {
		return password, err
	}

	password.Hash = argon2.IDKey([]byte(passwordStr), password.Salt, HashTime, HashMem, uint8(runtime.NumCPU()), HashOutputLen)

	return password, err
}

func createSalt() (salt []byte, err error) {

	for i := range salt {
		if val, err := rand.Int(rand.Reader, big.NewInt(255)); err != nil {
			return salt, err
		} else {
			salt[i] = byte(val.Uint64())
		}
	}

	return salt, err
}

func VerifyPassword(inputPasswordStr string, actualPassword Password) bool {
	inputPasswordHash := argon2.IDKey([]byte(inputPasswordStr), actualPassword.Salt, HashTime, HashMem, uint8(runtime.NumCPU()), HashOutputLen)

	return bytes.Equal(inputPasswordHash, actualPassword.Hash)
}
