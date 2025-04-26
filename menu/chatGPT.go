package menu

import (
	"context"
	"fmt"
	"log"

	"github.com/sashabaranov/go-openai"
)

type ChatGPT struct {
	Token string
}

func (c *ChatGPT) RunQuery(query *Query) ([]Menu, error) {
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

	fmt.Println(resp.Choices[0].Message.Content)
	res := make([]Menu, 0)

	return res, nil
}
