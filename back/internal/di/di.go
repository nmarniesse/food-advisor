package di

import (
	"database/sql"
	"os"

	"github.com/nmarniesse/food-advisor/internal/model"
	"github.com/nmarniesse/food-advisor/internal/query"
	"github.com/nmarniesse/food-advisor/internal/storage"
)

type DI struct {
	Connection             *sql.DB
	ConversationRepository storage.ConversationRepository
	Ia                     model.IAProvider
}

func NewDI() *DI {
	connection := storage.CreateConnection()
	conversationRepository := storage.GetConversationRepository(connection)

	isFake := os.Getenv("FAKE_AI") == "1"
	var ia model.IAProvider
	if isFake {
		ia = &query.Fake{}
	} else {
		ia = &query.ChatGPT{Token: os.Getenv("CHATGPT_TOKEN"), ConversationRepository: conversationRepository}
	}

	return &DI{
		Connection:             connection,
		Ia:                     ia,
		ConversationRepository: conversationRepository,
	}
}

func (di *DI) Shutdown() {
	storage.CloseConnection(di.Connection)
}
