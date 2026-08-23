package ports

import "nead-desk/src/domain"

type CalledStorage interface {
	Save(called *domain.Called) error
	FindAll() ([]*domain.Called, error)
	FindByID(id string) (*domain.Called, error)
	Update(called *domain.Called) error
	Delete(id string) error
}
