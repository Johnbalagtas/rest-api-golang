package models

import (
	"errors"

	"golang.com/rest/db"
	"golang.com/rest/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.password)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.Id = userId
	return err

}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)

	if err != nil {
		return errors.New("credentials Invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials Invalid")
	}

	return nil

}

func (u *User) AllUsersList() ([]User, error) {
	query := `SELECT id, email, password FROM users`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Email, &user.password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
