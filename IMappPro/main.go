package main
import (
	"fmt"
	"github.com/gin-gonic/gin"	
)
func main() {

	fmt.Println("Hello Gin!")

	r := gin.Default() // get an engine

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{ //map of string to any
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 
	// change port
	// r.Run(":8080") 
}