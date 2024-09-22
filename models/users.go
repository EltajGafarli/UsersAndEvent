package models

import (
	"BookingApp/db"
	"BookingApp/utils"
	"database/sql"
	"errors"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (user User) Save() error {
	insertQuery := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(insertQuery)

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	user.ID = int(userId)
	return err
}

func (user *User) ValidateUser() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		return errors.New("invalid email or password")
	}

	hash := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !hash {
		return errors.New("invalid email or password")
	}

	return nil
}
