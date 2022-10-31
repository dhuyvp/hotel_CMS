package controllers

import (
	"fmt"
	"hotel_cms/app/models"
	"hotel_cms/pkg/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SearchChannelListData(db *sqlx.DB, ChannelListName *string) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("CHANNEL_LIST"))
	return func(c *fiber.Ctx) error {
		queryDb := "SELECT * FROM " + tableName
		WHERECLause := " WHERE 1=1"
		if ChannelListName != nil {
			WHERECLause += " AND ChannelPackName LIKE '%" + *ChannelListName + "%'"
		}

		queryDb += WHERECLause
		var result []models.ChannelListManage
		err := db.Select(&result, queryDb)

		if err != nil {
			log.Println("Error to SEARCH channel list manage data", err)
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

func InsertChannelListData(db *sqlx.DB, DataStruct models.ChannelListManage) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("CHANNEL_LIST"))
	queryColumns, queryValues := utils.GetColumnsAndValuesChannelListManage(DataStruct)
	return func(c *fiber.Ctx) error {
		queryDb := "INSERT INTO " + tableName + "(" + queryColumns + ") VALUES (" + queryValues + ")"
		_, errInsert := db.Exec(queryDb)

		if errInsert != nil {
			log.Println("Error to INSERT channel list manage data", errInsert)
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

func UpdateChannelListData(db *sqlx.DB, DataStruct models.ChannelListManage, UpdateID int) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("CHANNEL_LIST"))
	queryUpdate := utils.GetQueryUpdateChannelListManage(DataStruct)
	return func(c *fiber.Ctx) error {
		queryDb := "UPDATE " + tableName + " SET " + queryUpdate + " WHERE ChannelListID=?"
		_, errUpdate := db.Exec(queryDb, UpdateID)

		if errUpdate != nil {
			log.Println("Error to UPDATE channel list manage data", errUpdate)
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

func DeleteChannelListData(db *sqlx.DB, deleteID int) fiber.Handler {
	tableName := fmt.Sprintf("%s", os.Getenv("CHANNEL_LIST"))
	return func(c *fiber.Ctx) error {
		queryDb := "DELETE FROM " + tableName + " WHERE ChannelListID=?"
		querySelect := "SELECT * FROM " + tableName + " WHERE ChannelListID=?"

		var result []models.ChannelListManage
		err := db.Select(&result, querySelect, deleteID)

		if err != nil {
			log.Println("Error SELECT channel list manage data that needs to be DELETE", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		_, errDelete := db.Exec(queryDb, deleteID)

		if errDelete != nil {
			log.Println("Error to DELETE channel list manage data", errDelete)
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
