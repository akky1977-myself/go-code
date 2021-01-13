package model

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/akky/webservice/entity"
	"github.com/gin-gonic/gin"
)

// Model is interface of ProductModel
type Model interface {
	FindALl() ([]entity.Product, error)
}

// ProductModel have DB
type ProductModel struct {
	DB *sql.DB
}

// FindAll return
func (ProductModel ProductModel) FindAll() ([]entity.Product, error) {

	rows, err := ProductModel.DB.Query("SELECT * FROM test")
	defer ProductModel.DB.Close()

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return ProductModel.scan(rows)
}

// FindByID return
func (ProductModel ProductModel) FindByID(id string) ([]entity.Product, error) {

	rows, err := ProductModel.DB.Query("SELECT * FROM test WHERE id=?", id)
	defer ProductModel.DB.Close()

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return ProductModel.scan(rows)
}

// Insert return
func (ProductModel ProductModel) Insert(c *gin.Context) ([]entity.Product, error) {

	name := c.DefaultPostForm("name", "")
	stmt, err := ProductModel.DB.Prepare("INSERT test SET name=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err2 := stmt.Exec(name)
	if err2 != nil {
		return nil, err2
	}

	id, err3 := res.LastInsertId()
	if err3 != nil {
		return nil, err3
	}

	return ProductModel.FindByID(strconv.FormatInt(id, 10))
}

// Update return
func (ProductModel ProductModel) Update(c *gin.Context) ([]entity.Product, error) {
	name := c.DefaultPostForm("name", "")
	id := c.Param("id")

	defer ProductModel.DB.Close()

	if !ProductModel.check(id) {
		err := fmt.Errorf("No product found for id: %s", id)
		return nil, err
	}

	stmt, err1 := ProductModel.DB.Prepare("UPDATE test SET name=? WHERE  id= ?")

	if err1 != nil {
		return nil, err1
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(name, id)
	if err2 != nil {
		return nil, err2
	}

	return ProductModel.FindByID(id)
}

// Delete return
func (ProductModel ProductModel) Delete(c *gin.Context) bool {
	id := c.Param("id")

	defer ProductModel.DB.Close()

	if !ProductModel.check(id) {
		log.Printf("No product found for id: %s", id)
		return false
	}

	stmt, err2 := ProductModel.DB.Prepare("DELETE FROM test WHERE  id= ?")

	if err2 != nil {
		log.Println(err2.Error())
		return false
	}
	defer stmt.Close()

	_, err3 := stmt.Exec(id)
	if err3 != nil {
		log.Println(err3.Error())
		return false
	}

	return true
}

// scan return
func (ProductModel ProductModel) scan(rows *sql.Rows) ([]entity.Product, error) {
	products := []entity.Product{}
	for rows.Next() {
		var id int64
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
		product := entity.Product{
			ID:   id,
			Name: name,
		}
		products = append(products, product)
	}
	return products, nil
}

// check return
func (ProductModel ProductModel) check(id string) bool {
	products, err := ProductModel.FindByID(id)
	if err != nil {
		return false
	}

	if len(products) < 1 {
		return false
	}

	return true
}
