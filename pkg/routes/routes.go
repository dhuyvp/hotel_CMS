package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func PublicRoutes(app *fiber.App, db *sqlx.DB) {
	GroupAPI := app.Group("/api")
	PublicHotelManageRoutes(app, db, GroupAPI)
	PublicChannelPackManageRoutes(app, db, GroupAPI)
	PublicChannelListManageRoutes(app, db, GroupAPI)
	PublicChannelManageRoutes(app, db, GroupAPI)
	PublicDeviceManageRoutes(app, db, GroupAPI)
}
