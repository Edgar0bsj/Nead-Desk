package service

import (
	"nead-desk/src/domain"
	"nead-desk/src/dto"
	"nead-desk/src/storage/ports"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo ports.UserStorage
}

func NewUserService(repository ports.UserStorage) *UserService {
	return &UserService{repo: repository}
}

// Validação pendente
func (us *UserService) CreateUser(userDto *dto.CreateUserRequest) (*domain.User, error) {

	pHash, err := bcrypt.GenerateFromPassword([]byte(userDto.Senha), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	cargo := domain.ParseCargo(userDto.Cargo)

	entity := domain.User{
		ID:        uuid.New().String(),
		Nome:      userDto.Nome,
		Cargo:     cargo,
		Unidade:   userDto.Unidade,
		Email:     userDto.Email,
		SenhaHash: string(pHash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := us.repo.Save(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (us *UserService) GetUserByID(id string) (*domain.User, error) {

	return us.repo.FindByID(id)
}

func (us *UserService) GetAllUser() ([]*domain.User, error) {

	return us.repo.FindAll()

}

// Validação pendente
func (us *UserService) UpdateUser(id string, userDto *dto.UpdateUserRequest) (*domain.User, error) {
	entity, err := us.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	cargo := domain.ParseCargo(userDto.Cargo)

	hash, err := bcrypt.GenerateFromPassword([]byte(userDto.Senha), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	entity.Nome = userDto.Nome
	entity.Cargo = cargo
	entity.Unidade = userDto.Unidade
	entity.Email = userDto.Email
	entity.SenhaHash = string(hash)
	entity.UpdatedAt = time.Now()

	if err := us.repo.Update(entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (us *UserService) DeleteUser(id string) error {
	return us.repo.Delete(id)
}
