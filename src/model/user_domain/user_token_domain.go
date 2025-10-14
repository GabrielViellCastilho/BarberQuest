package user_domain

import (
	"fmt"
	"github.com/GabrielViellCastilho/BarberQuest/src/configuration/rest_err"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

var JWT_SECRET_KEY = "JWT_SECRET_KEY"

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.id,
		"name":  ud.name,
		"role":  ud.role,
		"email": ud.email,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(
			fmt.Sprintf("Error while generating token: %v", err))
	}
	return tokenString, nil
}

func GeneratePasswordResetToken(email string) (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
		"iat":     time.Now().Unix(),
		"jti":     "unique-token-id",
		"purpose": "password_reset",
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(
			fmt.Sprintf("Error while generating token: %v", err))
	}
	return tokenString, nil
}

func VerifyToken(tokenvalue string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenvalue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("Invalid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedRequestError("Invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &userDomain{
			id:    int(claims["id"].(float64)),
			name:  claims["name"].(string),
			role:  claims["role"].(string),
			email: claims["email"].(string),
		}, nil
	}
	return nil, rest_err.NewUnauthorizedRequestError("Invalid token")
}

func ValidatePasswordResetToken(tokenvalue string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	token, err := jwt.Parse(RemoveBearerPrefix(tokenvalue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("Invalid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedRequestError("Invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		purpose, ok := claims["purpose"].(string)
		if !ok || purpose != "password_reset" {
			return nil, rest_err.NewUnauthorizedRequestError("Invalid token")
		}

		return &userDomain{
			email: claims["email"].(string),
		}, nil
	}
	return nil, rest_err.NewUnauthorizedRequestError("Invalid token")
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ")
	}
	return token
}
