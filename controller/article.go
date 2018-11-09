package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/tommbee/go-article-feed/model"
)

// Article controller to handle article routes
type Article struct{}

func getArticle(id int64) (model.Article, error) {
	// Query repo
	return model.Article{}, errors.New("")
}

func (a Article) handleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	article, err := getArticle(1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(article)
}
