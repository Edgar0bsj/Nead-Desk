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

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

func (u UserRole) IsValid() bool {
	if u == RoleAdmin || u == RoleUser {
		return true
	}
	return false
}

// STRUCK USER
type User struct {
	ID            string    `json:"id"`
	Nome          string    `json:"nome"`
	Email         string    `json:"email"`
	Password_hash string    `json:"password_hash"`
	Role          UserRole  `json:"role"`
	Created_at    time.Time `json:"created_at"`
	Updated_at    time.Time `json:"updated_at"`
}
