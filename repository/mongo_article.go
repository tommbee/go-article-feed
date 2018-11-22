package repository

import (
	"log"

	"github.com/tommbee/go-article-feed/model"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoArticleRepository interfaces with a mongo db instance
type MongoArticleRepository struct {
	Server       string
	DatabaseName string
	Collection   string
	session      *mgo.Session
}

// Connect to the db instance
func (r *MongoArticleRepository) Connect() {
	session, err := mgo.Dial(r.Server)
	if err != nil {
		log.Fatal(err)
	}
	r.session = session
}

// Fetch all records from the article repository
func (r *MongoArticleRepository) Fetch(num int64) ([]*model.Article, error) {
	r.Connect()
	defer r.session.Close()
	log.Print("Getting articles")
	var articles []*model.Article
	err := r.session.DB(r.DatabaseName).C(r.Collection).Find(bson.M{}).All(&articles)
	return articles, err
}

// GetByID an entity
func (r *MongoArticleRepository) GetByID(id int64) (*model.Article, error) {
	return nil, nil
}

// GetByTitle entity
func (r *MongoArticleRepository) GetByTitle(title string) (*model.Article, error) {
	return nil, nil
}

// GetByUrl entity
func (r *MongoArticleRepository) GetByUrl(URL string) (*model.Article, error) {
	return nil, nil
}
