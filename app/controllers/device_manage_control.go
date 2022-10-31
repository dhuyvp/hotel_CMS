package controllers

import (
	"fmt"
	"hotel_cms/app/models"
	"hotel_cms/pkg/utils"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SearchDeviceData(db *sqlx.DB, HotelID *int, ChannelName *string, DeviceName *string, IsActive *bool) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("DEVICE"))
	tableHotelName := fmt.Sprintf("%s", os.Getenv("HOTEL"))

	return func(c *fiber.Ctx) error {
		queryDb := "SELECT " + tableName + ".* FROM " + tableName + "," + tableHotelName
		WHERECLause := " WHERE 1=1"
		if ChannelName != nil {
			WHERECLause += " AND DeviceName LIKE '%" + *ChannelName + "%'"
		}
		if IsActive != nil {
			WHERECLause += " AND IsActive=" + strconv.FormatBool(*IsActive)
		}

		var result []models.DeviceManage

		if HotelID == nil {
			log.Println("Don't have hotel that want to SEARCH device manage data")
			return c.Status(fiber.StatusOK).JSON(utils.Response{
				Success:    true,
				StatusCode: fiber.StatusOK,
				Data:       result,
			})
		}
		WHERECLause += " AND " + tableName + ".HotelID=" + tableHotelName + ".HotelID" +
			"AND " + tableName + ".HotelID=?"
		queryDb += WHERECLause

		err := db.Select(&result, queryDb, HotelID)

		if err != nil {
			log.Println("Error to SEARCH device manage data", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		return c.Status(fiber.StatusOK).JSON(utils.Response{
			Success:    true,
			StatusCode: fiber.StatusOK,
			Data:       result,
		})
	}
}

func InsertDeviceData(db *sqlx.DB, DataStruct models.DeviceManage, HotelID int) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("DEVICE"))
	tableHotelName := fmt.Sprintf("%s", os.Getenv("HOTEL"))

	queryColumns, queryValues := utils.GetColumnsAndValuesDeviceManage(DataStruct)
	return func(c *fiber.Ctx) error {
		var result []models.HotelManage
		queryHotel := "SELECT * FROM " + tableHotelName + " WHERE HotelID=?"

		err := db.Select(&result, queryHotel, HotelID)
		if err != nil {
			log.Println("Error to SELECT Hotel that want to INSERT device manage data", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		if result[0].DevicesNumber >= result[0].DevicesLimit {
			log.Println("Enough devices", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		queryDb := "INSERT INTO " + tableName + "(" + queryColumns + ") VALUES (" + queryValues + ")"
		_, errInsert := db.Exec(queryDb)

		if errInsert != nil {
			log.Println("Error to INSERT hotel manage data", errInsert)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		result[0].DevicesNumber += 1
		UpdateHotelData(db, result[0], HotelID)

		return c.Status(fiber.StatusOK).JSON(utils.Response{
			Success:    true,
			StatusCode: fiber.StatusOK,
			Data:       DataStruct,
		})
	}
}

func UpdateDeviceData(db *sqlx.DB, DataStruct models.DeviceManage, UpdateID int) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
	queryUpdate := utils.GetQueryUpdateDeviceManage(DataStruct)
	return func(c *fiber.Ctx) error {
		queryDb := "UPDATE " + tableName + " SET " + queryUpdate + " WHERE DeviceID=?"
		_, errUpdate := db.Exec(queryDb, UpdateID)

		if errUpdate != nil {
			log.Println("Error to UPDATE device manage data", errUpdate)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		return c.Status(fiber.StatusOK).JSON(utils.Response{
			Success:    true,
			StatusCode: fiber.StatusOK,
			Data:       DataStruct,
		})
	}
}

func DeleteDeviceData(db *sqlx.DB, deleteID int) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
	return func(c *fiber.Ctx) error {
		queryDb := "DELETE FROM " + tableName + " WHERE DeviceID=?"
		querySelect := "SELECT * FROM " + tableName + " WHERE DeviceID=?"

		var result []models.DeviceManage
		err := db.Select(&result, querySelect, deleteID)

		if err != nil {
			log.Println("Error SELECT device manage data that needs to be DELETE", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		_, errDelete := db.Exec(queryDb, deleteID)

		if errDelete != nil {
			log.Println("Error to DELETE device manage data", errDelete)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		return c.Status(fiber.StatusOK).JSON(utils.Response{
			Success:    true,
			StatusCode: fiber.StatusOK,
			Data:       result,
		})
	}
}
