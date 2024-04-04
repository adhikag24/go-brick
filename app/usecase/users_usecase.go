package usecase

import (
	"github.com/adhikag24/go-brick/app/domain"
	"github.com/adhikag24/go-brick/app/repository"
	"github.com/adhikag24/go-brick/app/utils"
)

type UserUseCase struct {
	rpo repository.Repo
}

func (u *UserUseCase) GetUser(userId string) (*domain.User, *utils.ApplicationError) {
	user, err := u.rpo.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) GetAllUsers() ([]*domain.User, *utils.ApplicationError) {

	user, err := u.rpo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return user, nil
}
