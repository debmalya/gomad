package main

import (
	"net/http"
        "os"	
        "io"	
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, do you hear echo!")
	})
        e.POST("/users", func(c echo.Context) error {
	    u := new(User)
	    if err := c.Bind(u); err != nil {
	        	return err
	    }
	    return c.JSON(http.StatusCreated, u)
	    // or
	    // return c.XML(http.StatusCreated, u)
        })
        e.GET("/users/:id", getUser)
        // e.PUT("/users/:id", updateUser)
        // e.DELETE("/users/:id", deleteUser)
        e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":1323"))
}


// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
  	// User ID from path `users/:id`
  	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}


// e.POST("/save", save)
func save(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
  	avatar, err := c.FormFile("avatar")
  	if err != nil {
 		return err
 	}
 
 	// Source
 	src, err := avatar.Open()
 	if err != nil {
 		return err
 	}
 	defer src.Close()
 
 	// Destination
 	dst, err := os.Create(avatar.Filename)
 	if err != nil {
 		return err
 	}
 	defer dst.Close()
 
 	// Copy
 	if _, err = io.Copy(dst, src); err != nil {
  		return err
  	}

	return c.HTML(http.StatusOK, "<b>Thank you! " + name + "</b>")
}
