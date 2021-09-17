package controllers

import (
	"github.com/gin-gonic/gin"
)

func GeneralResponse(r interface{}) gin.H {
	return gin.H{
		"code":    200,
		"message": "OK",
		"data":    r,
	}
}

func Return404(c *gin.Context) {
	c.JSON(404, gin.H{
		"code":    404,
		"message": "Not Found",
	})
}
