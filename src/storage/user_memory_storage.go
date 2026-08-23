package storage

import (
	"nead-desk/src/domain"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserMemoryStorage struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

func NewUserMemoryStorage() *UserMemoryStorage {
	teste := make(map[string]*domain.User)

	senhaHash, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	teste["cbebb439-80d8-4f4e-9b3c-f23abfa8b64c"] = &domain.User{
		ID:        "cbebb439-80d8-4f4e-9b3c-f23abfa8b64c",
		Nome:      "Desenvolvedor Master",
		Cargo:     domain.Programador,
		Unidade:   "Mesquita",
		Email:     "edgar@email.com",
		SenhaHash: string(senhaHash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// return &UserMemoryStorage{
	// 	users: make(map[string]*domain.User),
	// }

	return &UserMemoryStorage{
		users: teste,
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
