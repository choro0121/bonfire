package main

import (
	"os"
	"net/http"
	"github.com/labstack/echo"

	"github.com/lib/pq"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type User struct {
	UserId   int    `gorm:"primary_key; auto_increment"`
	Username string `gorm:"unique"`
}

var db *gorm.DB

func main() {
	// commect postgres
	connection, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})

	// create echo
	e := echo.New()

	e.GET("/books/:id", func (c echo.Context) error {
		return c.String(http.StatusOK, c.Param("id"))
	})

	e.POST("/users/:name", func (c echo.Context) error {
		user := &User{Username: c.Param("name")}
		db.Create(user)
		return c.JSON(http.StatusOK, user)
	})
	e.GET("/users", func (c echo.Context) error {
		var users []User
		db.Model(&User{}).Find(&users)
		return c.JSON(http.StatusOK, users)
	})

	// start server
	e.Start(":" + os.Getenv("PORT"))
}
