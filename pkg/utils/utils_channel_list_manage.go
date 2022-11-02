package utils

import (
	"fmt"
	"hotel_cms/app/models"
	"reflect"
)

func GetColumnsAndValuesChannelListManage(DataStruct models.ChannelListManage) (string, string) {
	v := reflect.ValueOf(DataStruct)
	typeOfStruct := v.Type()

	queryColumns := ""
	queryValues := ""

	for i := 0; i < v.NumField(); i++ {
		ColType := fmt.Sprintf("%s", v.Field(i).Type())
		ColValue := fmt.Sprintf("%v", v.Field(i).Interface())

		if ColType == "string" && ColValue == "" {
			continue
		}
		if ColType == "int" && ColValue == "0" {
			continue
		}

		queryColumns += "," + typeOfStruct.Field(i).Name
		if ColType != "string" {
			queryValues += "," + ColValue
		} else {
			queryValues += ",'" + ColValue + "'"
		}
	}

	return queryColumns[1:], queryValues[1:]
}

func GetQueryUpdateChannelListManage(DataStruct models.ChannelListManage) string {
	v := reflect.ValueOf(DataStruct)
	typeOfStruct := v.Type()

	queryUpdate := ""

	for i := 0; i < v.NumField(); i++ {
		ColType := fmt.Sprintf("%s", v.Field(i).Type())
		ColValue := fmt.Sprintf("%v", v.Field(i).Interface())

		if ColType == "string" && ColValue == "" {
			continue
		}
		if ColType == "int" && ColValue == "0" {
			continue
		}

		queryUpdate += "," + typeOfStruct.Field(i).Name + "="
		if ColType != "string" {
			queryUpdate += ColValue
		} else {
			queryUpdate += "'" + ColValue + "'"
		}
	}

	return queryUpdate[1:]
}
