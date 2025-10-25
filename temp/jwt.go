package temp

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	key []byte
}

func New() *JWT {
	key := os.Getenv("JWT_SECRET")
	return &JWT{key: []byte(key)}
}

func (j *JWT) Issue(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	jwtStr, _ := token.SignedString(j.key)
	return jwtStr
}

func (j *JWT) Parse(tokenStr string) string {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.key), nil
	})
	if err != nil || !token.Valid {
		return ""
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims["email"].(string)
}
