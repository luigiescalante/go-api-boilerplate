package db

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.api-boilerplate/models"
	"testing"
)

func TestUserDb_Save(t *testing.T) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(1)
	mock.ExpectQuery(signUp).WithArgs("john", "down", "test@test.com", "test123").WillReturnRows(rows)

	user := &models.Users{
		FirstName: "john",
		LastName:  "down",
		Email:     "test@test.com",
		Password:  "test123",
	}
	userRepo := NewUsersRepo(mockDB)
	userDb, err := userRepo.Save(user)
	assert.NotNil(t, userDb)
	assert.NoError(t, err)
}

func TestUserDb_SaveError(t *testing.T) {
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	mock.ExpectQuery(signUp).WithArgs("john", "down", "test@test.com", "test123").WillReturnError(fmt.Errorf("error on database"))
	user := &models.Users{
		FirstName: "john",
		LastName:  "down",
		Email:     "test@test.com",
		Password:  "test123",
	}
	userRepo := NewUsersRepo(mockDB)
	userDb, err := userRepo.Save(user)
	assert.Nil(t, userDb)
	assert.Error(t, err)
}

func TestUsersDb_GetByEmail(t *testing.T) {
	email := "test@test.com"
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password"})
	rows.AddRow(1, "john", "down", "test@test.com", "test123")
	mock.ExpectQuery(getByEmail).WithArgs(email).WillReturnRows(rows)
	userRepo := NewUsersRepo(mockDB)
	userDb, err := userRepo.GetByEmail(email)
	assert.NotNil(t, userDb)
	assert.NoError(t, err)
}

func TestUsersDb_GetByEmailNotFound(t *testing.T) {
	email := "test@test.com"
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email", "password"})
	mock.ExpectQuery(getByEmail).WithArgs(email).WillReturnRows(rows)
	userRepo := NewUsersRepo(mockDB)
	userDb, err := userRepo.GetByEmail(email)
	assert.Nil(t, userDb)
	assert.NoError(t, err)
}

func TestUsersDb_GetByEmailError(t *testing.T) {
	email := "test@test.com"
	mockDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	mock.ExpectQuery(getByEmail).WithArgs(email).WillReturnError(fmt.Errorf("error database"))
	userRepo := NewUsersRepo(mockDB)
	userDb, err := userRepo.GetByEmail(email)
	assert.Nil(t, userDb)
	assert.Error(t, err)
}
