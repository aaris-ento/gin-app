package utils

import (
	"errors"
	"gin-app/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var jwtKey = []byte(GetEnv("jwt-secret", "secret"))

func GenerateJWT(user *models.User) string {
	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"user_role": user.Role,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	val, err := token.SignedString(jwtKey)
	if err != nil {

	}
	return val
}

func ValidateToken(tokenStr string) (uint, string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return 0, "", errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil || claims["user_role"] == nil {
		return 0, "", errors.New("invalid claims in token")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, "", errors.New("user_id claim is not valid")
	}
	userRoleString, ok := claims["user_role"].(string)
	if !ok {
		return 0, "", errors.New("user_role claim is not valid")
	}

	return uint(userIDFloat), userRoleString, nil
}
