package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tommbee/go-article-feed/controller"
	"github.com/tommbee/go-article-feed/repository"
	"github.com/tommbee/go-article-feed/router"
)

var r *router.Router
var repo *repository.ArticleRepository

func newRepo() *repository.MongoArticleRepository {
	r := &repository.MongoArticleRepository{
		Database:   os.Getenv("DB"),
		Collection: os.Getenv("ARTICLE_COLLECTION"),
	}
	r.Connect()
	return r
}

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
	repo := newRepo()
	r.Add("/articles", controller.Index{Repository: repo}) // Add mongo repo instance
	log.Fatal(http.ListenAndServe(":"+port, r))
}
