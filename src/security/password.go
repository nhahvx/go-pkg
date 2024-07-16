package security

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// ref > https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72

func HasAndSalt(pwd string) *string {
	return hashAndSalt([]byte(pwd))
}

func hashAndSalt(pwd []byte) *string {
	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return nil
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	hasPassword := string(hash)
	return &hasPassword
}
func ComparePasswords(hashedPwd string, plainPwd string) bool {
	return comparePasswords(hashedPwd, []byte(plainPwd))
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
