package repository

import (
	"../model"
)

// ArticleRepository handles the interface to persistant storage
type ArticleRepository interface {
	Fetch(cursor string, num int64) ([]*model.Article, error)
	GetByID(id int64) (*model.Article, error)
	GetByTitle(title string) (*model.Article, error)
	GetByUrl(URL string) (*model.Article, error)
	Update(article *model.Article) (*model.Article, error)
	Store(a *model.Article) (int64, error)
	Delete(id int64) (bool, error)
}
