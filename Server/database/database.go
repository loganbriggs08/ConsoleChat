package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTables() bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return false
	} else {
		tablesToCreate := `
		CREATE TABLE IF NOT EXISTS active_heartbeats(authorization VARCHAR(255), last_heartbeat VARCHAR(255));
		`

		_, err := database.Exec(tablesToCreate)

		if err != nil {
			return false
		} else {
			return true
		}
	}
}

func AddHeartBeat(authorization string, snowflake uint64) bool {
	database, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		return false
	} else {
		_, err := database.Exec("INSERT INTO active_heartbeats VALUES(?, ?)", authorization, time.Now())

		if err != nil {
			return false
		} else {
			return true
		}
	}
}
