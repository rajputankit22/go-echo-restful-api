package main

import (
	"fmt"
	// "go-echo-restful-api/config"
	"go-echo-restful-api/db"
	"go-echo-restful-api/handler"
	"go-echo-restful-api/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main()  {
	e := echo.New()
	e.Validator = middlewares.InitCustomValidator()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())

	api := e.Group("/api/v1", serverHeader)
	api.GET("/products", handler.GetProducts)          // Returns all resources of this product
	api.POST("/products", handler.CreateProduct)       // Creates a resource of this product and stores the data you posted, then returns the ID
	
	fmt.Println("Ankit")

	err := db.Ping()
	if err != nil {
		logrus.Fatal(err)
	}

	// service start at port :9090
	err = e.Start(":9090")
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

