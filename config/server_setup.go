package config

import (
	"github.com/Ceazzer/personal-website-server/cmd/http/handler"
	"github.com/labstack/echo/v4"
)

// TODO: Add Handlers as parameters
func ServerSetup() *echo.Echo {
	e := echo.New()

	// Handlers
	profileHandler := handler.NewProfileHandler(e, nil)
	messageHandler := handler.MessageHandler(e)

	// Middleware
	v := e.Group("/v1")

	// Routes
	// // Profile
	v.Group("/profile")
	e.GET("/", profileHandler.CreateProfileHandler)
	e.POST("/", profileHandler.GetProfileHandler)

	// // Message
	v.Group("/message")
	e.GET("/", messageHandler.CreateMessageHandler)

	return e
}
