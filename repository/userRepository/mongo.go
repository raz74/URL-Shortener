package userRepository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"shortened_link/model"
	"shortened_link/repository"
)

type mongoRepo struct {
	Mongodb    *mongo.Client
	collection *mongo.Collection
}

func NewMongoUserRepositoryImpl(mongodb *mongo.Client) repository.UserRepository {
	return &mongoRepo{
		Mongodb:    mongodb,
		collection: mongodb.Database("Shortener_Url").Collection("user"),
	}
}

func (m *mongoRepo) CreateUser(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (m *mongoRepo) CheckUniqueEmail(Email string) error {
	//TODO implement me
	panic("implement me")
}

func (m *mongoRepo) GetUserByEmail(Email string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
