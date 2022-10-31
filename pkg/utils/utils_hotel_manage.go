package utils

import (
	"fmt"
	"hotel_cms/app/models"
	"reflect"
)

func GetColumnsAndValuesHotelManage(DataStruct models.HotelManage) (string, string) {
	v := reflect.ValueOf(DataStruct)
	typeOfStruct := v.Type()

	queryColumns := ""
	queryValues := ""

	for i := 0; i < v.NumField(); i++ {
		queryColumns += "," + typeOfStruct.Field(i).Name
		if fmt.Sprintf("%s", v.Field(i).Type()) != "string" {
			queryValues += "," + fmt.Sprintf("%v", v.Field(i).Interface())
		} else {
			queryValues += ",'" + fmt.Sprintf("%v", v.Field(i).Interface()) + "'"
		}
	}

	return queryColumns[1:], queryValues[1:]
}

func GetQueryUpdateHotelManage(DataStruct models.HotelManage) string {
	v := reflect.ValueOf(DataStruct)
	typeOfStruct := v.Type()

	queryUpdate := ""

	for i := 0; i < v.NumField(); i++ {
		queryUpdate += "," + typeOfStruct.Field(i).Name + "="
		if fmt.Sprintf("%s", v.Field(i).Type()) != "string" {
			queryUpdate += fmt.Sprintf("%v", v.Field(i).Interface())
		} else {
			queryUpdate += "'" + fmt.Sprintf("%v", v.Field(i).Interface()) + "'"
		}
	}

	return queryUpdate[1:]
}
