package service

import (
	"nead-desk/src/domain"
	"nead-desk/src/dto"
	"nead-desk/src/storage/ports"
	"time"

	"github.com/google/uuid"
)

type CategoriesService struct {
	repo ports.CategoriesStorage
}

func NewCategoriesService(repository ports.CategoriesStorage) *CategoriesService {
	return &CategoriesService{
		repo: repository,
	}
}

func (c *CategoriesService) CreateCategorie(categoriDto *dto.CreateCategorieDto) (*domain.Categories, error) {
	categorie := domain.Categories{
		ID:          uuid.New().String(),
		Name:        categoriDto.Name,
		Description: categoriDto.Description,
		Is_active:   true,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
	}

	if err := c.repo.Save(&categorie); err != nil {
		return nil, err
	}
	return &categorie, nil
}

func (c *CategoriesService) FindAllCategories() []*domain.Categories {
	allCategories, _ := c.repo.FindAll()
	return allCategories
}

func (c *CategoriesService) FindByIdCategories(id string) (*domain.Categories, error) {

	categories, err := c.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return categories, nil

}

func (c *CategoriesService) UpdateCategories(id string, cateDto *dto.UpdateCategorieDto) (*domain.Categories, error) {

	entity, err := c.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	entity.Name = cateDto.Name
	entity.Description = cateDto.Description
	entity.Updated_at = time.Now()

	if err := c.repo.Update(entity); err != nil {
		return nil, err
	}
	return entity, nil
}

func (c *CategoriesService) DeleteCategories(id string) error {
	if err := c.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (c *CategoriesService) ExisteCategoriName(name string) bool {
	return c.repo.ExistName(name)
}

func (c *CategoriesService) DisableCategory(id string) (bool, error) {
	categori, err := c.repo.FindByID(id)

	if err != nil {
		return false, err
	}

	categori.Is_active = !categori.Is_active
	if err := c.repo.Update(categori); err != nil {
		return false, err
	}

	return categori.Is_active, nil
}
