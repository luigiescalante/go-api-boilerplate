package mock

import (
	"github.com/stretchr/testify/mock"
	"go.api-boilerplate/models"
)

type UsersRepoMock struct {
	mock.Mock
}

func (m *UsersRepoMock) Save(user *models.Users) (*models.Users, error) {
	args := m.Mock.Called(user)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Users), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *UsersRepoMock) GetByEmail(email string) (*models.Users, error) {
	args := m.Mock.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Users), args.Error(1)
	}
	return nil, args.Error(1)
}
