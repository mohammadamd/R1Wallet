package routes

import (
	"github.com/labstack/echo/v4"
	"r1wallet/handler"
)

func RegisterRoutes(e *echo.Echo, handler *handler.BaseHandler) {
	api := e.Group("/api")
	api.GET("/balance/:user_id", handler.Credit.HandleGetBalanceRequest())
}
