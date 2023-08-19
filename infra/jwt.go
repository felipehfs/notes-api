package infra

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JwtSecret = "NOT_SAFE_TO_WORK_2030203"
)

func CreateToken(id string, email string, duration time.Duration) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(duration).Unix(),
		"id":     id,
		"email":  email,
		"iat":    time.Now(),
		"issuer": "test",
	})
	tokenSigned, err := token.SignedString([]byte(JwtSecret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenSigned
}

func JwtAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		headerPieces := strings.Split(authHeader, "Bearer ")
		log.Printf("Login pieces %v", headerPieces)
		if len(headerPieces) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string][]string{
				"errors": {"header malformed"},
			})
			return
		}
		token, err := jwt.Parse(headerPieces[1], func(t *jwt.Token) (interface{}, error) {

			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(JwtSecret), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string][]string{"errors": {err.Error()}})
		}

		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Not authorized")
		}

	})
}
