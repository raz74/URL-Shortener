package url

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"shortened_link/model"
	"shortened_link/repository"
)

type mongoRepo struct {
	Mongodb    *mongo.Client
	collection *mongo.Collection
}

func (m mongoRepo) Create(shortUrl *model.ShortedUrl) (*model.ShortedUrl, error) {
	_, err := m.collection.InsertOne(context.TODO(), shortUrl)
	if err != nil {
		return nil, err
	}
	return shortUrl, nil
}

func (m mongoRepo) Count() int64 {
	return 10
}

func NewMongoUrlRepoImp(mongodb *mongo.Client) repository.UrlRepository {
	return &mongoRepo{
		Mongodb:    mongodb,
		collection: mongodb.Database("Shortener_Url").Collection("ShortUrl"),
	}
}
