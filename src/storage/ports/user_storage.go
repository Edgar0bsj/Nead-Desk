package ports

import "nead-desk/src/domain"

type UserStorage interface {
	Save(user *domain.User) error
	FindAll() ([]*domain.User, error)
	FindByID(id string) (*domain.User, error)
	Update(User *domain.User) error
	Delete(id string) error
}
