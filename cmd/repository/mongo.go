package repository

import (
	"gitlab.bitkub.io/bo/gokub-boilerplate/cmd/domain/user"
)

// MongoDBRepo represent mongodb type
type MongoDBRepo struct{}

type MongoConfigs struct{}

// NewMongoDB instant new mongo repo
// TODO: dependencies injection mongo client
func NewMongoDB(cfg *MongoConfigs) (*MongoDBRepo, error) {
	// mongo logic connection here
	return &MongoDBRepo{}, nil
}

func (r *MongoDBRepo) Create(usr *user.Model) (*user.User, error) {
	return &user.User{
		Name:  "test",
		Email: "test@test.com",
	}, nil
}
