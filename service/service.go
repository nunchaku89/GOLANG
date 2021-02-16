package service // import "service"

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	nullable "gopkg.in/guregu/null.v3"

	"conndb"
	"model"
)

// Login -
func Login(c echo.Context) error {
	var idx int
	var name, email string
	auth := new(model.JwtCustomClaims)
	auth.Idx = idx
	auth.Name = name
	auth.Email = email
	if err := c.Bind(auth); err != nil {
		return err
	}

	if auth.Name != "Joe" || auth.Email != "joe@abc.com" {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := model.JwtCustomClaims{
		Idx:   auth.Idx,
		Name:  auth.Name,
		Email: auth.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// Accessible -
func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// Restricted -
func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

// SelectPersons -
func SelectPersons(c echo.Context) error {
	db, err := conndb.ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT IDX, NAME, EMAIL FROM PERSON")
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
	pg := new(model.Paging)
	pg.Limit = limit
	pg.Offset = offset
	if err := c.Bind(pg); err != nil {
		return err
	}

	var Persons = []model.Person{}

	rows, err := db.Query("SELECT IDX, NAME, EMAIL FROM PERSON LIMIT ? OFFSET ?", pg.Limit, pg.Offset)
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
			Idx:   idx,
			Name:  name,
			Email: email,
		}
		Persons = append(Persons, p)
	}

	resopnse := Persons
	return c.JSON(http.StatusOK, resopnse)
}

// PageNum -
func PageNum(c echo.Context) error {
	db, err := conndb.ConnectToDb()
	if err != nil {
		return err
	}
	defer db.Close()

	var page int
	pg := new(model.Paging)
	pg.Page = page
	pg.Limit = 3
	if err := c.Bind(pg); err != nil {
		fmt.Println(err)
		return err
	}
	pg.Offset = (pg.Page - 1) * pg.Limit
	fmt.Println("page:", pg.Page, ", limit:", pg.Limit, ", offset:", pg.Offset)

	var Persons = []model.Person{}

	rows, err := db.Query("SELECT IDX, NAME, EMAIL FROM PERSON LIMIT ? OFFSET ?", pg.Limit, pg.Offset)
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
			Idx:   idx,
			Name:  name,
			Email: email,
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

	result, err2 := stmt.Exec(p.Name.String, p.Email.String, p.Idx.Int64)
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
