package handler

import (
	"fmt"
	"nead-desk/src/dto"
	"nead-desk/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CategoriesHandler struct {
	svc *service.CategoriesService
}

func NewCategoriesHandler(service *service.CategoriesService) *CategoriesHandler {
	return &CategoriesHandler{
		svc: service,
	}
}

func (h *CategoriesHandler) HandlerCreateCategories(c *gin.Context) {
	validate := validator.New()
	var req dto.CreateCategorieDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "corpo da solicitação inválido",
		})
		return
	}

	//Validação de campos
	if err := validate.Struct(req); err != nil {
		var errs []map[string]string
		for _, e := range err.(validator.ValidationErrors) {
			errs = append(errs, map[string]string{
				"field":   e.Field(),
				"message": fmt.Sprintf("failed on '%s' validation", e.Tag()),
			})
		}
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	// verificar duplicidade
	if exist := h.svc.ExisteCategoriName(req.Name); exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "O nome da categoria já existe!",
		})
		return
	}

	// Persistencia
	categorie, err := h.svc.CreateCategorie(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error ao salvar Categoria !",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":          categorie.ID,
		"name":        categorie.Name,
		"description": categorie.Description,
		"is_active":   categorie.Is_active,
		"created_at":  categorie.Created_at,
	})
}

func (h *CategoriesHandler) HandlerListAllCategores(c *gin.Context) {
	categories := h.svc.FindAllCategories()

	var data []map[string]string

	for _, v := range categories {
		data = append(data, map[string]string{
			"id":          v.ID,
			"name":        v.Name,
			"description": v.Description,
			"is_active":   strconv.FormatBool(v.Is_active),
		})

	}

	c.JSON(http.StatusOK, data)
}

func (h *CategoriesHandler) HandlerUpdateCategores(c *gin.Context) {
	categoriId := c.Param("id")
	validate := validator.New()
	var req dto.UpdateCategorieDto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Corpo da requisição inválida!"})
		return
	}
	// Validação
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := h.svc.UpdateCategories(categoriId, &req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error na atualização da entidade",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          result.ID,
		"name":        result.Name,
		"description": result.Description,
		"is_active":   result.Is_active,
		"updated_at":  result.Updated_at,
	})
}
