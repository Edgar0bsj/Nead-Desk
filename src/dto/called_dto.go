package dto

type CreateCalledRequest struct {
	Titulo     string `json:"titulo"`
	Descricao  string `json:"descricao"`
	Prioridade string `json:"prioridade"`
}

type UpdateCalledRequest struct {
	Titulo      string `json:"titulo"`
	Descricao   string `json:"descricao"`
	Prioridade  string `json:"prioridade"`
	AtendenteID string `json:"atendente_id"`
}
