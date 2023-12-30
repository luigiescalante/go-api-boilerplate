package models

type UsersRepo interface {
	GetByEmail(email string) (*Users, error)
	Save(user *Users) (*Users, error)
}
