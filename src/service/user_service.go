package service

import (
	"nead-desk/src/domain"
	"nead-desk/src/dto"
	"nead-desk/src/storage/ports"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	repo ports.UserStorage
}

func NewUserService(repository ports.UserStorage) *UserService {
	return &UserService{repo: repository}
}

func (c *UserService) FindUserByEmail(userEmail string) (*domain.User, error) {

	userEntity, err := c.repo.FindByEmail(userEmail)

	if err != nil {
		return nil, domain.ErrInvalidUser
	}

	return userEntity, nil
}

func (c *UserService) CreateUser(userDto *dto.UserAuthRegisterDto) (*domain.User, error) {
	entity := domain.User{
		ID:            uuid.New().String(),
		Name:          userDto.Name,
		Email:         userDto.Email,
		Password_hash: userDto.Password,
		Role:          domain.RoleUser,
		Created_at:    time.Now(),
		Updated_at:    time.Now(),
	}
	if err := c.repo.Save(&entity); err != nil {
		return nil, domain.ErrSaveFailedUser
	}

	return &entity, nil
}

func (c *UserService) FindAllUser() []*domain.User {
	users, _ := c.repo.FindAll()

	return users
}

func (c *UserService) FindByIdUser(id string) (*domain.User, error) {
	user, err := c.repo.FindByID(id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *UserService) ChangeUser(id string, userDto *dto.UserChangeResponseDto) (*domain.User, error) {
	user, err := c.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	user.Name = userDto.Name
	user.Role = domain.UserRole(userDto.Role)

	if err := c.repo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}
