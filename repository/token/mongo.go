package token

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"shortened_link/model"
	"shortened_link/repository"
	"time"
)

type mongoRepo struct {
	Mongodb    *mongo.Client
	collection *mongo.Collection
}

func NewMongoTokenRepoImp(mongodb *mongo.Client) repository.TokenRepo {
	return &mongoRepo{
		Mongodb:    mongodb,
		collection: mongodb.Database("Shortener_Url").Collection("token"),
	}
}

func (m mongoRepo) Create(user *model.User) (*model.SessionCookie, error) {
	//create a new random cookie session
	cookieToken := uuid.NewString()
	expiresAt := time.Now().Add(7 * time.Second)
	var token model.SessionCookie
	token = model.SessionCookie{
		UserID: user.Id,
		Value:  cookieToken,
		Expire: expiresAt,
	}
	//mongodb
	_, err := m.collection.InsertOne(context.TODO(), token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (m mongoRepo) Get(header string) (*model.SessionCookie, error) {
	substr := header[6:]
	println(substr)
	var token model.SessionCookie
	//get mongo
	err := m.collection.FindOne(context.TODO(), bson.D{{"value", substr}}).Decode(&token)
	if err != nil {
		return nil, mongo.ErrNoDocuments
	}
	// check the expire
	if token.Expire.Before(time.Now()) {
		_, err2 := m.collection.DeleteOne(context.TODO(), bson.D{{"value", substr}})
		if err2 != nil {
			return nil, err2
		}
	}

	return &token, nil
}
