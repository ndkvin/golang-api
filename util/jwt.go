package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

	"os"
)

type JwtClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(id uint, email string) (token string, err error) {
	godotenv.Load()
	secret := os.Getenv("JWT_SECRET")

	claims := &JwtClaims{
		id,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyJWT(token *jwt.Token) (claims *JwtClaims, err error) {

	claims = token.Claims.(*JwtClaims)

	return
}
