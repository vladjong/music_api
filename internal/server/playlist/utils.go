package playlist

import (
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type Error struct {
	Error string `json:"Error"`
}

func writeResponseJson(w http.ResponseWriter, body interface{}) {
	resp, err := jsoniter.Marshal(body)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
}

func writeError(w http.ResponseWriter, error error, status int) {
	w.WriteHeader(status)
	e := &Error{Error: error.Error()}
	byteMessage, _ := json.Marshal(e)
	_, err := w.Write(byteMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
