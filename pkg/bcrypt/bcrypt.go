package bcrypt

import "golang.org/x/crypto/bcrypt"

type Hasher interface {
	HashPassword(password string) (string, error)

	VerifyPassword(password, hashedPassword string) error
}

type hasherImpl struct {
	cost int
}

func NewHasher(cost int) Hasher {
	return &hasherImpl{cost}
}

func (h *hasherImpl) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h *hasherImpl) VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
