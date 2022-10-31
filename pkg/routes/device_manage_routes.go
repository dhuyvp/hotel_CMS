package routes

import (
	"hotel_cms/app/controllers"
	"hotel_cms/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PublicDeviceManageRoutes(app *fiber.App, db *sqlx.DB, GroupAPI fiber.Router) {
	route := GroupAPI.Group("/device_manage")

	// Data to INSERT
	HotelIDInsert := 1
	DeviceManageInsert := models.DeviceManage{
		DeviceID:          1,
		HotelID:           1,
		MacWired:          "",
		MacWireless:       "",
		DeviceName:        "device",
		DetailDescription: "",
		IsActive:          true,
	}

	// Data to SELECT
	var HotelIDSelect *int
	var IsActive *bool
	var ChannelName *string
	var DeviceName *string

	ValueHotelID := 1
	HotelIDSelect = &ValueHotelID
	ValueChannelName := ""
	ChannelName = &ValueChannelName
	ValueDeviceName := ""
	DeviceName = &ValueDeviceName

	ValueIsActive := true
	IsActive = &ValueIsActive

	// Data to UPDATE
	DeviceIdUpdate := 1
	DeviceManageUpdate := models.DeviceManage{
		DeviceID:          DeviceIdUpdate,
		HotelID:           1,
		MacWired:          "",
		MacWireless:       "",
		DeviceName:        "device",
		DetailDescription: "",
		IsActive:          true,
	}

	// HotelID to DELETE
	DeviceIdDelete := 2

	route.Post("/insert", controllers.InsertDeviceData(db, DeviceManageInsert, HotelIDInsert))
	route.Get("/get", controllers.SearchDeviceData(db, HotelIDSelect, ChannelName, DeviceName, IsActive))
	route.Put("/put", controllers.UpdateDeviceData(db, DeviceManageUpdate, DeviceIdUpdate))
	route.Delete("/delete", controllers.DeleteChannelData(db, DeviceIdDelete))

	route.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Successfull!")
	})
}
