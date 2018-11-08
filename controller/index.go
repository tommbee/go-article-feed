package controller

import (
	"encoding/json"
	"errors"
	"net/http"

	"../model"
	"../repository"
)

// Index is the article listing controller
type Index struct {
	repository *repository.ArticleRepository
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
