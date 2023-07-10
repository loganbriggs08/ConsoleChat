package main

import (
	"fmt"
	"net/http"

	"github.com/pterm/pterm"
)

type apiError struct {
	ErrorCode uint16 `json:"error"`
}

func heartbeatSend(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "EndPoint Reached")
}

func main() {
	http.HandleFunc("/api/heartbeat/send", heartbeatSend)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}
}
