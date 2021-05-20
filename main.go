package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	indoWords := TranslateWords("Selamat Pagi Semua. Senang rasanya bisa berada di sini")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": indoWords})
	})

	r.Run()
}
