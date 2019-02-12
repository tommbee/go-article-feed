package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tommbee/go-article-feed/repository"
)

// Index is the article listing controller
type Index struct {
	Repository repository.ArticleRepository
}

// HandleRequest hnadle the response for a given request
func (i Index) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	page := 1
	keys, ok := r.URL.Query()["page"]
	if ok && len(keys[0]) > 0 {
		i, err := strconv.Atoi(keys[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		page = i
	}
	articles, err := i.Repository.Fetch(100, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(articles)
}
