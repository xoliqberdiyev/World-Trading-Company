package auth

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password string) (string, error) {
	hasPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hasPass), nil
}

func CompareHashPassword(password, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
