package main

import (
	"fmt"
	"net/http"

	"github.com/chonla/cacheman"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	cacheConfig := &cacheman.Config{
		Enabled: true,
		Verbose: true,
		AdditionalHeaders: map[string]string{
			"X-Cache":         "true",
			"X-Cache-Manager": "cacheman",
		},
		Paths: []string{
			"/",
			"/user/:name",
		},
	}
	cache, err := cacheman.NewBigCache(cacheConfig)
	if err == nil {
		fmt.Printf("%s is used", cache.Type())
		e.Use(cacheman.Middleware(cacheConfig, cache))
	} else {
		e.Logger.Error(err.Error())
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, fmt.Sprintf("Hello, %s!", name))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
