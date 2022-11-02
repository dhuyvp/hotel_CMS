package routes

import (
	"hotel_cms/app/controllers"
	"hotel_cms/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PublicChannelListManageRoutes(app *fiber.App, db *sqlx.DB, GroupAPI fiber.Router) {
	route := GroupAPI.Group("/channel_list_manage")

	// Data to INSERT
	ChannelListManageInsert := models.ChannelListManage{
		ChannelPackID:     2,
		ChannelListName:   "ABC1",
		DetailDescription: "mo_ta12",
		SortOrder:         1,
	}

	// Data to SELECT
	var ChannelListName *string
	ValueChannelListName := "VTDV"
	ChannelListName = &ValueChannelListName

	// Data to UPDATE
	ChannelListIdUpdate := 5
	ChannelListManageUpdate := models.ChannelListManage{
		ChannelListID:     ChannelListIdUpdate,
		ChannelPackID:     2,
		ChannelListName:   "ABC5",
		DetailDescription: "mota5",
		SortOrder:         8,
	}

	// HotelID to DELETE
	ChannelListIdDelete := 2

	route.Post("/insert", controllers.InsertChannelListData(db, ChannelListManageInsert))
	route.Get("/get", controllers.SearchChannelListData(db, ChannelListName))
	route.Put("/put", controllers.UpdateChannelListData(db, ChannelListManageUpdate, ChannelListIdUpdate))
	route.Delete("/delete", controllers.DeleteChannelListData(db, ChannelListIdDelete))

	route.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Successfull!")
	})
}
