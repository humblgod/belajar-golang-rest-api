package auth

import "golang.org/x/crypto/bcrypt"

func CreateHashedPassword(plainPass string) (string, error) {
	passByte := []byte(plainPass)

	// encrypt password 
	hashedPassword, err := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(plainPass, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))
	return err == nil
}
