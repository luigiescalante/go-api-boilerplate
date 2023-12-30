package services

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.api-boilerplate/models"
	"go.api-boilerplate/services/mock"
	"go.api-boilerplate/utils"
	"testing"
)

func TestUserSignUp(t *testing.T) {
	usersRepoMock := mock.UsersRepoMock{}
	userExpected := &models.Users{
		Repo:      &usersRepoMock,
		FirstName: "john",
		LastName:  "down",
		Email:     "test@test.com",
		Password:  "pass123",
		AuthType:  "test",
	}
	usersRepoMock.Mock.On("Save", userExpected).Return(userExpected, nil)

	user := &models.Users{
		Repo:      &usersRepoMock,
		FirstName: "john",
		LastName:  "down",
		Email:     "test@test.com",
		Password:  "pass123",
		AuthType:  "test",
	}
	auth := utils.PasswordAuthDummy{}
	signUp, err := UserSignUp(user, auth)
	assert.NotNil(t, signUp)
	assert.Equal(t, "test-token", signUp)
	assert.NoError(t, err)
}

func TestUserSignUpError(t *testing.T) {
	usersRepoMock := mock.UsersRepoMock{}
	userExpected := &models.Users{
		Repo:      &usersRepoMock,
		FirstName: "john",
		LastName:  "down",
		Email:     "test@test.com",
		Password:  "pass123",
		AuthType:  "test",
	}
	usersRepoMock.Mock.On("Save", userExpected).Return(nil, fmt.Errorf("error on database"))

	user := &models.Users{
		Repo:      &usersRepoMock,
		FirstName: "john",
		LastName:  "down",
		Email:     "test@test.com",
		Password:  "pass123",
		AuthType:  "test",
	}
	signUp, err := UserSignUp(user, utils.PasswordAuthDummy{})
	assert.Empty(t, signUp)
	assert.Error(t, err)
}

func TestUserLogin(t *testing.T) {
	userRepoMock := mock.UsersRepoMock{}
	userExpected := &models.Users{
		Repo:      &userRepoMock,
		ID:        1,
		FirstName: "john",
		LastName:  "down",
		Email:     "test@test.com",
		Password:  "pass123",
		AuthType:  "test",
	}
	userRepoMock.Mock.On("GetByEmail", "test@test.com").Return(userExpected, nil)

	user := &models.Users{
		Repo:      &userRepoMock,
		FirstName: "john",
		LastName:  "down",
		Email:     "test@test.com",
		Password:  "pass123",
		AuthType:  "test",
	}
	usrLogin, err := Login(user, utils.PasswordAuthDummy{})
	assert.NotNil(t, usrLogin)
	assert.NoError(t, err)
}

func TestUserLoginError(t *testing.T) {
	userRepoMock := mock.UsersRepoMock{}
	userRepoMock.Mock.On("GetByEmail", "test@test.com").Return(nil, fmt.Errorf("error db"))

	user := &models.Users{
		Repo:      &userRepoMock,
		FirstName: "john",
		LastName:  "down",
		Email:     "test@test.com",
		Password:  "pass123",
		AuthType:  "test",
	}
	usrLogin, err := Login(user, utils.PasswordAuthDummy{})
	assert.Empty(t, usrLogin)
	assert.Error(t, err)
}

func TestUserLoginUserNotExist(t *testing.T) {
	userRepoMock := mock.UsersRepoMock{}
	userRepoMock.Mock.On("GetByEmail", "test@test.com").Return(nil, nil)

	user := &models.Users{
		Repo:      &userRepoMock,
		FirstName: "john",
		LastName:  "down",
		Email:     "test@test.com",
		Password:  "pass123",
		AuthType:  "test",
	}
	usrLogin, err := Login(user, utils.PasswordAuthDummy{})
	assert.Empty(t, usrLogin)
	assert.Error(t, err)
}
