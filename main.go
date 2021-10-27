package main

import (
	"net/http"
	"tia"
)

func main() {
	r := tia.New()
	r.GET("/", func(c *tia.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Tin</h1>")
	})

	r.GET("/hello", func(c *tia.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *tia.Context) {
		c.JSON(http.StatusOK, tia.J{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
