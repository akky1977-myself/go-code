package product

import (
	"github.com/akky/webservice/db"
	"github.com/akky/webservice/entity"
	"github.com/akky/webservice/model"
	"github.com/gin-gonic/gin"
)

// GetAll return products
func GetAll() []entity.Product {
	ProductModel := getProductModel()
	products, err := ProductModel.FindAll()
	ProductModel.DB.Close()
	if err != nil {
		panic(err.Error())
	}
	return products
}

// View return products
func View(c *gin.Context) []entity.Product {
	id := c.Param("id")
	ProductModel := getProductModel()
	products, err := ProductModel.FindByID(id)
	ProductModel.DB.Close()
	if err != nil {
		panic(err.Error())
	}
	return products
}

// Create return products
func Create(c *gin.Context) []entity.Product {
	ProductModel := getProductModel()
	products, err := ProductModel.Insert(c)
	if err != nil {
		panic(err.Error())
	}
	return products
}

// Update return products
func Update(c *gin.Context) []entity.Product {
	ProductModel := getProductModel()
	products, err := ProductModel.Update(c)
	if err != nil {
		panic(err.Error())
	}
	return products
}

// Delete return products
func Delete(c *gin.Context) bool {
	ProductModel := getProductModel()
	return ProductModel.Delete(c)
}

func getProductModel() model.ProductModel {
	db := db.GetDB()
	ProductModel := model.ProductModel{
		DB: db,
	}
	return ProductModel
}
