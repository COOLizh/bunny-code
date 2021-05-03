// Package token for user's authentication
package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"gitlab.com/greenteam1/task_repo/pkg/models"
)

type tokenClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
}

// Create returns token string with given user's ID
func Create(u *models.User, salt string) (string, error) {
	claims := tokenClaims{
		ID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(salt))
}

// Parse extracts ID from JWT string
func Parse(token, salt string) (int, error) {
	tk, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := tk.Claims.(*tokenClaims); ok && tk.Valid {
		return claims.ID, nil
	}
	return 0, err
}
