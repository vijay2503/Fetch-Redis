package router

import (
	"errors"
	"fmt"
	"net/http"
	"qwik/model"
	s "qwik/server"

	"github.com/labstack/echo/v4"
)

var LocalServe s.LocalService

func CreateData(c echo.Context) error {
	create := model.ServerData{}
	if err := c.Bind(&create); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Failed to Bind Data")
	}
	if err := LocalServe.Insert(create); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Failed to Insert Data to Postgres")
	}
	val, err := LocalServe.SetRedis(create.Key, create.Value)
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to Store the Data in redis")
	}

	return c.String(http.StatusOK, "Successfully Inserted\n"+val)
}
func GetData(c echo.Context) error {
retrieve:
	key := c.QueryParam("key")
	v, err := LocalServe.GetRedis(key)
	if v == "" || err != nil {
		rr, value := LocalServe.GET(key)
		if rr != nil {
			return c.String(http.StatusBadRequest, "Failed to retrieve data from Local Data-Base")
		}
		_, er := LocalServe.SetRedis(key, value)
		if er != nil {
			return c.String(http.StatusBadRequest, "Failed to SeT data from Local Data-Base to Redis")
		}
		goto retrieve
	}

	return c.String(http.StatusOK, "Retrieve from Redis-Server "+`{`+key+`:`+v+`}`)
}
func DeleteData(c echo.Context) error {
	delKey := c.QueryParam("delete-key")
	err := LocalServe.DelRedis(delKey)
	if err != nil {
		if err.Error() == errors.New("Key").Error() {
			return c.String(http.StatusBadRequest, "Deleted Key Does Not Exist")
		} else {
			return c.String(http.StatusBadRequest, "Failed to Delete the data")
		}
	}
	return c.String(http.StatusOK, "Success fully deleted the data")
}

func UpdateData(c echo.Context) error {
	updateKey := c.QueryParam("update-key")
	updateValue := c.QueryParam("update-value")
	err := LocalServe.Update(updateKey, updateValue)
	if err != nil {
		if err.Error() == "record not found" {
			return c.String(http.StatusBadRequest, "Updated Key Does Not Exist")
		} else {
			return c.String(http.StatusBadRequest, "Failed to Update the data")
		}
	}

	val, err := LocalServe.SetRedis(updateKey, updateValue)
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to Store the Data in redis")
	}

	return c.String(http.StatusOK, val+"Success fully Updated the data\n"+`{`+updateKey+`:`+updateValue+`}`)
}