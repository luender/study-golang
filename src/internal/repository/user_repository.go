package repository

import (
	"api-project/src/internal/model"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUser(email string) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
	UpdateUser(id string, user *model.User) error
	DeleteUser(id string) error
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

func (repo *userRepository) CreateUser(user *model.User) error {
	_, err := repo.DB.Exec(`INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`,
		user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) GetUser(email string) (*model.User, error) {
	var user model.User

	err := repo.DB.QueryRow(`SELECT * FROM users WHERE email = $1`, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User

	err := repo.DB.QueryRow(`SELECT * FROM users WHERE id = $1`, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (repo *userRepository) UpdateUser(id string, user *model.User) error {
	_, err := repo.DB.Exec(`UPDATE users SET name = $1, email = $2 WHERE id = $3`,
		user.Name, user.Email, id)

	if err != nil {
		return err
	}

	return nil
}

func (repo *userRepository) DeleteUser(id string) error {
	_, err := repo.DB.Exec(`DELETE FROM users WHERE id = $1`, id)

	if err != nil {
		return err
	}

	return nil
}


