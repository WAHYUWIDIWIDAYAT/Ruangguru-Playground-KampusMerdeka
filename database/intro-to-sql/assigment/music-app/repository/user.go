package repository

import (
	"database/sql"

	"github.com/ruang-guru/playground/backend/database/intro-to-sql/assigment/music-app/model"
)

type UserRepositoryInterface interface {
	FetchUserByID(id int) (*model.User, error)
	DeleteUserByID(id int) error
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (ur *UserRepository) FetchUserByID(id int) (*model.User, error) {
	var sqlStatement string

	sqlStatement = "SELECT * FROM users WHERE id = ?"

	var user model.User
	row := ur.db.QueryRow(sqlStatement, id)
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) DeleteUserByID(id int) error {
	var sqlStatement string


	sqlStatement = "DELETE FROM users WHERE id = ?"

	_, err := ur.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}
