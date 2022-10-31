package routes

import (
	"hotel_cms/app/controllers"
	"hotel_cms/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PublicChannelPackManageRoutes(app *fiber.App, db *sqlx.DB, GroupAPI fiber.Router) {
	route := GroupAPI.Group("/channel_pack_manage")

	// Data to INSERT
	ChannelPackManageInsert := models.ChannelPackManage{
		HotelID:           3,
		Logo:              "VTVCab",
		ChannelPackName:   "VTVCab",
		DetailDescription: "Tuyet1",
		Note:              "",
		IsActive:          true,
	}

	// Data to SELECT
	var ChannelPackName *string
	var IsActive *bool
	ValueChannelPackName := "VTDV"
	ValueIsActive := true
	ChannelPackName = &ValueChannelPackName
	IsActive = &ValueIsActive

	// Data to UPDATE
	ChannelPackIdUpdate := 1
	ChannelPackManageUpdate := models.ChannelPackManage{
		ChannelPackID:     ChannelPackIdUpdate,
		HotelID:           3,
		Logo:              "VTVCab",
		ChannelPackName:   "VTVCab",
		DetailDescription: "Tuyetvoi",
		Note:              "",
		IsActive:          true,
	}

	// HotelID to DELETE
	ChannelPackIdDelete := 2

	route.Post("/insert", controllers.InsertChannelPackData(db, ChannelPackManageInsert))
	route.Get("/get", controllers.SearchChannelPackData(db, ChannelPackName, IsActive))
	route.Put("/put", controllers.UpdateChannelPackData(db, ChannelPackManageUpdate, ChannelPackIdUpdate))
	route.Delete("/delete", controllers.DeleteChannelPackData(db, ChannelPackIdDelete))

	route.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Successfull!")
	})
}
