package models

// Authed User

type AuthedUser struct {
	UserID   string `json:"-"`
	UserName string
	JWT      string
}

// Airport Service

type AirportDetailsRes struct {
	Distance float64
}

