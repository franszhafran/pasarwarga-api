package main

import (
	"pasarwarga-service-api/src/controllers/article"
	"pasarwarga-service-api/src/controllers/category"

	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {
	articleGroup := r.Group("/article")
	{
		articleGroup.GET("/", article.Index)
		articleGroup.GET("/:article_id", article.Show)
		articleGroup.POST("/", article.Store)
		articleGroup.DELETE("/:article_id", article.Destroy)
		articleGroup.PATCH("/", article.Edit)
	}

	categoryGroup := r.Group("/category")
	{
		categoryGroup.GET("/", category.Index)
		categoryGroup.GET("/:category_id", category.Show)
		categoryGroup.POST("/", category.Store)
		categoryGroup.DELETE("/", category.Destroy)
		categoryGroup.PATCH("/", category.Edit)
	}
}
