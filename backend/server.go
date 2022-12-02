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

	r.GET("/api/userinfo", func(c *gin.Context) {
		c.String(http.StatusOK, "Sign In Successful. Here are your details : %s", c.Request.Header["X-Ms-Client-Principal"])
	})

	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
