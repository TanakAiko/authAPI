package dbmanager

import (
	"auth/internals/tools"
	md "auth/models"
	"database/sql"
	"log"
	"os"
	"time"
)

func CreateUser(user md.User, db *sql.DB) error {
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
	hashedPassword, err := tools.HashPassword(user.Password)
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
