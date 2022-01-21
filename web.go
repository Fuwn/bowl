package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func mountRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", map[string]interface{}{})
	})

	router.POST("/api/v1/entry", func(c *gin.Context) {
		formType := c.PostForm("type")
		formNotes := c.PostForm("notes")
		formAccessCode := c.PostForm("access_code")

		if formAccessCode == accessCode {
			log.Println(formType, formNotes, formAccessCode)

			write(formType, formNotes, c.ClientIP())

			c.HTML(http.StatusOK, "notice.tmpl", map[string]interface{}{
				"Notice": "submission successfully logged",
				"Data":   fmt.Sprintf("{ type: \"%s\", notes: \"%s\" }", formType, formNotes),
			})
		} else {
			log.Printf("invalid access code: %s, expected: %s\n", formAccessCode, accessCode)

			c.HTML(http.StatusUnauthorized, "notice.tmpl", map[string]interface{}{
				"Notice": "invalid access code",
			})
		}
	})
}
