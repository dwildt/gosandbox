package main

import "github.com/gin-gonic/gin"
import "net/http"
import "strconv"

func main() {
	r := gin.Default()
	r.GET("/fizzbuzz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "fizz",
		})
	})

	r.GET("/fb/:number", func(c *gin.Context) {
		number := c.Param("number")
		value, _ := strconv.Atoi(number)
		number = fizzbuzz(value)

		c.String(http.StatusOK, "%s", number)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
