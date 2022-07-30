package main

import (
	"go-echo-restful-api/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main()  {
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	api := e.Group("/api/v1", serverHeader)
	api.GET("/products", handler.GetProducts)          // Returns all resources of this product

	// service start at port :9090
	err := e.Start(":9090")
	if err != nil {
		logrus.Fatal(err)
	}
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("x-version", "Test/v1.0")
		return next(c)
	}
}

