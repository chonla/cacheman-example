package main

import (
	"fmt"
	"net/http"

	"github.com/chonla/cacheman"
	"github.com/labstack/echo/v4"
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
		Server:        "localhost:6379",
		Password:      "",
		Database:      0,
		CacheInfoPath: "/cache/info",
	}
	cache, err := cacheman.NewRedis(cacheConfig)
	if err == nil {
		e.Use(cacheman.MiddlewareV4(cacheConfig, cache))
		fmt.Printf("%s is used", cache.Type())
	} else {
		e.Logger.Error(err.Error())
		fmt.Println(err.Error())
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
