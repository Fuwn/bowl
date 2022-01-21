package main

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func setupTemplates(router *gin.Engine) {
	tmpl := template.Must(template.New("").ParseFS(
		virtualFilesystem,
		"templates/partials/*.tmpl",
		"templates/pages/*.tmpl"),
	)
	router.SetHTMLTemplate(tmpl)
}
