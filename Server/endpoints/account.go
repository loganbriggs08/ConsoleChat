package endpoints

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/NotKatsu/ConsoleChat/authentication"
	"github.com/NotKatsu/ConsoleChat/database"
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
	Authorization := authentication.Encode(strconv.Itoa(int(SnowFlake))) + "." + authentication.Encode(strconv.Itoa(int(authentication.Since_Epoch()))) + "." + authentication.Encode(authentication.String(25))

	databaseResponse := database.AddHeartBeat(Authorization)

	if databaseResponse == true {
		newUser := accountCreated{
			SnowFlake:     uint64(SnowFlake),
			Authorization: Authorization,
		}

		w.WriteHeader(http.StatusOK)

		newUserJson, _ := json.Marshal(newUser)
		w.Write(newUserJson)
	} else {
		apiError := apiError{
			ErrorCode: http.StatusBadRequest,
		}

		w.WriteHeader(http.StatusBadRequest)

		errorJson, _ := json.Marshal(apiError)
		w.Write(errorJson)
	}
}
