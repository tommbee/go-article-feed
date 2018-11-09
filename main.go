package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tommbee/go-article-feed/controller"
	"github.com/tommbee/go-article-feed/router"
)

var r *router.Router

func newRouter() *router.Router {
	return &router.Router{
		Mux: make(map[string]controller.AppController),
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT env var not set")
	}
	r := newRouter()
	r.Add("/", controller.Index{Repository: nil}) // Add mongo repo instance
	log.Fatal(http.ListenAndServe(":"+port, r))
}
