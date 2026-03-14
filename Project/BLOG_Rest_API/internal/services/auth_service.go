package services

import (
	"blog_rest_api/internal/config"
	"blog_rest_api/internal/middlewares"
	"blog_rest_api/pkg/utils"
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int `json:"uid"`
	jwt.RegisteredClaims
}

func UserAuthService(ctx context.Context, r *http.Request) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", utils.ErrorHandler(err, "Unable to load the config")
	}

	cookie, err := r.Cookie("Bearer")
	if err != nil {
		return "", utils.ErrorHandler(err, "Unable to get the JWT token")
	}

	claims := &Claims{}

	parsedToken, err := jwt.ParseWithClaims(cookie.Value, claims, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing algorithm")
		}
		return []byte(cfg.JWT_SECRET), nil
	},
	)

	if err != nil {
		return "", utils.ErrorHandler(err, "Unable to parse the token")
	}

	if !parsedToken.Valid {
		return "", utils.ErrorHandler(errors.New("invalid token"), "Invalid token")
	}

	authID := strconv.Itoa(claims.UserID)

	return authID, nil

}

func AuthUser(r *http.Request, field string) (interface{}, error) {
	switch field {
	case "id":
		id, ok := r.Context().Value(middlewares.UserIDKey).(float64)
		if !ok {
			return 0, utils.ErrorHandler(errors.New("unable to get the ID from JWT"), "Unable to get the ID from JWT")
		} else {
			return int(id), nil
		}
	case "username":
		username, ok := r.Context().Value(middlewares.UsernameKey).(string)
		if !ok {
			return 0, utils.ErrorHandler(errors.New("unable to get the username from JWT"), "Unable to get the ID from JWT")
		} else {
			return username, nil
		}
	case "role":
		role, ok := r.Context().Value(middlewares.RoleKey).(string)
		if !ok {
			return 0, utils.ErrorHandler(errors.New("unable to get the role from JWT"), "Unable to get the ID from JWT")
		} else {
			return role, nil
		}
	}
	return nil, utils.ErrorHandler(errors.New("invalid JWT field request"), "Invalid JWT field request")
}
