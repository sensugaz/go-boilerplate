package services

import (
	"fmt"

	"gitlab.bitkub.io/bo/gokub-boilerplate/cmd/domain/user"
)

// UserService represent user service type
type UserService struct {
	usrRepo UserRepository
}

// NewUserService instant a new user service
func NewUserService(rpUsr UserRepository) (*UserService, error) {
	// service layer logic ....

	return &UserService{
		usrRepo: rpUsr,
	}, nil
}

func (s *UserService) Create(*user.User) (*user.User, error) {
	usrRepo, err := s.usrRepo.Create(&user.Model{})
	if err != nil {
		return nil, fmt.Errorf("[UserService.Create]: unable to create new user %w", err)
	}

	return &user.User{
		Name:  usrRepo.Name,
		Email: usrRepo.Email,
	}, nil
}
