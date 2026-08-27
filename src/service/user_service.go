package service

import (
	"nead-desk/src/domain"
	"nead-desk/src/storage/ports"
)

type UserService struct {
	repo ports.UserStorage
}

func NewUserService(repository ports.UserStorage) *UserService {
	return &UserService{repo: repository}
}

// Validação pendente
func (c *UserService) FindUserByEmail(userEmail string) (*domain.User, error) {

	userEntity, err := c.repo.FindByEmail(userEmail)

	if err != nil {
		return nil, domain.ErrInvalidUser
	}

	return userEntity, nil
}
