package main

import (
	"log"
	"net/http"
	"os"

	"./controller"
	"./router"
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
	r.Add("/", controller.Index{})
	log.Fatal(http.ListenAndServe(":"+port, r))
}
