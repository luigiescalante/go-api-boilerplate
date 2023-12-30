package services

import (
	"fmt"
	"go.api-boilerplate/models"
	"go.api-boilerplate/utils"
)

func UserSignUp(user *models.Users, auth utils.PasswordAuthMethod) (string, error) {
	var err error
	user.Password, err = auth.SetPassword(user.Password)
	if err != nil {
		return "", err
	}
	_, err = user.Repo.Save(user)
	if err != nil {
		return "", err
	}
	token, err := auth.GetJwtToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func Login(user *models.Users, auth utils.PasswordAuthMethod) (string, error) {
	usrExist, err := user.Repo.GetByEmail(user.Email)
	if err != nil {
		return "", err
	}
	if usrExist == nil {
		return "", fmt.Errorf("user or password invalid")
	}

	hash, err := auth.CheckPasswordHash(user.Password, usrExist.Password)
	if err != nil {
		return "", fmt.Errorf("user or password invalid")
	}
	if !hash {
		return "", fmt.Errorf("user or password invalid")
	}
	token, err := auth.GetJwtToken(usrExist.ID)
	if err != nil {
		return "", fmt.Errorf("user or password invalid")
	}
	return token, nil
}
