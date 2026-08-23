package storage

import (
	"nead-desk/src/domain"
	"sync"
)

type CalledMemoryStorage struct {
	mu     sync.RWMutex
	called map[string]*domain.Called
}

func NewCalledMemoryStorage() *CalledMemoryStorage {
	return &CalledMemoryStorage{
		called: make(map[string]*domain.Called),
	}
}

func (c *CalledMemoryStorage) Save(called *domain.Called) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.called[called.ID] = called
	return nil
}

func (c *CalledMemoryStorage) FindByID(id string) (*domain.Called, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if called, exists := c.called[id]; exists {
		return called, nil
	}

	return nil, domain.ErrCalledNotFound
}

func (c *CalledMemoryStorage) FindAll() ([]*domain.Called, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	calledList := make([]*domain.Called, 0, len(c.called))

	for _, call := range c.called {
		calledList = append(calledList, call)
	}

	return calledList, nil
}

func (c *CalledMemoryStorage) Update(called *domain.Called) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.called[called.ID]; !exists {
		return domain.ErrCalledNotFound
	}

	c.called[called.ID] = called
	return nil
}

func (c *CalledMemoryStorage) Delete(id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.called[id]; !exists {
		return domain.ErrCalledNotFound
	}

	delete(c.called, id)
	return nil
}
