package au

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
)

func JwtHandler() {

}

var identityKey = "id"

func NewGinJWTMiddleware() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
	}
}
