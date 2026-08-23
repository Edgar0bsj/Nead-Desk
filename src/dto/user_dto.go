package dto

type CreateUserRequest struct {
	Nome    string `json:"nome"`
	Cargo   string `json:"cargo"`
	Unidade string `json:"unidade"`
	Email   string `json:"email"`
	Senha   string `json:"senha"`
}

type UpdateUserRequest struct {
	Nome    string `json:"nome"`
	Cargo   string `json:"cargo"`
	Unidade string `json:"unidade"`
	Email   string `json:"email"`
	Senha   string `json:"senha"`
}
