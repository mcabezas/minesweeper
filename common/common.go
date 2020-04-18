package common

import (
	"encoding/json"
	"log"
	"net/http"
)

func DecodePost(r *http.Request, structure interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Println("Error parsing post data")
		return err
	}
	return nil
}
