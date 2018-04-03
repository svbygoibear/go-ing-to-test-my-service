package models

import (
	"github.com/labstack/echo"
	"net/http"
)

//// user in this case is a person who will be making use of something
//type User struct {
//	name 	string
//	surname	string
//	email	string
//	age 	int32
//}

// e.GET("/users/:id", getUser)
func GetUser(c echo.Context) error {
	id := c.Param("id") // User ID from path `users/:id`
	return c.String(http.StatusOK, id)
}