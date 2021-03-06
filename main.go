package main

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		c.String(http.StatusOK, "Hello BOJACÁ, by CodeShip")
	})
	router.GET("/codebreaker/setup/:number", func(c *gin.Context){
		number := c.Param("number")
		setSecret(number)
		c.String(http.StatusOK, "secret number configured: %s",number)
	})

	 router.GET("/codebreaker/guess/:number", func(c *gin.Context) {
		name := c.Param("number")
		result := validate(name)
		c.String(http.StatusOK, "Answer: %s", result)
	})

	router.Run(":" + port)
}