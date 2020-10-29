package repository

import "github.com/gumpen/server-todo/src/domain/model"

type MessageRepository interface {
	FindAll() (messages []*model.Message, err error)
	Create(*model.Message) (message *model.Message, err error)
}
