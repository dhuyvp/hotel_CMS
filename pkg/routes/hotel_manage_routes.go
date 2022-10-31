package routes

import (
	"hotel_cms/app/controllers"
	"hotel_cms/app/models"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PublicHotelManageRoutes(app *fiber.App, db *sqlx.DB, GroupAPI fiber.Router) {
	route := GroupAPI.Group("/hotel_manage")

	// Data to INSERT
	HotelManageInsert := models.HotelManage{
		Logo:              "abc",
		HotelName:         "Hotel2",
		DevicesLimit:      100,
		DevicesNumber:     0,
		TotalRoom:         10,
		Address:           "VN",
		DetailDescription: "Perfect",
		IsActive:          true,
	}

	// Data to SELECT
	var HotelName *string
	var IsActive *bool
	ValueHotelName := "Hotel1"
	ValueIsActive := true
	HotelName = &ValueHotelName
	IsActive = &ValueIsActive

	// Data to UPDATE
	HotelIdUpdate := 2
	HotelManageUpdate := models.HotelManage{
		HotelID:           HotelIdUpdate,
		Logo:              "abcd",
		HotelName:         "Hotel2",
		DevicesLimit:      100,
		DevicesNumber:     0,
		TotalRoom:         10,
		Address:           "VN",
		DetailDescription: "Perfect",
		IsActive:          true,
	}

	// HotelID to DELETE
	HotelIdDelete := 1

	route.Post("/insert", controllers.InsertHotelData(db, HotelManageInsert))
	route.Get("/get", controllers.SearchHotelData(db, HotelName, IsActive))
	route.Put("/put", controllers.UpdateHotelData(db, HotelManageUpdate, HotelIdUpdate))
	route.Delete("/delete", controllers.DeleteHotelData(db, HotelIdDelete))

	route.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Successfull!")
	})
}
