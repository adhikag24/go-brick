package repository

import (
	"github.com/adhikag24/go-brick/app/domain"
	"github.com/adhikag24/go-brick/app/utils"
)

type Repo interface {
	GetAllUsers() ([]*domain.User, *utils.ApplicationError)
	GetUser(userID string) (*domain.User, *utils.ApplicationError)
}

func NewRepo() Repo {
	return &Repository{}
}
