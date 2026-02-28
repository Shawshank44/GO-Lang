package middlewares

import (
	"blog_rest_api/internal/config"
	"blog_rest_api/pkg/utils"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const (
	UserIDKey   ContextKey = "userId"
	RoleKey     ContextKey = "role"
	UsernameKey ContextKey = "username"
	Expiry      ContextKey = "expiresAt"
)

func AuthorizeUser(userRole string, allowedRoles ...string) (bool, error) {
	for _, allowedRole := range allowedRoles {
		if userRole == allowedRole {
			return true, nil
		}
	}
	return false, errors.New("unauathorized you are not privilaged to access the page")
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg, err := config.Load()
		if err != nil {
			http.Error(w, "unable to configure", http.StatusInternalServerError)
			return
		}

		token, err := r.Cookie("Bearer")
		if err != nil {
			http.Error(w, "Authorization Header Missing", http.StatusUnauthorized)
			return
		}

		jwtSecret := cfg.JWT_SECRET

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
		ctx = context.WithValue(ctx, ContextKey(UserIDKey), claims["uid"])

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
