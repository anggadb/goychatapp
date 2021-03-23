package lib

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

func GenerateSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)
	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}
	return salt
}
func HashPassword(password string, salt []byte) string {
	var passwordBytes = []byte(password)
	var sha = sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha.Write(passwordBytes)
	var hashedBytes = sha.Sum(nil)
	var encodedHashed = base64.URLEncoding.EncodeToString(hashedBytes)
	return encodedHashed
}
func PasswordMatcher(hashed, password string, salt []byte) bool {
	var passwordHashed = HashPassword(password, salt)
	return hashed == passwordHashed
}
