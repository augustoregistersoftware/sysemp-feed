package usecase

import (
	"context"
	"fmt"
	"strings"
	"sysemp_feed/config"
	"sysemp_feed/repository"
)

type AuthUseCase struct {
	userRepo *repository.UserRepository
}

func NewAuthUsecase(userRepo *repository.UserRepository) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userRepo,
	}
}

func (a *AuthUseCase) ValidateCredentials(ctx context.Context, email, password string) (string, bool, error) {
	email = strings.TrimSpace(email)

	user, err := a.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", false, err
	}

	if user == nil {
		return "", false, nil
	}

	validApprovedUser, err := a.userRepo.IsApprovedUser(ctx, user.ID)
	if err != nil {
		return "", false, err
	}

	if !validApprovedUser {
		return "", false, fmt.Errorf("Usuario não aprovado")
	}

	validPassword, err := config.VerifyPassword(password, user.PasswordHash)
	if err != nil {
		return "", false, err
	}

	if !validPassword {
		return "", false, nil
	}

	return user.Username, true, nil
}
