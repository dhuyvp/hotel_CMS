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

func SearchHotelData(db *sqlx.DB, HotelName *string, IsActive *bool) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
	return func(c *fiber.Ctx) error {
		queryDb := "SELECT * FROM " + tableName
		WHERECLause := " WHERE 1=1"
		if HotelName != nil {
			WHERECLause += " AND HotelName LIKE '%" + *HotelName + "%'"
		}
		if IsActive != nil {
			WHERECLause += " AND IsActive=" + strconv.FormatBool(*IsActive)
		}

		queryDb += WHERECLause
		var result []models.HotelManage
		err := db.Select(&result, queryDb)

		if err != nil {
			log.Println("Error to SEARCH hotel manage data", err)
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

func InsertHotelData(db *sqlx.DB, DataStruct models.HotelManage) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
	queryColumns, queryValues := utils.GetColumnsAndValuesHotelManage(DataStruct)
	return func(c *fiber.Ctx) error {
		queryDb := "INSERT INTO " + tableName + "(" + queryColumns + ") VALUES (" + queryValues + ")"
		_, errInsert := db.Exec(queryDb)

		if errInsert != nil {
			log.Println("Error to INSERT hotel manage data", errInsert)
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

func UpdateHotelData(db *sqlx.DB, DataStruct models.HotelManage, UpdateID int) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
	queryUpdate := utils.GetQueryUpdateHotelManage(DataStruct)
	return func(c *fiber.Ctx) error {
		queryDb := "UPDATE " + tableName + " SET " + queryUpdate + " WHERE HotelID=?"
		_, errUpdate := db.Exec(queryDb, UpdateID)

		if errUpdate != nil {
			log.Println("Error to UPDATE hotel manage data", errUpdate)
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

func DeleteHotelData(db *sqlx.DB, deleteID int) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("HOTEL"))
	return func(c *fiber.Ctx) error {
		queryDb := "DELETE FROM " + tableName + " WHERE HotelID=?"
		querySelect := "SELECT * FROM " + tableName + " WHERE HotelID=?"

		var result []models.HotelManage
		err := db.Select(&result, querySelect, deleteID)

		if err != nil {
			log.Println("Error SELECT hotel manage data that needs to be DELETE", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		_, errDelete := db.Exec(queryDb, deleteID)

		if errDelete != nil {
			log.Println("Error to DELETE hotel manage data", errDelete)
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
