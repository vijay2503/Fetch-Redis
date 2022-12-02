package helper

import (
	"encoding/json"
	"fmt"
)

func Response(code, status string, resp interface{}) string {
	var response = struct {
		Code     string
		Status   string
		Response interface{}
	}{Code: code, Status: status, Response: resp}
	result, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return string(result)
}
func ErrResponse(code, status string, resp interface{}) string {
	var response = struct {
		Code   string
		Status string
		Error  interface{}
	}{Code: code, Status: status, Error: resp}
	result, err := json.Marshal(response)
	if err != nil {
		return err.Error()
	}
	return string(result)
}
