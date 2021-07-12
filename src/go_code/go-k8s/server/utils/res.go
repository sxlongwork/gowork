package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func GetJsonResponse(code int, msg string, data string, err error) (int, string) {
	var resultCode int
	dataRes := Response{}
	dataRes.Data = data
	if err != nil {
		dataRes.Code = http.StatusInternalServerError
		dataRes.Msg = "request failed"
		resultCode = http.StatusInternalServerError
	} else {
		dataRes.Code = code
		dataRes.Msg = msg
		resultCode = code
	}

	result, err := json.Marshal(dataRes)
	if err != nil {
		log.Printf("json marshal data failed, Eroor = %s\n", err.Error())
		return http.StatusInternalServerError, ""
	}
	return resultCode, string(result)

}

func GetJsonData(data interface{}, err error) (string, error) {
	if err != nil {
		return "", err
	}
	result, err := json.Marshal(data)
	if err != nil {
		log.Printf("json marshal data failed, Eroor = %s\n", err.Error())
		return "", err
	}
	return string(result), nil
}
