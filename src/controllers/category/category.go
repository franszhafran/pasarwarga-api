package category

import (
	"pasarwarga-service-api/src/controllers"
	"pasarwarga-service-api/src/services/category"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	data, err := category.List()
	if err != nil {
		panic(err)
	}

	c.JSON(200, controllers.GeneralResponse(data))
}

func Show(c *gin.Context) {
	category_id, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		panic(err)
	}

	data, err := category.Get(category_id)
	if err != nil {
		panic(err)
	}
	if data.ID == 0 {
		controllers.Return404(c)
	} else {
		c.JSON(200, controllers.GeneralResponse(data))
	}
}

func Store(c *gin.Context) {
	r, err := category.Store(c.PostForm("name"))
	if err != nil {
		panic(err)
	}
	c.JSON(200, controllers.GeneralResponse(r))
}

func Edit(c *gin.Context) {
	category_id, err := strconv.Atoi(c.PostForm("category_id"))
	if err != nil {
		panic(err)
	}

	r, err := category.Update(category_id, c.PostForm("name"))
	if err != nil {
		panic(err)
	}

	c.JSON(200, controllers.GeneralResponse(r))
}

func Destroy(c *gin.Context) {
	category_id, err := strconv.Atoi(c.Query("category_id"))
	if err != nil {
		panic(err)
	}

	r, err := category.Delete(category_id)
	if err != nil {
		panic(err)
	}
	c.JSON(200, controllers.GeneralResponse(r))
}
