package controller

import (
	"encoding/json"
	"net/http"

	"github.com/tommbee/go-article-feed/repository"
)

// Index is the article listing controller
type Index struct {
	Repository repository.ArticleRepository
}

// HandleRequest hnadle the response for a given request
func (i Index) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	articles, err := i.Repository.Fetch(1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(articles)
}
