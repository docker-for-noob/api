package repositories

import "golang.org/x/crypto/bcrypt"

type BCryptRepository struct{}

func NewBcryptRepository() *BCryptRepository {
	return &BCryptRepository{}
}

func (m BCryptRepository) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (m BCryptRepository) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
