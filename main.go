package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setting up struct for POST body request
type MyInput struct {
	Word string `form:"word" json:"word" binding:"required"`
}

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	var input MyInput

	r.POST("/", func(c *gin.Context) {
		err := c.ShouldBindJSON(&input)
		if err != nil {
			panic(err)
		}

		fmt.Println("received input " + input.Word)

		englishWord := TranslateWords(input.Word)

		analyzeResult := Annotator(englishWord)

		c.JSON(http.StatusOK, gin.H{"result": analyzeResult})
	})

	r.Run()
}
