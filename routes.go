package main

import "github.com/gin-gonic/gin"

func createRouter() *gin.Engine {
	router := gin.Default()
	templates := NewRender("views")
	templates.AddDirectory("")
	router.HTMLRender = templates

	// static files
	router.Static("/assets", "./assets")

	router.GET("/", landingPageHandler)

	return router
}
