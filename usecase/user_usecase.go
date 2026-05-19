package usecase

import (
	"context"
	"sysemp_feed/model"
	"sysemp_feed/repository"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return UserUseCase{
		repository: userRepo,
	}
}

func (u *UserUseCase) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	_, err := u.repository.CreateUser(ctx, user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
