package router

import (
	"errors"
	"fmt"
	"net/http"
	h "qwik/helper"
	"qwik/model"
	s "qwik/server"
	"strconv"

	"github.com/labstack/echo/v4"
)

var LocalServe s.LocalService

//CreateData handles the POST request from client
func CreateData(c echo.Context) error {
	create := model.ServerData{}
	if err := c.Bind(&create); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Failed to Bind Data")
	}
	if err := LocalServe.Insert(create); err != nil {
		if err.Error() == `pq: duplicate key value violates unique constraint "server_data_key_key"` {
			return c.String(http.StatusBadRequest, h.ErrResponse(strconv.Itoa(http.StatusBadRequest),
				"Failed", err.Error()))
		}
		fmt.Println(err)
		return c.String(http.StatusBadRequest, h.ErrResponse(strconv.Itoa(http.StatusBadRequest),
			"Failed to Insert postgres", err.Error()))
	}
	_, err := LocalServe.SetRedis(create.Key, create.Value)
	if err != nil {
		return c.String(http.StatusBadRequest, h.ErrResponse(strconv.Itoa(http.StatusBadRequest),
			"Failed to Store the Data in redis", err.Error()))
	}
	return c.String(http.StatusOK, h.Response(strconv.Itoa(http.StatusOK),
		"Sucessfuly Inserted", "{"+create.Key+":"+create.Value+"}"))
}

//GetData handles the GET request from client
func GetData(c echo.Context) error {
	key := c.QueryParam("key")
	v, err := LocalServe.GetRedis(key)
	if v == "" || err != nil {
		rr, value := LocalServe.GET(key)
		if rr != nil {
			return c.String(http.StatusBadRequest,
				h.ErrResponse(strconv.Itoa(http.StatusBadRequest),
					"Failed to retrieve data from Local Data-Base", err.Error()))

		} else {
			_, er := LocalServe.SetRedis(key, value)
			if er != nil {
				return c.String(http.StatusBadRequest,
					h.ErrResponse(strconv.Itoa(http.StatusBadRequest),
						"Failed to SeT data from Local Data-Base to Redis", err.Error()))
			}
			return c.String(http.StatusOK, h.Response(strconv.Itoa(http.StatusOK),
				"Retrieve from Local Data-Base-Server", `{`+key+`:`+value+`}`))
		}
	}
	return c.String(http.StatusOK, h.Response(strconv.Itoa(http.StatusOK),
		"Retrieve from Redis-Server", `{`+key+`:`+v+`}`))
}

//DeleteData handles the DELETE request from client
func DeleteData(c echo.Context) error {
	delKey := c.QueryParam("delete-key")
	err := LocalServe.DelRedis(delKey)
	if err != nil {
		if err.Error() == errors.New("Key Does Not Exist").Error() {
			return c.String(http.StatusBadRequest,
				h.ErrResponse(strconv.Itoa(http.StatusBadRequest),
					"Failed to Delete", err.Error()))
		} else {
			return c.String(http.StatusBadRequest,
				h.ErrResponse(strconv.Itoa(http.StatusBadRequest),
					"Failed to Delete", err.Error()))
		}
	}
	return c.String(http.StatusOK, h.Response(strconv.Itoa(http.StatusOK),
		"Sucessfuly Deleted", "OK"))
}

//UpdateData handles the PUT request from client
func UpdateData(c echo.Context) error {
	updateKey := c.QueryParam("update-key")
	updateValue := c.QueryParam("update-value")
	err := LocalServe.Update(updateKey, updateValue)
	if err != nil {
		if err.Error() == "record not found" {
			return c.String(http.StatusBadRequest,
				h.ErrResponse(strconv.Itoa(http.StatusBadRequest),
					"Failed", err.Error()))
		} else {
			return c.String(http.StatusBadRequest,
				h.ErrResponse(strconv.Itoa(http.StatusBadRequest),
					"Failed to Update Data", err.Error()))
		}
	}

	val, err := LocalServe.SetRedis(updateKey, updateValue)
	if err != nil {
		return c.String(http.StatusBadRequest,
			h.ErrResponse(strconv.Itoa(http.StatusBadRequest),
				"Failed", err.Error()))
	}
	return c.String(http.StatusOK, h.Response(strconv.Itoa(http.StatusOK),
		val, `{`+updateKey+`:`+updateValue+`}`))
}
