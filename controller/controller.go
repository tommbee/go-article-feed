package controller

import (
	"net/http"
)

// AppController base interfae
type AppController interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}
