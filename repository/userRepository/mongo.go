package userRepository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"shortened_link/model"
	"shortened_link/repository"
)

type mongoRepo struct {
	Mongodb    *mongo.Client
	collection *mongo.Collection
}

func NewMongoUserRepoImpl(mongodb *mongo.Client) repository.UserRepository {
	return &mongoRepo{
		Mongodb:    mongodb,
		collection: mongodb.Database("Shortener_Url").Collection("users"),
	}
}

func (m *mongoRepo) CreateUser(user *model.User) error {
	_, err := m.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return err
}

func (m *mongoRepo) CheckUniqueEmail(Email string) error {
	var user model.User
	//If find the document, the err will be nil
	//err := m.collection.FindOne(context.T ODO(), bson.D{{"email", Email}}).Decode(&user)
	//if err == nil {
	//	err = mongo.ErrNoDocuments
	//	return err
	//}
	count, err := m.collection.CountDocuments(context.TODO(), bson.D{{"email", Email}})
	if err != nil {
		return err
	}
	if count >= 0 {
		return mongo.ErrNoDocuments
	}
	fmt.Printf("%+v", user)
	return nil
}

func (m *mongoRepo) GetUserByEmail(Email string) (*model.User, error) {
	var user model.User
	err := m.collection.FindOne(context.TODO(), bson.D{{"email", Email}}).Decode(&user)
	if err != nil {
		return nil, mongo.ErrNoDocuments
	}
	return &user, nil
}
