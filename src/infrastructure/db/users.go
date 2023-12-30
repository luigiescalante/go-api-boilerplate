package db

import (
	"database/sql"
	"go.api-boilerplate/models"
)

const (
	signUp = `INSERT INTO users (first_name, last_name, email, password, created_at) 
				VALUES ($1::varchar,$2::varchar,$3::varchar,$4::varchar,now()) RETURNING id`
	getByEmail = `SELECT id,first_name,last_name,email,password 
					FROM users 
					WHERE  email =$1 limit 1`
)

type UsersDb struct {
	Repo *sql.DB
}

func (usrDb *UsersDb) GetByEmail(email string) (*models.Users, error) {
	row, err := usrDb.Repo.Query(getByEmail, email)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	user := &models.Users{}
	for row.Next() {
		err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == 0 {
		return nil, nil
	}
	return user, nil
}
func (usrDb *UsersDb) Save(user *models.Users) (*models.Users, error) {
	row := usrDb.Repo.QueryRow(signUp, user.FirstName, user.LastName, user.Email, user.Password)
	if row.Err() != nil {
		return nil, row.Err()
	}
	err := row.Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUsersRepo(repo *sql.DB) *UsersDb {
	return &UsersDb{
		Repo: repo,
	}
}
