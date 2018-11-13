package repository

import (
	"log"

	"github.com/tommbee/go-article-feed/model"

	"gopkg.in/mgo.v2"
)

// MongoArticleRepository interfaces with a mongo db instance
type MongoArticleRepository struct {
	Database   string
	Collection string
}

//var db *mgo.Database
var db *mgo.Session

// Connect to the db
func (r *MongoArticleRepository) Connect() {
	db, err := mgo.Dial(r.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMode(mgo.Monotonic, true)
}

// Fetch all records from the article repository
func (r *MongoArticleRepository) Fetch(num int64) ([]*model.Article, error) {
	sessionCopy := db.Copy()
	defer sessionCopy.Close()
	c := sessionCopy.DB(r.Database).C(r.Collection)
	var articles []*model.Article
	log.Print("Getting articles")
	err := c.Find(nil).All(&articles)
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
