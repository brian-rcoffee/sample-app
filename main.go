package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func landingPageHandler(c *gin.Context) {
	c.HTML(200, "index", gin.H{"foo": "bar"})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)

	router := createRouter()
	router.Run(fmt.Sprintf(":%s", port))
}
