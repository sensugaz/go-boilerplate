package services

import (
	"gitlab.bitkub.io/bo/gokub-boilerplate/cmd/domain/user"
)

//go:generate mockgen -source=services.go -destination=./mock/repository.go UserRepository

// UserRepository represent database, model, layer interface
type UserRepository interface {
	// Create creates a new user
	Create(usr *user.Model) (*user.User, error)
}
