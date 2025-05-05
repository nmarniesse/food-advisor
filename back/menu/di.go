package menu

import (
	"database/sql"
	"os"
)

type DI struct {
	connection             *sql.DB
	ConversationRepository ConversationRepository
	ia                     IAProvider
}

func NewDI() *DI {
	connection := CreateConnection()
	conversationRepository := GetConversationRepository(connection)

	isFake := os.Getenv("FAKE_AI") == "1"
	var ia IAProvider
	if isFake {
		ia = &Fake{}
	} else {
		ia = &ChatGPT{Token: os.Getenv("CHATGPT_TOKEN"), conversationRepository: conversationRepository}
	}

	return &DI{
		connection:             connection,
		ia:                     ia,
		ConversationRepository: conversationRepository,
	}
}

func (di *DI) Shutdown() {
	CloseConnection(di.connection)
}
