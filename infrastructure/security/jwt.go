package security

import (
	"chat/domain/entities"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const secretKey = "mysecretkey"

func GenerateJWT(user entities.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":  user.Name,
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(), // Caducidad del token en 2 horas
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(fieldToGet, tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unespected sign method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", fmt.Errorf("error: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		valueExtracted := claims[fieldToGet].(string)
		return valueExtracted, nil
	}

	return "", fmt.Errorf("invalid token")
}
