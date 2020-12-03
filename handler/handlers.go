package handler

import (
	"github.com/labstack/echo/v4"
	"r1wallet/services"
)

type BaseHandler struct {
	Credit Credit
}

type Credit interface {
	HandleIncreaseRequestFromChannel(requestChannel chan string)
	HandleGetBalanceRequest() func(c echo.Context) error
}

func NewBaseHandler(services *services.Services) *BaseHandler {
	return &BaseHandler{
		Credit: newCreditHandler(services),
	}
}
