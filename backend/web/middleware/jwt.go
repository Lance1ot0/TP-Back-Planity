package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(id int) (string, error) {
	// Créez un token JWT avec une clé secrète
	token := jwt.New(jwt.SigningMethodHS256)

	// Créez les revendications (claims) que vous souhaitez inclure
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Durée de validité du token

	secretKeyStr, err := generateRandomKey(32)
	if err != nil {
		fmt.Println("Erreur lors de la génération de la clé secrète:", err)
		return "", err
	}

	// Signez le token avec une clé secrète
	secretKey := []byte(secretKeyStr)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateRandomKey(length int) (string, error) {
	// Créez un slice de bytes pour stocker les octets aléatoires
	keyBytes := make([]byte, length)

	// Lisez des octets aléatoires depuis /dev/urandom ou crypto/rand
	_, err := rand.Read(keyBytes)
	if err != nil {
		return "", err
	}

	// Encodez les octets aléatoires en base64 pour obtenir une clé lisible
	key := base64.StdEncoding.EncodeToString(keyBytes)
	// ---
	return key, nil
}
