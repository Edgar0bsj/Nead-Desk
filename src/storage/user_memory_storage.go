package storage

import (
	"nead-desk/src/domain"
	"sync"
)

// UserMemoryStorage persiste users em um map na memória do processo.
//
// Útil para estudo, protótipos e testes: zero configuração externa.
// Os dados são perdidos ao reiniciar o servidor — comportamento esperado.
//
// UserMemoryStorage implementa userRepository implicitamente.
type UserMemoryStorage struct {
	// sync.RWMutex protege o map em acessos concorrentes.
	// Leituras (RLock) podem ocorrer em paralelo; escritas (Lock) são exclusivas.
	// Em APIs HTTP, várias goroutines (uma por requisição) acessam o storage
	// simultaneamente — sem mutex, o map causaria race condition.
	mu    sync.RWMutex
	users map[string]*domain.User
}

// NewUserMemoryStorage é o construtor da implementação em memória.
//
// Convenção Go: funções New* retornam ponteiros quando a struct
// será mutada ou compartilhada entre camadas.
func NewUserMemoryStorage() *UserMemoryStorage {
	return &UserMemoryStorage{
		users: make(map[string]*domain.User),
	}
}

func (s *UserMemoryStorage) Save(user *domain.User) error {
	s.mu.Lock()
	defer s.mu.Unlock() // defer garante unlock mesmo se houver panic futuro

	s.users[user.ID] = user
	return nil
}

func (s *UserMemoryStorage) FindByID(id string) (*domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if user, exists := s.users[id]; exists {
		// Retorna ponteiro para o user existente no map.
		// Em produção com banco, aqui viria uma cópia ou nova alocação.
		return user, nil
	}

	return nil, domain.ErrUserNotFound
}

func (s *UserMemoryStorage) FindAll() ([]*domain.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Pré-aloca slice com capacidade = len(map) para evitar realocações.
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
