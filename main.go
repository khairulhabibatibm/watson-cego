package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// const INDEX = `<!DOCTYPE html>
// <html>
//   <head>
//     <title>Powered By Paketo Buildpacks</title>
//   </head>
//   <body>
// 	<h1>Hello World</h1>
//     <img style="display: block; margin-left: auto; margin-right: auto; width: 50%;" src="https://paketo.io/images/paketo-logo-full-color.png"></img>
//   </body>
// </html>`

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	fmt.Fprint(w, INDEX)
	// })

	// log.Fatal(http.ListenAndServe(":8080", nil))

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}
