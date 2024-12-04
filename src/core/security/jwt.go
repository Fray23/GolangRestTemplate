package security

import (
	"fmt"

	db_models "code/models"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

func GenerateJwtToken(user *db_models.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"login":   user.Login,
	})
	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeJwtToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
