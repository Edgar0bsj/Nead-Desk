package dto

type CreateCategorieDto struct {
	Name        string `validate:"required,min=2,max=50"`
	Description string `validate:"required,min=10,max=250"`
}

type UpdateCategorieDto struct {
	Name        string `validate:"required,min=2,max=50"`
	Description string `validate:"required,min=10,max=250"`
}
