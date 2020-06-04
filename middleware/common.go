package middleware

import "github.com/gin-gonic/gin"

func UseMiddlewares(r *gin.Engine) {
	UseAuthMiddleware(r)
}
