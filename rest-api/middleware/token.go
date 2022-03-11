package middleware

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func GenerateToken(claims *JwtClaims, expiresAt time.Time, userID string) (string, error) {
	claims.ExpiresAt = expiresAt.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
