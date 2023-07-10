package main

import (
	"net/http"

	"github.com/NotKatsu/ConsoleChat/endpoints"
	"github.com/pterm/pterm"
)

func main() {
	http.HandleFunc("/api/heartbeat/send", endpoints.HeartbeatSend)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}
}
