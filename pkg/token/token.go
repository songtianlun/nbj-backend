package token

import (
	"github.com/dgrijalva/jwt-go"
	"minepin-backend/config"
	"time"
)

// Context is the context of the JSON web token.
type Context struct {
	ID       uint64
	UUID     string
	nbf      int64
	iat      int64
	exp      int64
}

func Sign(c Context, secret string) (tokenString string, err error) {
	if secret == "" {
		secret = config.GetString("jwt_secret")
	}
	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"uuid":     c.UUID,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Unix() + 20,
	})
	tokenString, err = token.SignedString([]byte(secret))
	return
}
