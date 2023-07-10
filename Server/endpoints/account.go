package endpoints

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/NotKatsu/ConsoleChat/authentication"
)

type accountCreated struct {
	SnowFlake     uint64 `json:"snowflake"`
	Authorization string `json:"authorization"`
}

func generate_authentication(unique_id uint64) string {
	return authentication.Encode(strconv.Itoa(int(unique_id))) + "." + authentication.Encode(strconv.Itoa(int(authentication.Since_Epoch()))) + "." + authentication.Encode(authentication.String(25))
}

func AccountCreate(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	min := int64(10000000000000000)
	max := int64(99999999999999999)

	SnowFlake := rand.Int63n(max-min+1) + min
	Authorization := generate_authentication(uint64(SnowFlake))

	newUser := accountCreated{
		SnowFlake:     uint64(SnowFlake),
		Authorization: Authorization,
	}

	w.WriteHeader(http.StatusOK)

	newUserJson, _ := json.Marshal(newUser)
	w.Write(newUserJson)

}
