package menu

import "os"

type DI struct {
	ConversationRepository ConversationRepository
	ia                     IAProvider
}

func NewDI() *DI {
	conversationRepository := GetConversationRepository()

	isFake := os.Getenv("FAKE_AI") == "1"
	var ia IAProvider
	if isFake {
		ia = &Fake{}
	} else {
		ia = &ChatGPT{Token: os.Getenv("CHATGPT_TOKEN"), conversationRepository: conversationRepository}
	}

	return &DI{
		ia:                     ia,
		ConversationRepository: conversationRepository,
	}
}
