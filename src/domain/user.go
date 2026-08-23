package domain

import (
	"errors"
	"fmt"
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

var CargoStrings = [...]string{
	"Professor",
	"Administrador",
	"Coordenador",
	"Secretaria",
	"Ti",
	"Programador",
}

func (c Cargo) String() string {
	if int(c) >= 0 || int(c) <= len(CargoStrings) {
		return CargoStrings[int(c-1)]
	}
	return "Desconhecido"
}

func ParseCargo(value string) (Cargo, error) {
	for i, v := range CargoStrings {
		if v == value {
			return Cargo(i + 1), nil
		}
	}
	return 0, fmt.Errorf("Cargo inválido: %s", value)
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
