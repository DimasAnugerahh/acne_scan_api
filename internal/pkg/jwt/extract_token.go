package jwt

import (
	"acne-scan-api/internal/model/web"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractToken(tokenString string) (*web.AuthResponse, error) {

	type MyCustomClaims struct {
		Username string `json:"username"`
		Role     string `json:"role"`
		UserId   int `json:"user_id"`
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		extractedToken := &web.AuthResponse{
			Username: claims.Username,
			Role:     claims.Role,
			UserId:   claims.UserId,
		}
		return extractedToken, nil
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil, fmt.Errorf("invalid token")
}
