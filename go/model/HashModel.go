package model

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/akky/webservice/entity"
)

// HashInterface is interface of HashModel
type HashInterface interface {
	Save() ([]entity.Hash, error)
}

// HashModel has DB
type HashModel struct {
	DB *sql.DB
}

// FindAll return all records from test_2 table
func (HashModel HashModel) FindAll() ([]entity.Hash, error) {

	rows, err := HashModel.DB.Query("SELECT * FROM test_2")

	defer HashModel.DB.Close()

	if err != nil {
		return nil, err
	}

	hashs, err2 := HashModel.scan(rows)

	if err2 != nil {
		return nil, err2
	}

	return hashs, nil
}

// FindByID return
func (HashModel HashModel) FindByID() ([]entity.Hash, error) {
	rows, err := HashModel.DB.Query("SELECT * FROM hash")

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	hashs, err1 := HashModel.scan(rows)

	if err1 != nil {
		return hashs, err
	}

	return hashs, nil
}

// Save return Hash
func (HashModel HashModel) Save(hash string, country string, name string) (int64, error) {
	hashExist, err5 := HashModel.check(hash, country)

	if err5 != nil {
		return 500, err5
	}

	if hashExist {
		stmt, err := HashModel.DB.Prepare("UPDATE test_2 SET name=? WHERE  hash= ? AND country= ?")
		if err != nil {
			return 500, err
		}
		defer stmt.Close()
		_, err1 := stmt.Exec(name, hash, country)
		if err1 != nil {
			return 500, err
		}
	} else {
		stmt, err2 := HashModel.DB.Prepare("INSERT test_2 SET hash=?, country=?, name=?")
		if err2 != nil {
			return 500, err2
		}
		defer stmt.Close()

		_, err3 := stmt.Exec(hash, country, name)
		if err3 != nil {
			return 500, err3
		}
	}
	return 200, nil
}

// FindByHashAndCode return
func (HashModel HashModel) FindByHashAndCode(hash string, country string) ([]entity.Hash, error) {
	rows, err := HashModel.DB.Query("SELECT * FROM test_2 WHERE hash=? AND country=?", hash, country)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	hashs, err1 := HashModel.scan(rows)

	if err1 != nil {
		return nil, err
	}

	return hashs, nil
}

// Delete return
func (HashModel HashModel) Delete(hash string, country string) (int16, error) {
	hashExist, err := HashModel.check(hash, country)

	if err != nil {
		return 500, err
	}

	if !hashExist {
		err1 := errors.New("cannot find the request")
		return 400, err1
	}

	stmt, err2 := HashModel.DB.Prepare("DELETE FROM test_2 WHERE  hash= ? AND country=?")

	if err2 != nil {
		return 500, err2
	}
	defer stmt.Close()

	_, err3 := stmt.Exec(hash, country)
	if err3 != nil {
		return 500, err3
	}

	return 200, nil

}

// scan return
func (HashModel HashModel) scan(rows *sql.Rows) ([]entity.Hash, error) {
	hashs := []entity.Hash{}

	for rows.Next() {
		var id int64
		var name string
		var hash string
		var country string
		err := rows.Scan(&id, &name, &hash, &country)
		if err != nil {
			return nil, err
		}
		hashEntity := entity.Hash{
			ID:      id,
			Hash:    hash,
			Country: country,
			Name:    name,
		}
		hashs = append(hashs, hashEntity)
	}
	return hashs, nil
}

// check return
func (HashModel HashModel) check(hash string, country string) (bool, error) {

	hashs, err := HashModel.FindByHashAndCode(hash, country)

	fmt.Print(len(hashs))

	if err != nil {
		return false, err
	}

	if len(hashs) < 1 {
		return false, nil
	}

	return true, nil
}
