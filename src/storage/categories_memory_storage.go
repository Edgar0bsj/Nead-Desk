package storage

import (
	"nead-desk/src/domain"
	"sync"
	"time"

	"github.com/google/uuid"
)

type CategoriesMemoryStorage struct {
	mu       sync.RWMutex
	categori map[string]*domain.Categories
}

func NewCategoriesMemoryStorage() *CategoriesMemoryStorage {
	db := make(map[string]*domain.Categories)

	teste := domain.Categories{
		ID:          uuid.New().String(),
		Name:        "Rede",
		Description: "Problemas relacionados à rede.",
		Is_active:   true,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}
	db[teste.ID] = &teste

	return &CategoriesMemoryStorage{
		categori: db,
		// categori: make(map[string]*domain.Categories),
	}

}

func (c *CategoriesMemoryStorage) Save(cate *domain.Categories) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.categori[cate.ID] = cate
	return nil
}

func (c *CategoriesMemoryStorage) FindByID(id string) (*domain.Categories, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if entity, exists := c.categori[id]; exists {
		return entity, nil
	}

	return nil, domain.ErrCategoriesNotFound
}

func (c *CategoriesMemoryStorage) FindAll() ([]*domain.Categories, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	sliceCater := make([]*domain.Categories, 0, len(c.categori))

	for _, cate := range c.categori {
		sliceCater = append(sliceCater, cate)
	}

	return sliceCater, nil
}

func (c *CategoriesMemoryStorage) Update(cate *domain.Categories) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.categori[cate.ID]; !exists {
		return domain.ErrCategoriesNotFound
	}

	c.categori[cate.ID] = cate
	return nil
}

func (c *CategoriesMemoryStorage) Delete(id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.categori[id]; !exists {
		return domain.ErrCategoriesNotFound
	}

	delete(c.categori, id)
	return nil
}

func (c *CategoriesMemoryStorage) ExistName(name string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var result bool

	for _, v := range c.categori {
		if v.Name != name {
			continue
		}
		result = true
	}
	return result
}
