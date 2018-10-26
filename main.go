package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"./controller"
	"./router"
)

var r *router.Router

func newRouter() *router.Router {
	return &router.Router{
		Mux: make(map[string]controller.AppController),
	}
}

func getTemplate(templateName string) *template.Template {
	t, err := template.ParseFiles(filepath.Join("assets", templateName+".html"))
	if err != nil {
		panic("Template missing!")
	}
	return t
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT env var not set")
	}
	r := newRouter()
	r.Add("/", controller.Index{PageTemplate: getTemplate("index")})
	log.Fatal(http.ListenAndServe(":"+port, r))
}
