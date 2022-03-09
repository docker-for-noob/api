package ports

type PasswordValidatorRepository interface {
	VerifyPasswordStrenght(plainPassword string) error
}
