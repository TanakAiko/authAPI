package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	Id         string
	UserID     int
	Expiration time.Time
}

func (session *Session) CreateSession(user User, db *sql.DB) error {
	session.Id = uuid.New().String()

	session.Expiration = time.Now().Add(24 * time.Hour)

	// This code snippet is starting a new transaction ,`tx`, on the database connection `db`.
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	defer tx.Rollback()

	content, err := os.ReadFile("./databases/sqlRequests/insertNewSession.sql")
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

	_, err = stmt.Exec(
		session.Id,
		user.Id,
		session.Expiration.Format(time.RFC3339),
	)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (session *Session) TestSessionUser(db *sql.DB) error {
	err := db.QueryRow("SELECT userID, expiresAT FROM sessions WHERE sessionID = ?", session.Id).Scan(&session.UserID, &session.Expiration)
	if err != nil {
		return err
	}

	if time.Now().After(session.Expiration) {
		session.deleteSession(db)
		return fmt.Errorf("session expired")
	}

	return nil
}

func (session *Session) deleteSession(db *sql.DB) {
	_, err := db.Exec("DELETE FROM sessions WHERE sessionID = ?", session.Id)
	if err != nil {
		log.Printf("Error deleting session: %v", err)
	}
}
