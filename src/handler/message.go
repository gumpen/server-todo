package handler

import (
	"net/http"

	"github.com/gumpen/server-todo/src/domain/model"
	"github.com/gumpen/server-todo/src/usecase"
	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	messageUsecase usecase.MessageUsecase
}

func NewMessageHandler(messageUsecase usecase.MessageUsecase) MessageHandler {
	messageHandler := MessageHandler{messageUsecase: messageUsecase}
	return messageHandler
}

func (handler *MessageHandler) Index() echo.HandlerFunc {
	return func(c echo.Context) error {
		models, err := handler.messageUsecase.Index()
		if err != nil {
			return c.JSON(http.StatusBadRequest, models)
		}
		return c.JSON(http.StatusOK, models)
	}
}

func (handler *MessageHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var message model.Message
		c.Bind(&message)
		err := handler.messageUsecase.Create(&message)
		return c.JSON(http.StatusOK, err)
	}
}
