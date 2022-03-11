package models

// Authed User

type AuthedUser struct {
	UserID   string `json:"-"`
	UserName string
	JWT      string
}

