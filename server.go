package main // import "server"

import (
	"database/sql"
	// "fmt"
	// "log"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"

	// nullable "gopkg.in/guregu/null.v3"

	"service"
)

var (
	db *sql.DB
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/select", service.SelectPersons)
	e.POST("/insert", service.InsertPerson)
	e.PUT("/update", service.UpdatePerson)
	e.DELETE("/delete", service.DeletePersonByIdx)
	/* e.GET("/select", func(c echo.Context) error {
		rows, err := db.Query("SELECT IDX, NAME, EMAIL FROM person")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		persons := []Person{}

		log.Println(&rows)

		for rows.Next() {
			var name, email string
			var idx int

			err := rows.Scan(&idx, &name, &email)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(idx, name, email)

			p := Person{
				Idx:   idx,
				Name:  sql.NullString{String: name, Valid: true},
				Email: sql.NullString{String: email, Valid: true},
			}
			persons = append(persons, p)
		}

		response := persons
		return c.JSON(http.StatusOK, response)
	}) */

	e.Logger.Fatal(e.Start(":1323"))
}

// ConnectToDb - hahaha
/* func ConnectToDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:13306)/test")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db is connected")
	}
	return db, err
} */

// SelectPersons -
/* func SelectPersons(c echo.Context) error {
	db, err := ConnectToDb()
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

	Persons := []Person{}

	for rows.Next() {
		var idx nullable.Int
		var name nullable.String
		var email nullable.String

		err := rows.Scan(&idx, &name, &email)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(idx, name, email)

		p := Person{
			Idx: idx,
			Name : name,
			Email: email,
		}
		Persons = append(Persons, p)
	}

	response := Persons
	return c.JSON(http.StatusOK, response)
} */

// InsertPerson -
/* func InsertPerson(c echo.Context) error {
	db, err := ConnectToDb()
	if err != nil {
		return err
	}
	defer db.Close()

	p := new(Person)
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
} */

// UpdatePerson -
/* func UpdatePerson(c echo.Context) error {
	db, err := ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	p := new(Person)
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
} */

// DeletePersonByIdx -
/* func DeletePersonByIdx(c echo.Context) error {
	db, err := ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	p := new(Person)
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
} */
