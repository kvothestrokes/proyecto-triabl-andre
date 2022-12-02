package controller

import (
	"PROYECTO/dto"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type JwtController interface {
	GetToken(ctx *gin.Context)
	ValidateToken(tokenString string) bool
}

type jwtController struct{}

func NewJwtController() JwtController {
	return &jwtController{}
}

func getSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

// GetToken genera un token para el usuario
func (i *jwtController) GetToken(ctx *gin.Context) {

	//crear el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authenticated": true,
		"expire":        time.Now().Add(time.Hour * 8).Unix(),
	})

	//firmar el token
	tokenString, err := token.SignedString(getSecretKey())
	if err != nil {
		handleErroResponse(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	tokenResponse := &dto.TokenHanlder{
		Token: tokenString,
	}

	//Retornar la respuesta
	response := dto.SearchResponseDTO{
		Status: true,
		Error:  "",
		Data:   tokenResponse,
	}
	ctx.JSON(http.StatusOK, response)
}

// Validar el token
func (i *jwtController) ValidateToken(tokenString string) bool {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error al leer el token  %v", token)
		}
		return getSecretKey(), nil
	})

	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	}

	return false
}
