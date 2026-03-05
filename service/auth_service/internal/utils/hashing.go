package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	// Generate Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return ""
	}

	return string(hashedPassword)
}

func ComparePassword(password string, hashedPassword string) bool {
	// Compare Hash Password
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)
	return err == nil
}