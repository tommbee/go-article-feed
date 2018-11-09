package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/tommbee/go-article-feed/model"
	"github.com/tommbee/go-article-feed/repository"
)

// Index is the article listing controller
type Index struct {
	Repository *repository.ArticleRepository
}

func getLatestArticles() ([]*model.Article, error) {
	// Query repo
	return nil, errors.New("")
}

// HandleRequest hnadle the response for a given request
func (i Index) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	articles, err := getLatestArticles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(articles)
}
