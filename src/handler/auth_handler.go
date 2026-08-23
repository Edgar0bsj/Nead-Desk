package handler

import (
	"nead-desk/src/dto"
	"nead-desk/src/storage/ports"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepository ports.UserStorage
}

func NewAuthHandler(userRepo ports.UserStorage) *AuthHandler {
	return &AuthHandler{
		userRepository: userRepo,
	}
}

func (l *AuthHandler) Login(c *gin.Context) {
	var req dto.AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	// Buscar usuário no banco
	user, err := l.userRepository.FindByEmail(req.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email ou Senha inválido"})
		return
	}

	// Verificar se a senha bate com a do banco
	if err := bcrypt.CompareHashAndPassword([]byte(user.SenhaHash), []byte(req.Senha)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email ou Senha inválido"})
		return
	}

	// Gerar token JWT
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"cargo":   user.Cargo.String(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // expira em 24h
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//Lembrar de trocar a secret ^,^
	assinaturaToken, err := token.SignedString([]byte("nead_desk_secret"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error ao Gerar token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": assinaturaToken})

}
