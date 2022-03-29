package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Path segment
func pathSegment(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, "provided user id: "+id)
}

// Query parameter
func queryParameter(c echo.Context) error {
	year := c.QueryParam("year")
	return c.String(http.StatusOK, "requested book year: "+year)
}

// Header
func header(c echo.Context) error {
	token := c.Request().Header.Get("token")
	return c.String(http.StatusOK, "provided token: "+token)
}

/*
testing

http 'localhost:8088/auth' TOKEN:secured
http 'localhost:8088/book?year=2008'
http 'localhost:8088/user/123'

*/
func run() {
	e := echo.New()
	e.GET("/user/:id", pathSegment)
	e.GET("/book", queryParameter)
	e.GET("/auth", header)

	e.Logger.Fatal(e.Start(":8088"))
}
