package apimodel

import (
	"reflect"
)

type HttpResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type JsonObjectArray struct {
	Keys    []string      `json:"keys"`
	Objects []interface{} `json:"objects"`
}

type ObjectArrayRes struct {
	Response HttpResponse    `json:"response"`
	Data     JsonObjectArray `json:"data"`
}

type PageRes struct {
	Response HttpResponse `json:"response"`
	Data     PageData     `json:"data"`
}
type PageData struct {
	Total   int           `json:"total"`
	Keys    []string      `json:"keys"`
	Objects []interface{} `json:"objects"`
}

func NewJsonObjectArray(origin interface{}) (jsonArray JsonObjectArray) {
	if reflect.TypeOf(origin).Kind() != reflect.Slice {
		return
	}
	originVal := reflect.ValueOf(origin)
	if originVal.Len() < 1 {
		return
	}
	first := originVal.Index(0)
	for i := 0; i < first.NumField(); i++ {
		jsonArray.Keys = append(jsonArray.Keys, first.Type().Field(i).Name)
	}

	for i := 0; i < originVal.Len(); i++ {
		objVal := originVal.Index(i)
		var jsonObj []interface{}
		for j := 0; j < objVal.NumField(); j++ {
			jsonObj = append(jsonObj, objVal.Field(j).Interface())
		}
		jsonArray.Objects = append(jsonArray.Objects, jsonObj)
	}
	return
}
