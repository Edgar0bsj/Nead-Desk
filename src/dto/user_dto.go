package dto

import (
	"nead-desk/src/domain"
)

type UserAuthLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAuthRegisterDto struct {
	Name     string `validate:"required,min=3,max=50"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=25"`
}

type UpdateUserRequest struct {
	Name     string          `json:"name"`
	Email    string          `json:"email"`
	Password string          `json:"password"`
	Role     domain.UserRole `json:"role"`
}

type UserChangeResponseDto struct {
	Name string `validate:"required,min=3,max=50"`
	Role string `validate:"required,oneof=user admin"`
}
