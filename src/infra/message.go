package infra

import (
	"github.com/gumpen/server-todo/src/domain/model"
	"github.com/gumpen/server-todo/src/domain/repository"
)

type MessageRepository struct {
	SqlHandler
}

func NewMessageRepository(sqlHandler SqlHandler) repository.MessageRepository {
	messageRepository := MessageRepository{sqlHandler}
	return &messageRepository
}

func (messageRepo *MessageRepository) FindAll() (messages []*model.Message, err error) {
	rows, err := messageRepo.SqlHandler.Conn.Query("SELECT * from messages")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		message := model.Message{}

		rows.Scan(&message.ID, &message.Body, &message.CreatedAt)

		messages = append(messages, &message)
	}
	return messages, nil
}

func (messageRepos *MessageRepository) Create(message *model.Message) (*model.Message, error) {
	_, err := messageRepos.SqlHandler.Conn.Exec("INSERT INTO messages (body, created_at) VALUES (?, ?)", message.Body, message.CreatedAt)
	return message, err
}
