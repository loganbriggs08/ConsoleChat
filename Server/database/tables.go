package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Create() bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return false
	} else {
		tablesToCreate := `
		CREATE TABLE IF NOT EXISTS active_heartbeats(authorization VARCHAR(255), last_heartbeat VARCHAR(255));
		CREATE TABLE IF NOT EXISTS accounts(snowflake BIGINT UNSIGNED, authorization VARCHAR(255))
		`

		_, err := database.Exec(tablesToCreate)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}
