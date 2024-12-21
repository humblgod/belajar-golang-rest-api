package auth

import "golang.org/x/crypto/bcrypt"

func CreateHashedPassword(plainPass string) (string, error) {
	passByte := []byte(plainPass)

	hashedPassword, err := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(plainPass, hashedPass string) (error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))
	if err != nil {
		return err
	} 
	return nil
}
