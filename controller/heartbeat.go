package controller

import (
	"encoding/json"
	"net/http"
	"os"
)

// Heartbeat is the heartbeat controller
type Heartbeat struct {
}

// HandleRequest hnadle the response for a given request
func (i Heartbeat) HandleRequest(w http.ResponseWriter, r *http.Request) {
	commit := os.Getenv("APP_COMMIT_REF")
	buildDate := os.Getenv("APP_BUILD_DATE")
	data := make(map[string]string, 2)
	data["date"] = buildDate
	data["commit"] = commit
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
