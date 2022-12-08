package main

import (
	"encoding/base64"
	"encoding/json"
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
		var id identity
		json.Unmarshal([]byte(decodedUserInfo), &id)
		if ok {
			for _, role := range id.UserRoles {
				if role == "authenticated" {

					c.String(http.StatusOK, "You can see this because you are -> %s .Your RBAC role(s) are : %s . ", role, id.UserRoles)
				}
			}
		} else {
			c.String(http.StatusOK, "Please login to view the information.")
		}
	})

	r.GET("/api/data", func(c *gin.Context) {
		var arr string
		for k, vals := range c.Request.Header {
			arr = arr + k + "#" + (strings.Join(vals, ""))
		}
		// var id identity
		// json.Unmarshal([]byte(decodedUserInfo), &id)
		c.String(http.StatusOK, "Data -> %s", arr)
	})

	log.Printf("Listening on port %s", port)
	r.Run(":" + port)
}

type identity struct {
	IdentityProvider string   `json:"identityProvider"`
	UserID           string   `json:"userId"`
	UserDetails      string   `json:"userDetails"`
	UserRoles        []string `json:"userRoles"`
}
