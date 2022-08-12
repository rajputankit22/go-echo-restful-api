package handler

import (
	"fmt"
	"net/http"
	// "strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go-echo-restful-api/types"
	"go-echo-restful-api/db"

)

// GetProducts is func get all product
func GetProducts(c echo.Context) error {
	// data, err := db.GetAllProducts()
	// if err != nil {
	// 	return c.JSON(http.StatusNotFound, types.ParseStatus("NOT_FOUND", err.Error()))
	// }
	return c.JSON(http.StatusOK, "data")
}

// CreateProduct is func create new product
func CreateProduct(c echo.Context) error {
	var objRequest types.Product
	if err := c.Bind(&objRequest); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, types.ParseStatus("REQ_ERR", "Có lỗi xảy ra, vui lòng kiểm tra lại thông tin"))
	}
	fmt.Println("------------Product-----------",objRequest);
	fmt.Println("------------c.Validate(&objRequest)-----------",c.Validate(&objRequest));
	if err := c.Validate(&objRequest); err != nil {
		fmt.Println("------------err-----------",err);

		return c.JSON(http.StatusBadRequest, types.ParseStatus("REQ_INVALID", err.Error()))
	}

	data, err := db.CreateNewProduct(&objRequest)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, types.ParseStatus("NOT_ACCEPTED", err.Error()))
	}
	return c.JSON(http.StatusCreated, data)
}