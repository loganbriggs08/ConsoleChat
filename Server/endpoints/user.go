package endpoints

import (
	"encoding/json"
	"net/http"
)

type accountCreated struct {
	SnowFlake     uint64 `json:"snowflake"`
	Authorization string `json:"authorization"`
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	newUser := accountCreated{
		SnowFlake:     12345677889343,
		Authorization: "asdjakjdkjas",
	}

	w.WriteHeader(http.StatusOK)

	newUserJson, _ := json.Marshal(newUser)
	w.Write(newUserJson)
}
