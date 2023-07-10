package main

import (
	"encoding/json"
	"net/http"

	"github.com/pterm/pterm"
)

type apiError struct {
	ErrorCode uint16 `json:"error_code"`
}

func heartbeatSend(w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get("Authorization")

	if authorizationHeader == "" {
		err := apiError{
			ErrorCode: 65535,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		jsonError, _ := json.Marshal(err)
		_, writeError := w.Write(jsonError)

		if writeError != nil {
			pterm.Fatal.WithFatal(true).Println("Failed to write the data to the connection.")
		}
	} else {
		
	}
}

func main() {
	http.HandleFunc("/api/heartbeat/send", heartbeatSend)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		pterm.Fatal.WithFatal(true).Println(err)
	}
}
