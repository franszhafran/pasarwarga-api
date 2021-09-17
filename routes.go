package main

import "github.com/gin-gonic/gin"
import "pasarwarga-service-api/src/controllers/article"

func initRoutes(r *gin.Engine) {
	articleGroup := r.Group("/article")
	{
		articleGroup.GET("/", article.Index)
		// articleGroup.POST("/submit", submitEndpoint)
		// articleGroup.POST("/read", readEndpoint)
	}

	// categoryGroup := r.Group("/category")
	// {
	// 	categoryGroup.POST("/login", loginEndpoint)
	// 	categoryGroup.POST("/submit", submitEndpoint)
	// 	categoryGroup.POST("/read", readEndpoint)
	// }
}