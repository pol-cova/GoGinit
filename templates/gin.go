package templates

const GinTemplate = `package main
import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello, Gin!")
    })
    r.Run(":8080")
}`
