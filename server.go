package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type user struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var (
	users = map[int]*user{}
	seq = 1
)

/*
type Template struct {
	templates *template.Template
}
*/

func main() {
	e := echo.New()
	e.File("/", "public/index.html")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/*
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, aageboi!!")
	})
	*/
	e.POST("/users", saveUser)
	e.GET("/users",getUser)
	e.GET("/user/:id", getUser)
	e.PUT("/user/:id", updateUser)
	e.DELETE("/user/:id", deleteUser)
	
	e.GET("/show", show)

	// static file
	e.Static("/static", "assets")
	e.Logger.Fatal(e.Start(":1323"))
}

/**
* Handlers
*
*/
func getUser(c echo.Context) error{
	/*
	// show parameter id
	id := c.Param("id")
	return c.String(http.StatusOK, id)
	*/
	id, _ := strconv.Atoi(c.Param("id"))
	if id == 0 {
		return c.JSON(http.StatusOK, users)
	} else {
		return c.JSON(http.StatusOK, users[id])
	}
}

func saveUser(c echo.Context) error {
	/* 
	// using form
	name := c.FormValue("name")
	email := c.FormValue("email")

	return c.String(http.StatusOK, "name: "+name+", email: "+email)
	*/
	u := &user {
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

// /show?team=x-men&member=wolverine
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team: "+team+" member: "+member)
}
