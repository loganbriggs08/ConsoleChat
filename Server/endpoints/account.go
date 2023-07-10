package endpoints

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type accountCreated struct {
	SnowFlake     uint64 `json:"snowflake"`
	Authorization string `json:"authorization"`
}

func AccountCreate(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	min := int64(10000000000000000)
	max := int64(99999999999999999)

	SnowFlake := rand.Int63n(max-min+1) + min

	newUser := accountCreated{
		SnowFlake:     uint64(SnowFlake),
		Authorization: "asdjakjdkjas",
	}

	w.WriteHeader(http.StatusOK)

	newUserJson, _ := json.Marshal(newUser)
	w.Write(newUserJson)

}
