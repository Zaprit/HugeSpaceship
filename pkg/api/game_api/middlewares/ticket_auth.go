package middlewares

import (
	"HugeSpaceship/pkg/common/db/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TicketAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("MM_AUTH") // TODO: double check that this is the right cookie?
		if err != nil {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		session, exists := auth.GetSession(token)

		if !exists {
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Set("session", session) // TODO: Make this function more secure and more correct, idk
	}
}
