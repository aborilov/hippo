package response

import (
	"encoding/json"
	"net/http"

	"github.com/aborilov/hippo/api/sdk/http/errors"
)

func WriteJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		errors.Internal(w, "can't write json to response", err)
	}
}

func WriteJSONWithStatus(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if response != nil {
		json.NewEncoder(w).Encode(response)
	}
}
