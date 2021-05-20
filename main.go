package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyInput struct {
	word string
}

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	var input MyInput

	r.GET("/", func(c *gin.Context) {
		englishWord := TranslateWords("Selamat Pagi Semua. Senang rasanya bisa berada di sini")
		c.JSON(http.StatusOK, gin.H{"data": englishWord})
	})

	r.POST("/", func(c *gin.Context) {
		err := c.ShouldBindJSON(&input)
		if err != nil {
			panic(err)
		}

		fmt.Println(input.word)

		englishWord := TranslateWords(input.word)

		c.JSON(http.StatusOK, gin.H{"data": englishWord})
	})

	r.Run()
}
