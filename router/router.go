package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/tommbee/feed-app/controller"
)

// Router object to handle http routing to controllers
type Router struct {
	Mux map[string]controller.AppController
}

// Add a path to the router
func (r *Router) Add(path string, handle controller.AppController) {
	r.Mux[path] = handle
}

// GetPath gets the url path from the request URL
func GetPath(url string) string {
	sl := strings.Split(url, "/")
	return fmt.Sprintf("/%s", sl[1])
}

// ServeHTTP serves the relevant content according to the controller
func (r *Router) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	head := GetPath(rq.URL.Path)
	c, ok := r.Mux[head]
	if ok {
		c.HandleRequest(w, rq)
		return
	}
	http.NotFound(w, rq)
}
