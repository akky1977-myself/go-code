package main

import (
	"net/http"

	"github.com/akky/webservice/handler/hash"
	"github.com/akky/webservice/handler/product"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/product", func(c *gin.Context) {
		c.JSON(http.StatusOK, product.GetAll())
	})
	r.GET("/product/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, product.View(c))
	})
	r.POST("/product", func(c *gin.Context) {
		c.JSON(http.StatusOK, product.Create(c))
	})
	r.PUT("/product/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, product.Update(c))
	})
	r.DELETE("/product/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, product.Delete(c))
	})
	r.GET("/hash", func(c *gin.Context) {
		hash.GetAll(c)
	})
	r.POST("/hash", func(c *gin.Context) {
		hash.Save(c)
	})
	r.DELETE("/hash", func(c *gin.Context) {
		hash.Delete(c)
	})
	r.Run(":8999")
}
