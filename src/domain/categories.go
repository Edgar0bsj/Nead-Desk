package domain

import (
	"errors"
	"time"
)

// ERROS
var (
	ErrCategoriesNotFound   = errors.New("Categoria Não Encontrado")
	ErrInvalidCategories    = errors.New("Dados da categoria invalido")
	ErrSaveFailedCategories = errors.New("Error ao Persistir dados da Categoria")
)

// STRUCK CATEGORIES
type Categories struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Is_active   bool      `json:"is_active"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
