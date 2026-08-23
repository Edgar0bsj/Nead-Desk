package domain

import (
	"errors"
	"time"
)

// ERROS
var (
	ErrUserNotFound = errors.New("Usuario Não Encontrado")
	ErrInvalidUser  = errors.New("Dados do Usuario invalido")
)

// ENUM CARGO
type Cargo int

const (
	Professor Cargo = iota + 1
	Administrador
	Coordenador
	Secretaria
	Ti
	Programador
)

func ParseCargo(value string) Cargo {
	switch value {
	case "Professor":
		return Professor
	case "Administrador":
		return Administrador
	case "Coordenador":
		return Coordenador
	case "Secretaria":
		return Secretaria
	case "Ti":
		return Ti
	case "Programador":
		return Programador
	default:
		return 0
	}
}

func (c Cargo) String() string {
	cargos := [...]string{"Professor", "Administrador", "Coordenador", "Secretaria", "Ti", "Programador"}
	if c < Professor || c > Programador {
		return "Desconhecido"
	}
	return cargos[c]
}

// STRUCK USER
type User struct {
	ID        string    `json:"id"`
	Nome      string    `json:"nome"`
	Cargo     Cargo     `json:"cargo"`
	Unidade   string    `json:"unidade"`
	Email     string    `json:"email"`
	SenhaHash string    `json:"senha"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
