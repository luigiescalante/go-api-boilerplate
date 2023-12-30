package models

import (
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v5"
	"go.api-boilerplate/config"
	"strconv"
	"time"
)

type Users struct {
	Repo      UsersRepo
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	AuthType  string `json:"auth_type" validate:"required"`
}

func (usr *Users) Validate() error {
	var validate = validator.New()
	err := validate.Struct(usr)
	if err != nil {
		return err
	}
	return nil
}

func (usr *Users) ValidateLogin() error {
	var validate = validator.New()
	err := validate.Var(&usr.Email, "required,email")
	if err != nil {
		return err
	}
	err = validate.Var(&usr.Password, "required")
	if err != nil {
		return err
	}
	return nil
}

func (usr *Users) GetAuthToken(userId uint64) (string, error) {
	cfg := config.GetConfig()
	expiredAt, _ := strconv.ParseInt(cfg.GetExpiredTokenMinutes(), 10, 32)
	claims := jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Duration(int64(time.Minute) * expiredAt)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.GetJwtKey()))

}
