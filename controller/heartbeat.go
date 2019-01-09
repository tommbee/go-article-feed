package controller

import (
	"encoding/json"
	"net/http"
)

// Heartbeat is the heartbeat controller
type Heartbeat struct {
}

// HandleRequest hnadle the response for a given request
func (i Heartbeat) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("ok")
}
