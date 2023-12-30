package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"go.api-boilerplate/config"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type PasswordAuthMethod interface {
	SetPassword(passwordPlainText string) (string, error)
	CheckPasswordHash(passwordPlainText string, hash string) (bool, error)
	GetExpiredTokenMinutes() int
	GetJwtToken(userId uint64) (string, error)
}

type PasswordAuth struct {
}

func (pass PasswordAuth) SetPassword(passwordPlainText string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwordPlainText), 14)
	return string(bytes), err
}

func (pass PasswordAuth) GetExpiredTokenMinutes() int {
	cfg := config.GetConfig()
	expired, _ := strconv.Atoi(cfg.GetExpiredTokenMinutes())
	return expired
}

func (pass PasswordAuth) GetJwtToken(userId uint64) (string, error) {
	cfg := config.GetConfig()
	expiredAt, _ := strconv.ParseInt(cfg.GetExpiredTokenMinutes(), 10, 32)
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Duration(int64(time.Minute) * expiredAt)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.GetJwtKey()))
}

func (pass PasswordAuth) CheckPasswordHash(passwordPlainText, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordPlainText))
	if err != nil {
		return false, err
	}
	return true, err
}

type PasswordAuthDummy struct {
}

func (dummy PasswordAuthDummy) GetExpiredTokenMinutes() int {
	return 30
}

func (dummy PasswordAuthDummy) GetJwtToken(userId uint64) (string, error) {
	return "test-token", nil
}

func (dummy PasswordAuthDummy) SetPassword(passwordPlainText string) (string, error) {
	return passwordPlainText, nil
}

func (dummy PasswordAuthDummy) CheckPasswordHash(passwordPlainText string, hash string) (bool, error) {
	return true, nil
}
