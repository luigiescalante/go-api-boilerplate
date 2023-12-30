package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsers_Validate(t *testing.T) {
	usr := &Users{
		FirstName: "john",
		LastName:  "down",
		Email:     "test@test.com",
		Password:  "pass123",
		AuthType:  "test",
	}
	err := usr.Validate()
	assert.NoError(t, err)
}

func TestUsers_ValidateNotValid(t *testing.T) {
	usr := &Users{
		FirstName: "",
		LastName:  "",
		Email:     "test@test.com",
		Password:  "",
		AuthType:  "test",
	}
	err := usr.Validate()
	assert.Error(t, err)
}

func TestUsers_ValidateLogin(t *testing.T) {
	usr := &Users{
		Email:    "test@test.com",
		Password: "pass123",
	}
	err := usr.ValidateLogin()
	assert.NoError(t, err)
}

func TestUsers_ValidateLoginNotValid(t *testing.T) {
	usr := &Users{
		Email:    "test",
		Password: "pass123",
	}
	err := usr.ValidateLogin()
	assert.Error(t, err)
}

func TestUsers_GetAuthToken(t *testing.T) {
	usr := &Users{}
	token, err := usr.GetAuthToken(1)
	assert.NotNil(t, token)
	assert.NoError(t, err)
}
