package handler

import "github.com/labstack/echo/v4"

type MessageHandlerResponse struct {
	CreateMessageHandler func(c echo.Context) error
}

func MessageHandler(e *echo.Echo) *MessageHandlerResponse {

	cmh := func(c echo.Context) error {
		return nil
	}

	return &MessageHandlerResponse{
		CreateMessageHandler: cmh,
	}
}
