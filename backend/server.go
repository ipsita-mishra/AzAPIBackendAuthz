package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "6666"
		log.Printf("Defaulting to port %s", port)
	}

	r := gin.New()

	r.GET("/api/signup", func(c *gin.Context) {
		c.String(http.StatusOK, "Sign In Successful")
	})
	r.GET("/api/signin", func(c *gin.Context) {
		c.String(http.StatusOK, "Sign In Successful")
	})
	r.GET("/api/signout", func(c *gin.Context) {
		c.String(http.StatusOK, "Sign Out Successful")
	})
	r.GET("/api/insensitive", func(c *gin.Context) {
		c.String(http.StatusOK, "API data response post auth - Insensitive")
	})
	r.GET("/api/sensitive", func(c *gin.Context) {
		c.String(http.StatusOK, "API data response post auth - Sensitive")
	})

	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
