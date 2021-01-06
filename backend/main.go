package main

import (
	"os"
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/books/:id", func (c echo.Context) error {
		return c.String(http.StatusOK, c.Param("id"))
	})

	// サーバ起動
	e.Start(":" + os.Getenv("PORT"))
}
