package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "6666"
		log.Printf("Defaulting to port %s", port)
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.GET("/api/userinfo", func(c *gin.Context) {
		val, ok := c.Request.Header["X-Ms-Client-Principal"]
		decodedUserInfo, _ := base64.StdEncoding.DecodeString((strings.Join(val, "")))
		if ok {
			c.String(http.StatusOK, "Here are your details : %s", decodedUserInfo)
		} else {
			c.String(http.StatusOK, "Please login to view the information.")
		}
	})

	r.GET("/api/privi", func(c *gin.Context) {
		val, ok := c.Request.Header["X-Ms-Client-Principal"]
		decodedUserInfo, _ := base64.StdEncoding.DecodeString((strings.Join(val, "")))
		if ok {
			c.String(http.StatusOK, "Your RBAC role is : %s .Here are details you are allowed to view : %s", decodedUserInfo)
		} else {
			c.String(http.StatusOK, "Please login to view the information.")
		}
	})

	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}
