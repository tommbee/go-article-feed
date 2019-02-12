package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/mongodb/mongo-go-driver/options"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/tommbee/go-article-feed/model"
)

// MongoArticleRepository interfaces with a mongo db instance
type MongoArticleRepository struct {
	Server       string
	DatabaseName string
	AuthDatabase string
	DBSSL        string
	Collection   string
	Username     string
	Password     string
	ReplicaSet   string
	db           *mongo.Database
}

type key string

const (
	hostKey         = key("hostKey")
	usernameKey     = key("usernameKey")
	passwordKey     = key("passwordKey")
	databaseKey     = key("databaseKey")
	authDatabaseKey = key("authDatabaseKey")
	dBSSL           = key("dBSSL")
)

// Connect to the db instance
func (r *MongoArticleRepository) Connect() {
	log.Print("Connecting...")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ctx = context.WithValue(ctx, hostKey, r.Server)
	ctx = context.WithValue(ctx, usernameKey, r.Username)
	ctx = context.WithValue(ctx, passwordKey, r.Password)
	ctx = context.WithValue(ctx, databaseKey, r.DatabaseName)
	ctx = context.WithValue(ctx, authDatabaseKey, r.AuthDatabase)
	ctx = context.WithValue(ctx, dBSSL, r.DBSSL)
	db, err := configDB(ctx)
	if err != nil {
		log.Fatalf("todo: database configuration failed: %v", err)
	}
	r.db = db
}

func configDB(ctx context.Context) (*mongo.Database, error) {
	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s?authSource=%s&ssl=%s`,
		ctx.Value(usernameKey),
		ctx.Value(passwordKey),
		ctx.Value(hostKey),
		ctx.Value(databaseKey),
		ctx.Value(authDatabaseKey),
		ctx.Value(dBSSL),
	)
	log.Print(uri)
	client, err := mongo.NewClient(uri)
	if err != nil {
		return nil, fmt.Errorf("todo: couldn't connect to mongo: %v", err)
	}
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("todo: mongo client couldn't connect with background context: %v", err)
	}
	dbName := ctx.Value(databaseKey).(string)
	articleDB := client.Database(dbName)
	return articleDB, nil
}

// Fetch all records from the article repository
func (r *MongoArticleRepository) Fetch(batch int, page int) ([]*model.Article, error) {
	r.Connect()
	log.Printf("Getting articles... pages: %d", page)
	var articles []*model.Article
	collection := r.db.Collection(r.Collection)
	bs := int32(batch)
	skip := int64((batch * page))
	cur, err := collection.Find(context.Background(), nil, &options.FindOptions{
		BatchSize: &bs,
		Skip:      &skip,
	})
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		result := &model.Article{}
		err := cur.Decode(&result)
		if err != nil {
			return articles, err
		}
		articles = append(articles, result)
	}
	if err := cur.Err(); err != nil {
		return articles, err
	}
	return articles, err
}

// GetByUrl entity
func (r *MongoArticleRepository) GetByUrl(URL string) (*model.Article, error) {
	r.Connect()
	log.Printf("Getting article... (%s)", URL)
	article := &model.Article{}
	l := int64(1)
	cur, err := r.db.Collection(r.Collection).Find(context.Background(), nil, &options.FindOptions{
		Limit: &l,
	})
	if err != nil {
		return article, err
	}
	error := cur.Decode(&article)
	if error != nil {
		return article, error
	}
	return article, nil
}
