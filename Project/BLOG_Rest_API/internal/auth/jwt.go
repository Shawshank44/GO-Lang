package auth

import (
	"blog_rest_api/internal/config"
	"blog_rest_api/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignToken(userID int, Username, Role string) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", utils.ErrorHandler(err, "Unable to config")
	}

	jwtSecret := cfg.JWT_SECRET
	jwtExpiresIn := cfg.JWT_EXPIRES_IN

	claims := jwt.MapClaims{
		"uid":  userID,
		"user": Username,
		"role": Role,
	}

	if jwtExpiresIn != "" {
		duration, err := time.ParseDuration(jwtExpiresIn)
		if err != nil {
			return "", utils.ErrorHandler(err, "JWT internal error")
		}
		claims["exp"] = jwt.NewNumericDate(time.Now().Add(duration))
	} else {
		claims["exp"] = jwt.NewNumericDate(time.Now().Add(15 * time.Minute))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", utils.ErrorHandler(err, "JWT internal error")
	}

	return signedToken, nil
}
