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
	ro := &repository.MongoArticleRepository{
		Server:       os.Getenv("SERVER"),
		DatabaseName: os.Getenv("DB"),
		AuthDatabase: os.Getenv("AUTH_DB"),
		DBSSL:        os.Getenv("DB_SSL"),
		Collection:   os.Getenv("ARTICLE_COLLECTION"),
		Username:     os.Getenv("DB_USER"),
		Password:     os.Getenv("DB_PASSWORD"),
	}
	return ro
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
	r.Add("/article", controller.Article{Repository: repo})
	r.Add("/articles", controller.Index{Repository: repo})
	r.Add("/heartbeat", controller.Heartbeat{})
	log.Fatal(http.ListenAndServe(":"+port, r))
}
