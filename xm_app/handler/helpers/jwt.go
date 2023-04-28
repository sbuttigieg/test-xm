package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	app "github.com/sbuttigieg/test-xm/xm_app"
	"github.com/sbuttigieg/test-xm/xm_app/errors"
	"github.com/sbuttigieg/test-xm/xm_app/models"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func GenerateJWT(config *app.Config, email string, username string) (string, error) {
	claims := &models.JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.JWTExpiry).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSigned, err := token.SignedString(jwtKey)
	if err != nil {
		config.Log.Errorf("helpers, method: GenerateJWT, layer: http, error: %v", err)

		return "", errors.Internal("couldn't sign token")
	}

	return tokenSigned, nil
}

func ValidateToken(config *app.Config, signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		config.Log.Errorf("helpers, method: ValidateToken - parse claims, layer: http, error: %v", err)

		return errors.Internal("couldn't parse claims")
	}

	claims, ok := token.Claims.(*models.JWTClaim)
	if !ok {
		config.Log.Errorf("helpers, method: ValidateToken - check claims, layer: http, error: %v", err)

		return errors.Internal("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		config.Log.Errorf("helpers, method: ValidateToken - Expired, layer: http, error: %v", err)

		return errors.Internal("token expired")
	}

	return nil
}
