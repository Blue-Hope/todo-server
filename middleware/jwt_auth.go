package middleware

import (
    "fmt"
    "github.com/gin-gonic/gin"

)

func AuthMiddleware(c *gin.Context) {
    fmt.Println("JWT Auth Check.. ")
    c.Next()
}

func UseAuthMiddleware(r *gin.Engine) {
    r.Use(AuthMiddleware)
}

