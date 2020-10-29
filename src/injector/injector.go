package injector

import (
	"github.com/gumpen/server-todo/src/domain/repository"
	"github.com/gumpen/server-todo/src/handler"
	"github.com/gumpen/server-todo/src/infra"
	"github.com/gumpen/server-todo/src/usecase"
)

func InjectDB() infra.SqlHandler {
	sqlHandler := infra.NewSqlHandler()
	return *sqlHandler
}

func InjectMessageRepository() repository.MessageRepository {
	sqlHandler := InjectDB()
	return infra.NewMessageRepository(sqlHandler)
}

func InjectMessageUsecase() usecase.MessageUsecase {
	messageRepo := InjectMessageRepository()
	return usecase.NewMessageUsecase(messageRepo)
}

func InjectMessageHandler() handler.MessageHandler {
	return handler.NewMessageHandler(InjectMessageUsecase())
}
