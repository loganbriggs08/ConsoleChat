package main

import (
	"net/http"

	"github.com/NotKatsu/ConsoleChat/database"
	"github.com/NotKatsu/ConsoleChat/endpoints"
	"github.com/pterm/pterm"
)

func main() {
	result := database.CreateTables()

	if result == true {
		// http.HandleFunc("/api/heartbeat/send", endpoints.HeartbeatSend)
		http.HandleFunc("/api/account/new", endpoints.AccountCreate)

		err := http.ListenAndServe(":8080", nil)

		if err != nil {
			pterm.Fatal.WithFatal(true).Println(err)
		}
	} else {
		pterm.Fatal.WithFatal(true).Println("An error occured while creating the database tables.")
	}
}
