package handler

import (
	"errors"
	"nead-desk/src/domain"
	"nead-desk/src/dto"
	"nead-desk/src/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) UserAuth(c *gin.Context) {

	var req dto.UserAuth

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "corpo da solicitação inválido",
		})
		return
	}

	//Buscar usuario
	user, err := h.service.FindUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidUser) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Usuario Inválido",
			})
			return
		}
	}

	//Comparar hash da senha
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password_hash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email ou Senha incorreto",
		})
		return
	}

	//Gerar o token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"cargo":   user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // expira em 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Lembrar de trocar a secret ^,^
	assinaturaToken, err := token.SignedString([]byte("nead_desk_secret"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error ao Gerar token"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"access_token": assinaturaToken, "token_type": "Bearer"})
}
