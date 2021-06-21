package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.NEW()

	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee<h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		name := c.Query("name")
		c.String(http.StatusOK, "hello %s, you're at %s\n", name, c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
