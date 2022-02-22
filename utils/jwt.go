package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJwtToken(uid, seconds int64, accessSecret string) (string, error) {
	claims := make(jwt.MapClaims)
	iat := time.Now().Unix()
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	t := jwt.New(jwt.SigningMethodHS256)
	t.Claims = claims
	return t.SignedString([]byte(accessSecret))
}
