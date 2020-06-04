package route

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/server/middleware"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    middleware.UseMiddlewares(r)

    r.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })

    r.GET("/user/:name", func(c *gin.Context) {
        user := c.Params.ByName("name")
        c.JSON(http.StatusOK, gin.H{"user": user})
    })

    return r
}

