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
	e.POST("/", profileHandler.CreateProfileHandler)
	e.GET("/", profileHandler.GetProfileHandler)

	// // Message
	v.Group("/message")
	e.POST("/", messageHandler.CreateMessageHandler)

	return e
}
