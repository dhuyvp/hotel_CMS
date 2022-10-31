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

func SearchChannelData(db *sqlx.DB, ChannelName *string, ChannelListName *string, IsActive *bool) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("CHANNEL"))
	return func(c *fiber.Ctx) error {
		queryDb := "SELECT * FROM " + tableName
		WHERECLause := " WHERE 1=1"
		if ChannelName != nil {
			WHERECLause += " AND ChannelPackName LIKE '%" + *ChannelName + "%'"
		}
		if ChannelListName != nil {
			WHERECLause += " AND ChannelPackName LIKE '%" + *ChannelListName + "%'"
		}
		if IsActive != nil {
			WHERECLause += " AND IsActive=" + strconv.FormatBool(*IsActive)
		}

		queryDb += WHERECLause
		var result []models.ChannelManage
		err := db.Select(&result, queryDb)

		if err != nil {
			log.Println("Error to SEARCH channel manage data", err)
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

func InsertChannelData(db *sqlx.DB, DataStruct models.ChannelManage) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("CHANNEL"))
	queryColumns, queryValues := utils.GetColumnsAndValuesChannelManage(DataStruct)
	return func(c *fiber.Ctx) error {
		queryDb := "INSERT INTO " + tableName + "(" + queryColumns + ") VALUES (" + queryValues + ")"
		_, errInsert := db.Exec(queryDb)

		if errInsert != nil {
			log.Println("Error to INSERT channel manage data", errInsert)
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

func UpdateChannelData(db *sqlx.DB, DataStruct models.ChannelManage, UpdateID int) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("CHANNEL"))
	queryUpdate := utils.GetQueryUpdateChannelManage(DataStruct)
	return func(c *fiber.Ctx) error {
		queryDb := "UPDATE " + tableName + " SET " + queryUpdate + " WHERE ChannelID=?"
		_, errUpdate := db.Exec(queryDb, UpdateID)

		if errUpdate != nil {
			log.Println("Error to UPDATE channel manage data", errUpdate)
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

func DeleteChannelData(db *sqlx.DB, deleteID int) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("CHANNEL"))
	return func(c *fiber.Ctx) error {
		queryDb := "DELETE FROM " + tableName + " WHERE ChannelID=?"
		querySelect := "SELECT * FROM " + tableName + " WHERE ChannelID=?"

		var result []models.ChannelManage
		err := db.Select(&result, querySelect, deleteID)

		if err != nil {
			log.Println("Error SELECT channel manage data that needs to be DELETE", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		_, errDelete := db.Exec(queryDb, deleteID)

		if errDelete != nil {
			log.Println("Error to DELETE channel pack manage data", errDelete)
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
