package handler

import (
	"errors"
	"nead-desk/src/domain"
	"nead-desk/src/dto"
	"nead-desk/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CalledHandler struct {
	service *service.CalledService
}

func NewCalledHandler(service *service.CalledService) *CalledHandler {
	return &CalledHandler{
		service: service,
	}
}

func (h *CalledHandler) Create(c *gin.Context) {
	var req dto.CreateCalledRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "corpo da solicitação inválido",
		})
		return
	}

	called, err := h.service.CreateCalled(&req)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCalled) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if errors.Is(err, domain.ErrCalledNotFound) {
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

	c.JSON(http.StatusCreated, called)
}

func (h *CalledHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	called, err := h.service.GetCalledByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrCalledNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Chamado não encontrado",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro do Servidor Interno",
		})
		return
	}

	c.JSON(http.StatusOK, called)
}

func (h *CalledHandler) GetAll(c *gin.Context) {
	called, err := h.service.GetAllCalled()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro do Servidor Interno",
		})
		return
	}

	c.JSON(http.StatusOK, called)
}

func (h *CalledHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateCalledRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "corpo da solicitação inválido",
		})
		return
	}

	called, err := h.service.UpdateCalled(id, &req)
	if err != nil {
		if errors.Is(err, domain.ErrCalledNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Chamado não encontrada",
			})
			return
		}

		if errors.Is(err, domain.ErrInvalidCalled) {
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

	c.JSON(http.StatusOK, called)
}

func (h *CalledHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteCalled(id)
	if err != nil {
		if errors.Is(err, domain.ErrCalledNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Chamado não encontrado",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro do Servidor Interno",
		})
		return
	}

	c.Status(http.StatusNoContent)
}
