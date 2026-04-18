package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"
	helper "github.com/rchmachina/rach-fw/internal/utils/helper"
	jwtToken "github.com/rchmachina/rach-fw/internal/utils/jwt"

	"github.com/gin-gonic/gin"
)


type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}



func Auth(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, Result{Code: http.StatusUnauthorized, Message: "unauthorized"})
			c.Abort()
			return
		}

		token = strings.Split(token, "Bearer ")[1]

		claims, err := jwtToken.DecodeToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, Result{Code: http.StatusUnauthorized, Message: "unauthorized"})
			c.Abort()
			return
		}

		expiry, ok := claims["expiry"].(float64)
		if !ok {
			log.Println("expiry value is not of type float64")
			return
			
		}
		roles, ok := claims["roles"].(string)
		if !ok {
			log.Println("expiry value is not of type float64")
			return
		}
		if float64(time.Now().Unix()) > expiry {
			helper.JSONResponse(c,401,"token already expired")
			return
		}
		c.Set("roles",roles)
		handlerFunc(c)
	}
}