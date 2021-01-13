package hash

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/akky/webservice/db"
	"github.com/akky/webservice/model"
	"github.com/gin-gonic/gin"
)

// PostJSON struct
type PostJSON struct {
	Hash    string `json:"hash"`
	Country string `json:"country"`
	Name    string `json:"name"`
}

// GetAll return all records from test_2 table
func GetAll(c *gin.Context) {

	hashModel := getHashModel()

	hashs, err := hashModel.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, hashs)
	return
}

// Save return Hash
func Save(c *gin.Context) {

	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var PostJSON PostJSON
	err2 := json.Unmarshal(jsonData, &PostJSON)
	if err2 != nil {
		log.Print(err2)
		c.JSON(http.StatusBadRequest, err2)
		return
	}

	success, errorMessage := validatePostJSON(PostJSON, true)

	if !success {
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	hashModel := getHashModel()

	code, err3 := hashModel.Save(PostJSON.Hash, PostJSON.Country, PostJSON.Name)

	if err3 != nil {
		c.JSON(http.StatusNotFound, err3)
		return
	}

	c.JSON(http.StatusOK, code)
	return

}

// Delete return
func Delete(c *gin.Context) {

	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var PostJSON PostJSON
	err2 := json.Unmarshal(jsonData, &PostJSON)
	if err2 != nil {
		log.Print(err2)
		c.JSON(http.StatusBadRequest, err2)
		return
	}

	success, errorMessage := validatePostJSON(PostJSON, false)

	if !success {
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	hashModel := getHashModel()

	code, err3 := hashModel.Delete(PostJSON.Hash, PostJSON.Country)

	if err3 != nil {
		if code == 400 {
			c.JSON(http.StatusNotFound, "Cannot find the request.")
			return
		}
		c.JSON(http.StatusInternalServerError, "Internal server error.")
		return
	}

	c.JSON(http.StatusOK, code)
	return
}

func getHashModel() model.HashModel {
	db := db.GetDB()
	hashModel := model.HashModel{
		DB: db,
	}
	return hashModel
}

func validatePostJSON(PostJSON PostJSON, countryCheck bool) (bool, string) {

	errorMessage := ""

	if countryCheck && len(PostJSON.Name) < 1 {
		errorMessage = errorMessage + "Name cannot be empty. "
	}
	if len(PostJSON.Hash) < 1 {
		errorMessage = errorMessage + "Hash cannot be empty. "
	}
	if len(PostJSON.Country) != 2 {
		errorMessage = errorMessage + "Country allow only two characters."
	}

	if len(errorMessage) < 1 {
		return true, errorMessage
	}

	return false, errorMessage
}
