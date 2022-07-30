package handler

import (
	"net/http"
	// "strconv"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/gommon/log"
)

// GetProducts is func get all product
func GetProducts(c echo.Context) error {
	// data, err := db.GetAllProducts()
	// if err != nil {
	// 	return c.JSON(http.StatusNotFound, types.ParseStatus("NOT_FOUND", err.Error()))
	// }
	return c.JSON(http.StatusOK, "data")
}