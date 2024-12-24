// main.go

package main

import (
    "github.com/gin-gonic/gin"
    "github.com/pk1151222/bug-scanners/internal/scanner"
)

func main() {
    router := gin.Default()

    router.GET("/scan/:url", func(c *gin.Context) {
        url := c.Param("url")
        result, err := scanner.ScanURL(url)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, result)
    })

    router.Run(":8080")
}
