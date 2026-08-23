package handler

import (
	"errors"
	"nead-desk/src/domain"
	"nead-desk/src/dto"
	"nead-desk/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req dto.CreateUserRequest

	// ShouldBindJSON deserializa o body e valida tags básicas se houver.
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "corpo da solicitação inválido",
		})
		return
	}

	user, err := h.service.CreateUser(&req)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidUser) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro do Servidor Interno",
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id") // :id da rota GET /todos/:id

	user, err := h.service.GetUserByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Usuario não encontrada",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro do Servidor Interno",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro do Servidor Interno",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "corpo da solicitação inválido",
		})
		return
	}

	user, err := h.service.UpdateUser(id, &req)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Usuario não encontrada",
			})
			return
		}

		if errors.Is(err, domain.ErrInvalidUser) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro do Servidor Interno",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteUser(id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Usuario não encontrada",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro do Servidor Interno",
		})
		return
	}

	// 204 No Content: sucesso sem body — padrão REST para DELETE.
	c.Status(http.StatusNoContent)
}
