package repository

import (
	"github.com/tommbee/go-article-feed/model"
)

// ArticleRepository handles the interface to persistant storage
type ArticleRepository interface {
	Fetch(batch int, page int) ([]*model.Article, error)
	GetByUrl(URL string) (*model.Article, error)
}
