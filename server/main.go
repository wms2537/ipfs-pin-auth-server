package main

import (
	"fmt"
	"ipfs-pin-auth-server/models"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.JSON(http.StatusForbidden, gin.H{"error": "No Authorization header provided"})
			c.String(http.StatusForbidden, "No Authorization header provided")
			c.Abort()
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		if token == auth {
			c.JSON(http.StatusForbidden, gin.H{"error": "Could not find bearer token in Authorization header"})
			c.Abort()
			return
		}
		var mytoken models.Token
		if err := models.TokensCollection.FindOne(c, bson.M{"secret": token}).Decode(&mytoken); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

func IPFSPin(cid string) error {
	requestURL := fmt.Sprintf("http://localhost:5001/api/v0/pin/add?arg=%s", cid)
	req, err := http.NewRequest(http.MethodPost, requestURL, nil)
	if err != nil {
		return fmt.Errorf("client: could not create request: %s", err)

	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("client: error making http request: %s", err)
	}
	return nil
}

func main() {
	r := gin.New()
	r.Use(IsAuth())

	r.GET("/pin/:cid", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// it would print: "12345"
		cid := c.Param("cid")
		if err := IPFSPin(cid); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
		}
		log.Println(example)
	})

	// Listen and serve on 0.0.0.0:8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
