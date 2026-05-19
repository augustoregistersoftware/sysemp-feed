package repository

import (
	"fmt"
	"sysemp_feed/config"
	"sysemp_feed/model"
)

type UserRepository struct {
	*Repository
}

func NewUserRepository(baseRepo *Repository) UserRepository {
	return UserRepository{
		Repository: baseRepo,
	}
}

func (ur *UserRepository) CreateUser(user model.User) (int, error) {
	var id_user int
	query, err := ur.DB.Prepare("SELECT id_user FROM users WHERE email = $1")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.Email).Scan(&id_user)
	if err == nil {
		return 409, fmt.Errorf("email already exists")
	}

	query, err = ur.DB.Prepare("INSERT INTO users" +
		"(email, username, password)" +
		" VALUES ($1, $2, $3) RETURNING id_user")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	passwordEncrypted, err := config.NewCryptConfig(user.Password)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.Email, user.Username, passwordEncrypted).Scan(&id_user)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id_user, nil
}
