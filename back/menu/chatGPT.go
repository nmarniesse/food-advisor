package menu

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
)

type ChatGPT struct {
	Token                  string
	conversationRepository ConversationRepository
}

func (c *ChatGPT) RunQuery(query *Query) (*Response, error) {
	client := openai.NewClient(c.Token)

	message := query.formatToString()
	log.Println("Query is about to start with message:\n" + message)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil, err
	}
	response := resp.Choices[0].Message.Content
	log.Println("ChatGPT response: " + response)

	var res Response
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	uuid := uuid.New()
	res.Uuid = uuid

	// @TODO: decouple conversation save from the IA
	c.conversationRepository.SaveConversation(&Conversation{
		Uuid: uuid,
		Messages: []ChatMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: message,
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: response,
			},
		},
	})
	log.Println("Conversation Saved", uuid)

	return &res, nil
}

func (c *ChatGPT) RunRefineQuery(query *RefineQuery) (*Response, error) {
	client := openai.NewClient(c.Token)

	conversation, err := c.conversationRepository.GetConversation(query.uuid)
	if err != nil {
		log.Println("Error getting conversation:", err)
		return nil, err
	}

	if conversation == nil {
		log.Println("Conversation not found", query.uuid)
		return nil, fmt.Errorf("conversation not found")
	}

	log.Println("Conversation retrieved:", conversation)

	newMessage := `Garde les recettes pour les jours suivants: %s. Et pour les autres jours, donne moi d'autres recettes.
Donne moi le résultat avec le même format JSON que précédemment. C'est important que le résultat soit un JSON valide.
	`
	newMessage = fmt.Sprintf(newMessage, strings.Join(query.daysToKeep, ", "))

	conversation.Messages = append(conversation.Messages, ChatMessage{Role: openai.ChatMessageRoleUser, Content: newMessage})

	var openaiMessages []openai.ChatCompletionMessage
	for _, msg := range conversation.Messages {
		openaiMessages = append(openaiMessages, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: openaiMessages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil, err
	}

	response := resp.Choices[0].Message.Content
	log.Println("ChatGPT response: " + response)

	var res Response
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}

	// @TODO: decouple conversation save from the IA
	conversation.Messages = append(conversation.Messages, ChatMessage{Role: openai.ChatMessageRoleSystem, Content: response})
	c.conversationRepository.SaveConversation(conversation)

	return &res, nil
}
