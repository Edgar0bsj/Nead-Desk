package ports

import "nead-desk/src/domain"

type CategoriesStorage interface {
	Save(user *domain.Categories) error
	FindAll() ([]*domain.Categories, error)
	FindByID(id string) (*domain.Categories, error)
	Update(User *domain.Categories) error
	Delete(id string) error
	ExistName(name string) bool
}
