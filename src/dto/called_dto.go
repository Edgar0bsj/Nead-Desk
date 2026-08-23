package dto

type CreateCalledRequest struct {
	Titulo        string `json:"titulo"`
	Descricao     string `json:"descricao"`
	Prioridade    string `json:"prioridade"`
	SolicitanteID string `json:"solicitante_id"`
	AtendenteID   string `json:"atendente_id"`
}

type UpdateCalledRequest struct {
	Titulo        string `json:"titulo"`
	Descricao     string `json:"descricao"`
	Status        string `json:"status"`
	Prioridade    string `json:"prioridade"`
	SolicitanteID string `json:"solicitante_id"`
	AtendenteID   string `json:"atendente_id"`
}
