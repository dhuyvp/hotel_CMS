package routes

import (
	"hotel_cms/app/controllers"
	"hotel_cms/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PublicChannelManageRoutes(app *fiber.App, db *sqlx.DB, GroupAPI fiber.Router) {
	route := GroupAPI.Group("/channel_manage")

	// Data to INSERT
	ChannelManageInsert := models.ChannelManage{
		ChannelListID:     4,
		Logo:              "logo",
		ChannelName:       "123",
		Link:              "dsa.dsad",
		DetailDescription: "duoc",
		IsActive:          true,
	}

	// Data to SELECT
	var ChannelName *string
	var IsActive *bool
	var ChannelListName *string

	ValueChannelName := "VTDV"
	ChannelName = &ValueChannelName
	ValueChannelListName := ""
	ChannelListName = &ValueChannelListName

	ValueIsActive := true
	IsActive = &ValueIsActive

	// Data to UPDATE
	ChannelID := 1
	ChannelIdUpdate := 1
	ChannelManageUpdate := models.ChannelManage{
		ChannelID:         ChannelIdUpdate,
		ChannelListID:     1,
		Logo:              "Logo",
		ChannelName:       "123",
		Link:              "dsa.dsad",
		DetailDescription: "duoc",
		IsActive:          true,
	}

	// HotelID to DELETE
	ChannelIdDelete := 2

	route.Post("/insert", controllers.InsertChannelData(db, ChannelManageInsert))
	route.Get("/get", controllers.SearchChannelData(db, ChannelName, ChannelListName, IsActive))
	route.Put("/put", controllers.UpdateChannelData(db, ChannelManageUpdate, ChannelID))
	route.Delete("/delete", controllers.DeleteChannelData(db, ChannelIdDelete))

	route.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Successfull!")
	})
}
