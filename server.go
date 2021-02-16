package main // import "server"

import (
	"database/sql"
	"model"

	// "fmt"
	// "log"

	// "net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	// nullable "gopkg.in/guregu/null.v3"

	"service"
)

var (
	db *sql.DB
)

func main() {
	e := echo.New()

	/* e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	}) */

	e.POST("/login", service.Login)

	e.GET("/", service.Accessible)

	r := e.Group("/restricted")

	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}

	r.Use(middleware.JWTWithConfig(config))
	r.GET("", service.Restricted)

	e.GET("/select", service.SelectPersons)
	e.POST("/paging", service.SelectWithPaging)
	e.POST("/pagenum", service.PageNum)
	e.POST("/insert", service.InsertPerson)
	e.PUT("/update", service.UpdatePerson)
	e.DELETE("/delete", service.DeletePersonByIdx)

	e.Logger.Fatal(e.Start(":1323"))
}
