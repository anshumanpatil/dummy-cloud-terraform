package jwthelper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT() (tokenString string, err error) {

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte("mysecretkey"))

	return
}

func ValidateJWT(authorizationTokenString string) (*jwt.Token, error) {
	return jwt.Parse(authorizationTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mysecretkey"), nil
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if len(ctx.Request.Header["Authorization"]) <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Authorization required!"})
			ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("Authorization required!"))
			return
		}
		authorizationTokenString := ctx.Request.Header["Authorization"][0]

		_, err := ValidateJWT(authorizationTokenString)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		ctx.Next()
	}
}
