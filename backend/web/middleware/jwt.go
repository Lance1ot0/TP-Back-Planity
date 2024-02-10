package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(id int, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["user_role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	secretKeyStr, err := generateRandomKey(32)
	if err != nil {
		fmt.Println("Erreur lors de la génération de la clé secrète:", err)
		return "", err
	}

	secretKey := []byte(secretKeyStr)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateRandomKey(length int) (string, error) {
	keyBytes := make([]byte, length)

	_, err := rand.Read(keyBytes)
	if err != nil {
		return "", err
	}

	key := base64.StdEncoding.EncodeToString(keyBytes)
	return key, nil
}
