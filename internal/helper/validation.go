package helper

import (
	"crypto/sha512"
	"encoding/base64"
	"hash"
)

func HashPassword(password string, salt []byte) string {
	var (
		passwordBytes             []byte    = append([]byte(password), salt...)
		sha512Hasher              hash.Hash = sha512.New()
		hashedPasswordBytes       []byte
		base64EncodedPasswordHash string
	)

	sha512Hasher.Write(passwordBytes)

	hashedPasswordBytes = sha512Hasher.Sum(nil)

	base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)

	return base64EncodedPasswordHash
}

func VerifyPassword(hashedPassword, currPassword string, salt []byte) bool {
	return hashedPassword == HashPassword(currPassword, salt)
}
