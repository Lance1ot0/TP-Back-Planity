package middleware

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(id int, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["user_role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	secretKeyStr := "R73pY17oMjuVSuhi47okiB9BAzDkYFUb"
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

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Authorization header is required")
			return
		}

		splitToken := strings.Split(tokenHeader, "Bearer ")
		if len(splitToken) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Malformed token")
			return
		}

		tokenString := splitToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("R73pY17oMjuVSuhi47okiB9BAzDkYFUb"), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid token: %v", err)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid token")
			return
		}

		context := context.WithValue(r.Context(), "user", token.Claims)
		next.ServeHTTP(w, r.WithContext(context))
	})
}

func DecodeJWT(tokenString string) (jwt.MapClaims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("R73pY17oMjuVSuhi47okiB9BAzDkYFUb"), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
