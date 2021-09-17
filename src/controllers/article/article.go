package article

import (
	"pasarwarga-service-api/src/controllers"
	articleService "pasarwarga-service-api/src/services/article"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	data, err := articleService.List(c.DefaultQuery("title", ""), c.DefaultQuery("category", ""))
	if err != nil {
		panic(err)
	}

	c.JSON(200, controllers.GeneralResponse(data))
}

func Show(c *gin.Context) {
	article_id, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		panic(err)
	}
	data, err := articleService.Get(article_id)
	if err != nil {
		panic(err)
	}

	c.JSON(200, controllers.GeneralResponse(data))
}

func Store(c *gin.Context) {
	category_id, err := strconv.Atoi(c.PostForm("category_id"))
	if err != nil {
		panic(err)
	}

	data, err := articleService.Store(
		c.PostForm("title"),
		category_id,
		c.PostForm("content"),
	)
	if err != nil {
		panic(err)
	}

	c.JSON(200, controllers.GeneralResponse(data))
}

func Edit(c *gin.Context) {
	category_id, err := strconv.Atoi(c.PostForm("category_id"))
	if err != nil {
		panic(err)
	}

	article_id, err := strconv.Atoi(c.PostForm("article_id"))
	if err != nil {
		panic(err)
	}

	data, err := articleService.Update(article_id, c.PostForm("title"), category_id, c.PostForm("content"))
	if err != nil {
		panic(err)
	}

	c.JSON(200, controllers.GeneralResponse(data))
}

func Destroy(c *gin.Context) {
	article_id, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		panic(err)
	}

	_, err = articleService.Delete(article_id)
	if err != nil {
		panic(err)
	}

	c.JSON(200, controllers.GeneralResponse(""))
}
