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

func SearchChannelPackData(db *sqlx.DB, ChannelPackName *string, IsActive *bool) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("CHANNEL_PACK"))
	return func(c *fiber.Ctx) error {
		queryDb := "SELECT * FROM " + tableName
		WHERECLause := " WHERE 1=1"
		if ChannelPackName != nil {
			WHERECLause += " AND ChannelPackName LIKE '%" + *ChannelPackName + "%'"
		}
		if IsActive != nil {
			WHERECLause += " AND IsActive=" + strconv.FormatBool(*IsActive)
		}

		queryDb += WHERECLause
		var result []models.ChannelPackManage
		err := db.Select(&result, queryDb)

		if err != nil {
			log.Println("Error to SEARCH channel pack manage data", err)
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

func InsertChannelPackData(db *sqlx.DB, DataStruct models.ChannelPackManage) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("CHANNEL_PACK"))
	queryColumns, queryValues := utils.GetColumnsAndValuesChannelPackManage(DataStruct)
	return func(c *fiber.Ctx) error {
		queryDb := "INSERT INTO " + tableName + "(" + queryColumns + ") VALUES (" + queryValues + ")"
		_, errInsert := db.Exec(queryDb)

		if errInsert != nil {
			log.Println("Error to INSERT channel pack manage data", errInsert)
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

func UpdateChannelPackData(db *sqlx.DB, DataStruct models.ChannelPackManage, UpdateID int) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("CHANNEL_PACK"))
	queryUpdate := utils.GetQueryUpdateChannelPackManage(DataStruct)
	return func(c *fiber.Ctx) error {
		queryDb := "UPDATE " + tableName + " SET " + queryUpdate + " WHERE ChannelPackID=?"
		_, errUpdate := db.Exec(queryDb, UpdateID)

		if errUpdate != nil {
			log.Println("Error to UPDATE channel pack manage data", errUpdate)
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

func DeleteChannelPackData(db *sqlx.DB, deleteID int) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("CHANNEL_PACK"))
	return func(c *fiber.Ctx) error {
		queryDb := "DELETE FROM " + tableName + " WHERE ChannelPackID=?"
		querySelect := "SELECT * FROM " + tableName + " WHERE ChannelPackID=?"

		var result []models.ChannelPackManage
		err := db.Select(&result, querySelect, deleteID)

		if err != nil {
			log.Println("Error SELECT channel pack manage data that needs to be DELETE", err)
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
