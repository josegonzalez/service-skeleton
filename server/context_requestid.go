package server

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

const headerXRequestID = "X-Request-ID"

func ContextRequestidMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("request_id", requestid.Get(ctx))
		ctx.Next()
	}
}
