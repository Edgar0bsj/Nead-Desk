package dto

type AuthRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
