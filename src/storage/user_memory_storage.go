package storage

import (
	"nead-desk/src/domain"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserMemoryStorage struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

func NewUserMemoryStorage() *UserMemoryStorage {
	db := make(map[string]*domain.User)
	passHash, err := bcrypt.GenerateFromPassword([]byte("654321"), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	teste := domain.User{
		ID:            uuid.New().String(),
		Nome:          "Edgar",
		Email:         "Edgar@email.com",
		Password_hash: string(passHash),
		Role:          domain.RoleAdmin,
		Created_at:    time.Now(),
		Updated_at:    time.Now(),
	}
	db[teste.ID] = &teste

	return &UserMemoryStorage{
		users: db,
		// users: make(map[string]*domain.User),
	}

}

func (s *UserMemoryStorage) Save(user *domain.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.users[user.ID] = user
	return nil
}

func (s *UserMemoryStorage) FindByID(id string) (*domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if user, exists := s.users[id]; exists {
		return user, nil
	}

	return nil, domain.ErrUserNotFound
}

func (s *UserMemoryStorage) FindAll() ([]*domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]*domain.User, 0, len(s.users))

	for _, user := range s.users {
		users = append(users, user)
	}

	return users, nil
}

func (s *UserMemoryStorage) Update(user *domain.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[user.ID]; !exists {
		return domain.ErrUserNotFound
	}

	s.users[user.ID] = user
	return nil
}

func (s *UserMemoryStorage) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return domain.ErrUserNotFound
	}

	delete(s.users, id)
	return nil
}

func (s *UserMemoryStorage) FindByEmail(email string) (*domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var index string

	for i, v := range s.users {
		if email != v.Email {
			continue
		}
		index = i
	}

	if index == "" {
		return nil, domain.ErrUserNotFound
	}
	return s.users[index], nil
}
