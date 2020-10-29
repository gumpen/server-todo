package usecase

import (
	"github.com/gumpen/server-todo/src/domain/model"
	"github.com/gumpen/server-todo/src/domain/repository"
)

type MessageUsecase interface {
	Index() (messages []*model.Message, err error)
	Create(message *model.Message) error
}

type messageUsecase struct {
	messageRepo repository.MessageRepository
}

func NewMessageUsecase(messageRepo repository.MessageRepository) MessageUsecase {
	messageUsecase := messageUsecase{messageRepo: messageRepo}
	return &messageUsecase
}

func (usecase *messageUsecase) Index() ([]*model.Message, error) {
	messages, err := usecase.messageRepo.FindAll()
	return messages, err
}

func (usecase *messageUsecase) Create(message *model.Message) error {
	_, err := usecase.messageRepo.Create(message)
	return err
}
