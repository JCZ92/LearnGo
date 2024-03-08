package main

import (
	"fmt"
	"strings"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
func main() {

	fmt.Println("Hello Gin!")

	r := gin.Default() // get an route engine
	r.Use(cors.Default()) // CORS

	// a bare get method
	r.GET("/ping", func(c *gin.Context) {
		c.Writer.WriteString("Hello Joseph, This is Go.\n")
		c.String(200, "Hope you are doing %s", "good") // with status
		// c.JSON(200, gin.H{ //map of string to any
		// 	"message": "pong",
		// })
	})
	// a get method with query
	r.GET("/welcome", func(c *gin.Context) { // note that query is not in the path argument
		who := c.Query("who")
		c.Writer.WriteString("Hello " + who + ", This is Go.")
	})

	r.POST("/post/:name/:welcome", func(c *gin.Context) {
		name := c.Param("name")
		welcome := c.Param("welcome")
		c.Writer.WriteString(strings.ToTitle(welcome) + " " + name)

	})

	r.Run() // listen and serve on 0.0.0.0:8080 
	// change port
	// r.Run(":8080") 
}