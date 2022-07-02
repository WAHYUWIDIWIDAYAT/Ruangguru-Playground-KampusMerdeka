package repository

import (
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByID(id int64) (User, error) {
	var user User
	var sqlStatement string

	sqlStatement = "SELECT * FROM users WHERE id = ?"
	row := u.db.QueryRow(sqlStatement, id)

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	var users []User
	var sqlStatement string

	sqlStatement = "SELECT * FROM users"
	rows, err := u.db.Query(sqlStatement)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *UserRepository) Login(username string, password string) (*string, error) {

	var user User
	var sqlStatement string

	sqlStatement = "SELECT * FROM users WHERE username = ?"
	row := u.db.QueryRow(sqlStatement, username)

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)
	if err != nil {
		return nil, err
	}

	if user.Password == password {
		return &user.Role, nil
	}

	return nil, errors.New("username or password is incorrect")
}

func (u *UserRepository) InsertUser(username string, password string, role string, loggedin bool) error {
	var sqlStatement string

	sqlStatement = "INSERT INTO users (username, password, role, loggedin) VALUES (?, ?, ?, ?)"
	_, err := u.db.Exec(sqlStatement, username, password, role, loggedin)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) FetchUserRole(username string) (*string, error) {
	var user User
	var sqlStatement string

	sqlStatement = "SELECT * FROM users WHERE username = ?"
	row := u.db.QueryRow(sqlStatement, username)

	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.Loggedin)
	if err != nil {
		return nil, err
	}

	return &user.Role, nil
}
