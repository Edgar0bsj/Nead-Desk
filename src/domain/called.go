package domain

import (
	"errors"
	"time"
)

// Enum Status
type Status int

const (
	Aberto Status = iota + 1
	EmAtedimento
	Aguardando
	Resolvido
)

func StatusParse(value string) Status {
	switch value {
	case "Aberto":
		return Aberto
	case "EmAtedimento":
		return EmAtedimento
	case "Aguardando":
		return Aguardando
	case "Resolvido":
		return Resolvido
	default:
		return 0
	}
}

// Enum Prioridade
type Prioridade int

const (
	Baixa Prioridade = iota + 1
	Media
	Alta
	Critica
)

func (p Prioridade) String() string {
	prioridadeStrings := [...]string{
		"Baixa",
		"Media",
		"Alta",
		"Critica",
	}
	if int(p) > 0 && int(p) <= len(prioridadeStrings) {
		return prioridadeStrings[int(p)-1]
	}

	return "Desconhecido"
}

func PrioridadeParse(value string) Prioridade {
	switch value {
	case "Baixa":
		return Baixa
	case "Media":
		return Media
	case "Alta":
		return Alta
	case "Critica":
		return Critica
	default:
		return 0
	}
}

// Domain Called
type Called struct {
	ID            string     `json:"id"`
	Titulo        string     `json:"titulo"`
	Descricao     string     `json:"descricao"`
	Status        Status     `json:"status"`
	Prioridade    Prioridade `json:"prioridade"`
	SolicitanteID string     `json:"solicitante_id"`
	AtendenteID   string     `json:"atendente_id"`
	CreatedAt     time.Time  `json:"create_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

// ERROS
var (
	ErrCalledNotFound = errors.New("Chamado Não Encontrado")
	ErrInvalidCalled  = errors.New("Dados do Chamado invalido")
)
