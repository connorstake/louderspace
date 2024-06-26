package repositories

import (
	"database/sql"
	"louderspace/internal/models"
)

type UserStorage interface {
	Save(user *models.User) error
	UserByID(userID int) (*models.User, error)
	UserByUsername(username string) (*models.User, error)
	Users() ([]*models.User, error)
}

type UserDatabase struct {
	db *sql.DB
}

func NewUserDatabase(db *sql.DB) UserStorage {
	return &UserDatabase{db}
}

func (r *UserDatabase) Save(user *models.User) error {
	query := "INSERT INTO users (username, password, email, role, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	return r.db.QueryRow(query, user.Username, user.Password, user.Email, user.Role, user.CreatedAt).Scan(&user.ID)
}

func (r *UserDatabase) UserByID(userID int) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, username, email, role, created_at FROM users WHERE id = $1"
	if err := r.db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserDatabase) UserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, username, email, password, role, created_at FROM users WHERE username = $1"
	if err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserDatabase) Users() ([]*models.User, error) {
	query := "SELECT id, username, email, role, created_at FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
