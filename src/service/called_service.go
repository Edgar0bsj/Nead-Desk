package service

import (
	"nead-desk/src/domain"
	"nead-desk/src/dto"
	"nead-desk/src/storage/ports"
	"time"

	"github.com/google/uuid"
)

type CalledService struct {
	repo ports.CalledStorage
}

func NewCalledService(repository ports.CalledStorage) *CalledService {
	return &CalledService{repo: repository}
}

// Validação pendente
func (c *CalledService) CreateCalled(calledDto *dto.CreateCalledRequest, solicitanteID string) (*domain.Called, error) {

	prioridade := domain.PrioridadeParse(calledDto.Prioridade)
	if prioridade == 0 {
		return nil, domain.ErrInvalidCalled
	}

	entity := domain.Called{
		ID:            uuid.New().String(),
		Titulo:        calledDto.Titulo,
		Descricao:     calledDto.Descricao,
		Status:        domain.Aberto,
		Prioridade:    prioridade,
		SolicitanteID: solicitanteID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := c.repo.Save(&entity); err != nil {
		return nil, err
	}

	return &entity, nil
}

func (c *CalledService) GetCalledByID(id string) (*domain.Called, error) {

	return c.repo.FindByID(id)
}

func (c *CalledService) GetAllCalled() ([]*domain.Called, error) {

	return c.repo.FindAll()

}

// Validação pendente
func (c *CalledService) UpdateCalled(id string, calledDto *dto.UpdateCalledRequest) (*domain.Called, error) {
	entity, err := c.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	prioridade := domain.PrioridadeParse(calledDto.Prioridade)

	entity.Titulo = calledDto.Titulo
	entity.Descricao = calledDto.Descricao
	entity.Prioridade = prioridade
	if entity.AtendenteID == "" {
		entity.AtendenteID = calledDto.AtendenteID
		entity.Status = domain.EmAtedimento
	}
	entity.UpdatedAt = time.Now()

	if err := c.repo.Update(entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (c *CalledService) DeleteCalled(id string) error {
	return c.repo.Delete(id)
}
