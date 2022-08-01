package auth

import (
	"fmt"
	plat_user "rt-msg-carrier/models/plat_user"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		if userName, passWord, ok := c.Request.BasicAuth(); !ok {
			c.AbortWithStatus(401)
		} else {
			if userName != "seu" || passWord != "admin@123" {
				c.AbortWithStatus(401)
			}
		}
	}
}

func Token2User() gin.HandlerFunc {
	return func(c *gin.Context) {
		var platUser plat_user.PlatUser
		reqToken := c.Request.Header.Get("Authorization")
		// headerAuthorization := c.Request.Header["Authorization"][0]
		if reqToken == "awsl5201314" {
			platUser = new(plat_user.ThirdPartyUser)
		} else {
			splitToken := strings.Split(reqToken, "Bearer ")
			if len(splitToken) != 2 {
				// Error: Bearer token not in proper format
				c.AbortWithStatus(401)
			}
			reqToken = strings.TrimSpace(splitToken[1])
			//TODO
			id, err := strconv.Atoi(reqToken)
			if err != nil {
				c.AbortWithStatus(401)
			}
			platUser = plat_user.NewPlatUser(id)
		}

		fmt.Println(platUser)
		c.Next()
	}
}
