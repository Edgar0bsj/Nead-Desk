package dto

import (
	"nead-desk/src/domain"
)

type UserAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Nome     string          `json:"nome"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
	Role     domain.UserRole `json:"role"`
}

type UpdateUserRequest struct {
	Nome     string          `json:"nome"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
	Role     domain.UserRole `json:"role"`
}
