package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/pterm/pterm"
)

type apiError struct {
	ErrorCode uint16 `json:"error_code"`
}

func HeartbeatSend(w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get("Authorization")

	if authorizationHeader == "" {
		err := apiError{
			ErrorCode: http.StatusUnauthorized,
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
