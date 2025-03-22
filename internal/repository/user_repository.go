package repository

import (
	"database/sql"
	"github.com/Elex1337/user-service/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserByID(id int) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id int) error
}

type UserRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) CreateUser(user entity.User) (entity.User, error) {

	query := `INSERT INTO users 
              (user_name, password, created_at, updated_at) 
              VALUES 
              (:user_name, :password, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) 
              RETURNING id, user_name, password, created_at, updated_at`

	var createdUser entity.User

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return entity.User{}, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&createdUser)
		if err != nil {
			return entity.User{}, err
		}
	}

	return createdUser, nil
}

func (r *UserRepositoryImpl) GetUserByID(id int) (entity.User, error) {
	var user entity.User
	err := r.db.Get(&user, "SELECT id, user_name, created_at, updated_at FROM users WHERE id = $1", id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user entity.User) (entity.User, error) {

	query := `UPDATE users SET 
              user_name = :user_name, 
              password = :password, 
              updated_at = CURRENT_TIMESTAMP 
              WHERE id = :id 
              RETURNING id, user_name, password, created_at, updated_at`

	rows, err := r.db.NamedQuery(query, user)
	if err != nil {
		return entity.User{}, err
	}
	defer rows.Close()

	var updatedUser entity.User
	if rows.Next() {
		err = rows.StructScan(&updatedUser)
		if err != nil {
			return entity.User{}, err
		}
	} else {
		return entity.User{}, sql.ErrNoRows
	}

	return updatedUser, nil
}
func (r *UserRepositoryImpl) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = $1"

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
