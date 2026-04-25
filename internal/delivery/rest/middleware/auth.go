package middleware

import (
	"context"
	"net/http"
	"strings"

	constant "github.com/rchmachina/rach-fw/internal/const"
	logg "github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	"github.com/rchmachina/rach-fw/internal/utils/helper"
	jwtToken "github.com/rchmachina/rach-fw/internal/utils/jwt"

	"github.com/gin-gonic/gin"
)

func Auth(logger logg.Logger, secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		logger := logger.WithCtx(c)
		token := c.GetHeader("Authorization")
		if token == "" {
			logger.Error("unauthorized", logg.ToField("unauthorized", "token not valid"))
			helper.JSONResponse(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		parts := strings.SplitN(token, "Bearer ", 2)
		if len(parts) != 2 {
			logger.Error("unauthorized", logg.ToField("unauthorized", "token split is less than 2 word"))
			helper.JSONResponse(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		token = parts[1]

		claims, err := jwtToken.DecodeToken(token, secretKey)
		if err != nil {
			helper.JSONResponse(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), constant.UserInfoKey, claims)
		c.Set("userInfo", claims)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
