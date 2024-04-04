package usecase

import (
	"github.com/adhikag24/go-brick/app/domain"
	"github.com/adhikag24/go-brick/app/repository"
	"github.com/adhikag24/go-brick/app/utils"
)

type UC interface {
	GetUser(userId string) (*domain.User, *utils.ApplicationError)
	GetAllUsers() ([]*domain.User, *utils.ApplicationError)
}

func NewUC(repo repository.Repo) UC {
	return &UserUseCase{rpo: repo}
}
