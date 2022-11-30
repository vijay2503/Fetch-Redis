package router

import "github.com/labstack/echo/v4"

func Routing() *echo.Echo {
	e := echo.New()
	e.POST("/setData", CreateData)
	e.GET("/getData", GetData)
	e.PUT("/UpdateData", UpdateData)
	e.DELETE("/DeleteData", DeleteData)
	return e
}
