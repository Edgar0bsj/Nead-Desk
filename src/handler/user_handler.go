package handler

import (
	"errors"
	"fmt"
	"nead-desk/src/domain"
	"nead-desk/src/dto"
	"nead-desk/src/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

// ======================================
// Login
// ======================================
func (h *UserHandler) UserAuthLogin(c *gin.Context) {

	var req dto.UserAuthLoginDto

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
		"user_id":   user.ID,
		"user_role": user.Role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // expira em 24h
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

// ======================================
// Register
// ======================================
func (h *UserHandler) UserAuthRegister(c *gin.Context) {
	validate := validator.New()
	var req dto.UserAuthRegisterDto

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

	//Verificar se email já existe.
	if exist, _ := h.service.FindUserByEmail(req.Email); exist != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Email já esta sendo utilizado",
		})
		return
	}

	//Gerar hash da senha
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error na Senha",
		})
		return
	}
	req.Password = string(passwordHash)

	// Criar usuário
	userEntity, err := h.service.CreateUser(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "falhas ao salvar",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":         userEntity.ID,
		"name":       userEntity.Name,
		"email":      userEntity.Email,
		"role":       userEntity.Role,
		"created_at": userEntity.Created_at,
	})
}
