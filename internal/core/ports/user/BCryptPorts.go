package ports

type BCryptRepository interface {
	CheckPasswordHash(password, hash string) bool
	HashPassword(password string) (string, error)
}
