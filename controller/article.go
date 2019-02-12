package controller

import (
	"encoding/json"
	"net/http"

	"github.com/tommbee/go-article-feed/repository"
)

// Article is the article listing controller
type Article struct {
	Repository repository.ArticleRepository
}

func (a Article) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	article, err := a.Repository.GetByUrl("asdf")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(article)
}
