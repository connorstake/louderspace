package models

import "time"

type Role string

const (
	RoleFree    Role = "free"
	RolePremium Role = "premium"
	RoleAdmin   Role = "admin"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Role      Role      `json:"role"`
}
