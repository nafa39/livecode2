package routes

import (
	"livecode2/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/users/register", controllers.Register)
	// e.POST("/users/login", controllers.Login)

	// // Protected routes
	// e.GET("/bookings", controllers.GetBookings, middlewares.IsAuthenticated)
	// e.GET("/bookings/unpaid", controllers.GetUnpaidBookings, middlewares.IsAuthenticated)
	// e.GET("/vehicles/availability", controllers.GetVehicleAvailability, middlewares.IsAuthenticated)
	// e.GET("/reports/total-customers", controllers.GetTotalCustomers, middlewares.IsAuthenticated)
	// e.GET("/reports/bookings-per-tour", controllers.GetBookingsPerTour, middlewares.IsAuthenticated)
}
