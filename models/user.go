package models

import (
	"database/sql"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                int       `json:"userId"`
	Nickname          string    `json:"nickname"`
	Age               int       `json:"age"`
	Gender            string    `json:"gender"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Identifier        string    `json:"identifier"`
	SessionID         string    `json:"sessionID"`
	SessionExpireTime time.Time `json:"sessionExpireTime"`
	CreateAt          time.Time `json:"createAt"`
}

func (user *User) CreateUser(db *sql.DB) error {
	// This code snippet is starting a new transaction ,`tx`, on the database connection `db`.
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	content, err := os.ReadFile("./databases/sqlRequests/insertNewUser.sql")
	if err != nil {
		return err
	}

	// This code snippet is preparing a SQL statement for execution within a transaction.
	stmt, err := tx.Prepare(string(content))
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	// Crypte the password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		user.Nickname,
		user.Age,
		user.Gender,
		user.FirstName,
		user.LastName,
		user.Email,
		hashedPassword,
		time.Now().Format(time.RFC3339),
	)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
		return err
	}

	return err
}

func (user *User) GetUser(db *sql.DB) (string, error) {
	var hashedPassword string
	err := db.QueryRow("SELECT id, nickname, age, gender, firstName, lastName, email, password, createdAt FROM users WHERE nickname = ?", user.Identifier).Scan(
		&user.Id,
		&user.Nickname,
		&user.Age,
		&user.Gender,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&hashedPassword,
		&user.CreateAt,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			err2 := db.QueryRow("SELECT id, nickname, age, gender, firstName, lastName, email, password, createdAt FROM users WHERE email = ?", user.Identifier).Scan(
				&user.Id,
				&user.Nickname,
				&user.Age,
				&user.Gender,
				&user.FirstName,
				&user.LastName,
				&user.Email,
				&hashedPassword,
				&user.CreateAt,
			)
			if err2 != nil {
				return "", err2
			}
		} else {
			return "", err
		}

	}
	return hashedPassword, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
