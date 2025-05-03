package menu

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/sashabaranov/go-openai"
)

type ChatGPT struct {
	Token string
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

	return &res, nil
}

func (c *ChatGPT) RunRefineQuery(query *RefineQuery) (*Response, error) {
	client := openai.NewClient(c.Token)

	/**
	Note: It doesn't work because chatGPT doesn't remember the previous context.
	To do that, we neet to keep the conversation history like that and give it to chatGPT:

	[
	  {"role": "system", "content": "You are a helpful assistant."},
	  {"role": "user", "content": "Hello, who won the world cup in 2018?"},
	  {"role": "assistant", "content": "France won the 2018 FIFA World Cup."},
	  {"role": "user", "content": "Who was the top scorer?"}
	]

	Each time we want to ask something, add the {"role": "user", "content": "<question>"} at the end of
	the list and send the whole content.
	*/

	message := `Peutx tu me donner d'autres propositions en gardant uniquement les jours suivants: %s.
Donne moi le résultat avec le même format que précédemment.
	`
	message = fmt.Sprintf(message, strings.Join(query.DaysToKeep, ", "))

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

	return &res, nil
}
