package middlewares

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"order_mgt/pkg/utils"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const (
	UserIDkey   ContextKey = "userId"
	RoleKey     ContextKey = "role"
	UsernameKey ContextKey = "username"
	Expiry      ContextKey = "expiresAt"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("Bearer")
		if err != nil {
			http.Error(w, "Authorization Header Missing", http.StatusUnauthorized)
			return
		}

		jwtSecret := os.Getenv("JWT_SECRET")

		parsedToken, err := jwt.Parse(token.Value, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method : %v", t.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				http.Error(w, "Token Expired", http.StatusUnauthorized)
				return
			} else if errors.Is(err, jwt.ErrTokenMalformed) {
				http.Error(w, "Token Malformed", http.StatusUnauthorized)
				return
			}
			utils.ErrorHandler(err, "")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if parsedToken.Valid {
			log.Println("Valid JWT")
		} else {
			http.Error(w, "Invalid login token", http.StatusUnauthorized)
			log.Println("Invalid JWT : ", token.Value)
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid Login Token", http.StatusUnauthorized)
			log.Println("Invalid login token", token.Value)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKey(RoleKey), claims["role"])
		ctx = context.WithValue(ctx, ContextKey(Expiry), claims["exp"])
		ctx = context.WithValue(ctx, ContextKey(UsernameKey), claims["user"])
		ctx = context.WithValue(ctx, ContextKey(UserIDkey), claims["uid"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
