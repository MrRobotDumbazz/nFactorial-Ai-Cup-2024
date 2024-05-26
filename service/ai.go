package service

import (
	"context"
	"fmt"
	"os"

	"github.com/neurocult/agency"
	"github.com/neurocult/agency/providers/openai"
)

type Ai interface {
	WriteMessageChatGPT(message string) ([]string, error)
	// ReadFileToChatGPT()
}

func WriteMessageChatGPT(message string) ([]agency.Message, error) {
	const op = "delivery.WriteMessageChatGPT"
	assistant := openai.
		New(openai.Params{Key: os.Getenv("OPENAI_API_KEY")}).
		TextToText(openai.TextToTextParams{Model: "gpt-3.5-turbo"}).
		SetPrompt("You are helpful assistant in the golang. You checking code and help people")

	messages := []agency.Message{}
	ctx := context.Background()

	input := agency.UserMessage(message)
	messages = append(messages, input)

	answer, err := assistant.SetMessages(messages).Execute(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	messages = append(messages, answer)
	return messages, nil
}
