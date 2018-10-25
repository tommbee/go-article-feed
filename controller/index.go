package controller

import (
	"html/template"
	"net/http"
)

// Index is the article listing controller
type Index struct {
	PageTemplate *template.Template
}

// HandleRequest hnadle the response for a given request
func (i Index) HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	err := i.PageTemplate.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
