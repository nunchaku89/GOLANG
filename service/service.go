package service // import "service"

import (
	// "database/sql"
	"fmt"
	"log"

	"net/http"

	"github.com/labstack/echo/v4"
	nullable "gopkg.in/guregu/null.v3"

	"model"
	"conndb"
)

/* var (
	db *sql.DB
)

// ConnectToDb - hahaha
func ConnectToDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:13306)/test")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db is connected")
	}
	return db, err
} */

// SelectPersons -
func SelectPersons(c echo.Context) error {
	db, err := conndb.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT IDX, NAME, EMAIL FROM person")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	log.Println(&rows)

	Persons := []model.Person{}

	for rows.Next() {
		var idx nullable.Int
		var name nullable.String
		var email nullable.String

		err := rows.Scan(&idx, &name, &email)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(idx, name, email)

		p := model.Person{
			Idx:   idx,
			Name:  name,
			Email: email,
		}
		Persons = append(Persons, p)
	}

	response := Persons
	return c.JSON(http.StatusOK, response)
}

// SelectWithPaging - 
func SelectWithPaging(c echo.Context) error {
	db, err := conndb.ConnectToDb()
	if err != nil {
		return err
	}
	defer db.Close()

	var limit, offset int
	var Persons = []model.Person{}
	
	rows, err := db.Query("SELECT IDX, NAME, EMAIL FROM PERSON LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		//return err
		fmt.Println(err)
	}
	defer rows.Close()
	
	for rows.Next() {
		var idx nullable.Int
		var name, email nullable.String

		err := rows.Scan(&idx, &name, &email)
		if err != nil {
			return err
		}

		p := model.Person{
			Idx : idx,
			Name : name,
			Email : email,
		}
		Persons = append(Persons, p)
	}

	resopnse := Persons
	return c.JSON(http.StatusOK, resopnse)
}

// InsertPerson -
func InsertPerson(c echo.Context) error {
	db, err := conndb.ConnectToDb()
	if err != nil {
		return err
	}
	defer db.Close()

	var idx int
	var name, email string

	p := new(model.Person)
	if p.Idx.Valid {
		p.Idx.Int64 = int64(idx)
	}
	if p.Name.Valid {
		p.Name.String = name
	}
	if p.Email.Valid {
		p.Email.String = email
	}
	if err := c.Bind(p); err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO PERSON(NAME, EMAIL) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(p.Name.String, p.Email.String)
	if err2 != nil {
		return err2
	}

	fmt.Println(result.LastInsertId())
	return c.JSON(http.StatusOK, p)
}

// UpdatePerson -
func UpdatePerson(c echo.Context) error {
	db, err := conndb.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	p := new(model.Person)
	if err := c.Bind(p); err != nil {
		return err
	}

	stmt, err := db.Prepare("UPDATE PERSON SET NAME = ?, EMAIL= ? WHERE IDX = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err2 :=stmt.Exec(p.Name.String, p.Email.String, p.Idx.Int64)
	if err2 != nil {
		return err2
	}
	fmt.Println(result.LastInsertId())
	return c.JSON(http.StatusOK, "1 row updated")
}

// DeletePersonByIdx -
func DeletePersonByIdx(c echo.Context) error {
	db, err := conndb.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	p := new(model.Person)
	if err := c.Bind(p); err != nil {
		return err
	}

	stmt, err := db.Prepare("DELETE FROM PERSON WHERE IDX = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err2 := stmt.Exec(p.Idx.Int64)
	if err2 != nil {
		return err2
	}
	fmt.Println(result.LastInsertId())
	return c.JSON(http.StatusOK, "1 row deleted")
}
