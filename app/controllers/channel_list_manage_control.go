package controllers

import (
	"fmt"
	"hotel_cms/app/models"
	"hotel_cms/pkg/utils"
	"log"
	"math"
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

		var result []models.ChannelListManage
		querySelect := "SELECT SortOrder FROM " + tableName
		err := db.Select(&result, querySelect)

		if err != nil {
			log.Println("Error to SELECT before INSERT channel list manage data", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		for i := 0; i < len(result)-1; i++ {
			RowValue := result[i]
			if RowValue.SortOrder >= DataStruct.SortOrder {
				query := "UPDATE " + tableName + " SET " + "SortOrder=? WHERE ChannelListID=?"
				_, errUpdate := db.Exec(query, RowValue.SortOrder+1, RowValue.ChannelListID)

				if errUpdate != nil {
					log.Println("Error to SELECT before INSERT channel list manage data", errUpdate)
					return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
						Success:    false,
						StatusCode: fiber.StatusBadRequest,
					})
				}
			}
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
	return func(c *fiber.Ctx) error {
		queryUpdate := utils.GetQueryUpdateChannelListManage(DataStruct)

		var result []models.ChannelListManage
		querySelect := "SELECT * FROM " + tableName
		err := db.Select(&result, querySelect)

		if err != nil {
			log.Println("Error to SELECT before UPDATE channel list manage data", err)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		var resultGet []models.ChannelListManage
		queryGetSortOrder := "SELECT SortOrder FROM " + tableName + " WHERE ChannelListID=?"
		errGet := db.Select(&resultGet, queryGetSortOrder, UpdateID)
		if errGet != nil {
			log.Println("Error to SELECT before UPDATE channel list manage data", errGet)
			return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		LeftOrder := int(math.Min(float64(resultGet[0].SortOrder), float64(DataStruct.SortOrder)))
		RightOrder := int(math.Max(float64(resultGet[0].SortOrder), float64(DataStruct.SortOrder)))

		count := 1
		if resultGet[0].SortOrder <= DataStruct.SortOrder {
			count = -1
		}

		for i := 0; i < len(result); i++ {
			RowValue := result[i]
			if RowValue.SortOrder >= LeftOrder && RowValue.SortOrder <= RightOrder {

				query := "UPDATE " + tableName + " SET " + "SortOrder=? WHERE ChannelListID=?"
				_, errUpdate := db.Exec(query, RowValue.SortOrder+count, RowValue.ChannelListID)

				if errUpdate != nil {
					log.Println("Error to SELECT before UPDATE channel list manage data", errUpdate)
					return c.Status(fiber.StatusBadRequest).JSON(utils.Response{
						Success:    false,
						StatusCode: fiber.StatusBadRequest,
					})
				}
			}
		}

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
