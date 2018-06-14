package security

import (
	"log"

	"github.com/raja/argon2pw"
	// "github.com/dwin/goArgonPass"
)

// VerifyPassword : check password (2nd argument) against hash (1st argument)
func VerifyPassword(hashedPassword string, password string) (bool, error) {

	// Test correct password in constant time
	valid, err := argon2pw.CompareHashWithPassword(hashedPassword, password)
	log.Printf("The password validity is %t against the hash", valid)

	return valid, err
}

// GenerateHash : generate hashed password
func GenerateHash(password string) (string, error) {

	// Generate a hashed password
	hashedPassword, err := argon2pw.GenerateSaltedHash(password)
	if err != nil {
		log.Printf("Hash generated returned error: %v", err)
	}
	log.Println("Hash Output: ", hashedPassword)

	return hashedPassword, err
}
