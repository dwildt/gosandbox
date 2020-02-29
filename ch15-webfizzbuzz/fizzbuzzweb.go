package main

import "github.com/gin-gonic/gin"
import "net/http"
import "strconv"

func main() {
	defaultRequest := "/fb/:number"

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "use %s to play with fizzbuzz", defaultRequest)
	})

	r.GET("/fb/:number", func(c *gin.Context) {
		number := c.Param("number")
		value, _ := strconv.Atoi(number)
		result := fizzbuzz(value)

		c.JSON(200, gin.H{
			"value":  value,
			"result": result,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
