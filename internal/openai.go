package internal

import (
	"DiskusiTugas/config"
	"context"

	"github.com/sashabaranov/go-openai"
)

func SetupOpenAI(apiKey string) (*openai.Client, error) {
	client := openai.NewClient(apiKey)
	return client, nil
}

func ResponOpenAI(prompt string) (string, error) {
	client := openai.NewClient(config.InitConfig().OpenAIKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
