package main

import (
	"embed"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

//go:embed templates
var virtualFilesystem embed.FS

var accessCode string

func init() {
	if gin.Mode() == gin.DebugMode {
		accessCode = "test"
	} else {
		accessCode = os.Getenv("ACCESS_CODE")
	}
}

func main() {
	router := gin.Default()

	setupTemplates(router)
	mountRoutes(router)

	err := router.RunTLS(":8080", ".bowl/certificates/bowl.crt", ".bowl/certificates/bowl.key")
	if err != nil {
		log.Fatalf("error starting bowl: %s\n", err)

		return
	}
}
