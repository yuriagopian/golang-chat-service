package chatcompletionstream

import (
	"context"
	"errors"

	openai "github.com/sashabaranov/go-openai"
	"github.com/yuriagopian/golang-chat-service/internal/domain/entity"
	"github.com/yuriagopian/golang-chat-service/internal/domain/gateway"
)

type ChatCompletionConfigInputDTO struct {
	Model                string
	ModelMaxTokens       int
	Temperature          float32
	TopP                 float32
	N                    int
	Stop                 []string
	MaxTokens            int
	PresencePenalty      float32
	FrequencyPenalty     float32
	InitialSystemMessage string
}

type ChatCompletionInputDTO struct {
	ChatID      string
	UserID      string
	UserMessage string
	Config      ChatCompletionConfigInputDTO
}

type ChatCompletionOutputDTO struct {
	ChatID  string
	UserID  string
	Content string
}

type ChatCompletionUseCase struct {
	ChatGateway  gateway.ChatGateway
	OpenAiClient *openai.Client
}

func NewChatCompletionUseCase(chatGateway gateway.ChatGateway, openAiClient *openai.Client) *ChatCompletionUseCase {
	return &ChatCompletionUseCase{
		ChatGateway:  chatGateway,
		OpenAiClient: openAiClient,
	}
}

func (uc *ChatCompletionUseCase) Execute(ctx context.Context, input ChatCompletionInputDTO) (*ChatCompletionOutputDTO, error) {
	chat, err := uc.ChatGateway.FindChatByID(ctx, input.ChatID)

	if err != nil {
		if err.Error() == "chat not found" {
			// Create new chat(entity)
			chat, err = CreateNewChat(input)

			if err != nil {
				return nil, errors.New("error creating chat: " + err.Error())
			}
			// Save on database
			err = uc.ChatGateway.CreateChat(ctx, chat)

			if err != nil {
				return nil, errors.New("error persisting new chat: " + err.Error())
			}
		} else {
			return nil, errors.New("error fetching existing chat: " + err.Error())
		}
	}
}

func CreateNewChat(input ChatCompletionInputDTO) (*entity.Chat, error) {
	model := entity.NewModel(input.Config.Model, input.Config.ModelMaxTokens)
	chatConfig := &entity.ChatConfig{
		Temperature:      input.Config.Temperature,
		TopP:             input.Config.TopP,
		N:                input.Config.N,
		Stop:             input.Config.Stop,
		MaxTokens:        input.Config.MaxTokens,
		PresencePenalty:  input.Config.PresencePenalty,
		FrequencyPenalty: input.Config.FrequencyPenalty,
		Model:            model,
	}

	initialMessage, err := entity.NewMessage("system", input.Config.InitialSystemMessage, model)

	if err != nil {
		return nil, errors.New("error creating initial message: " + err.Error())
	}

	chat, err := entity.NewChat(input.UserID, initialMessage, chatConfig)

	if err != nil {
		return nil, errors.New("error creating new chat: " + err.Error())
	}

	return chat, nil
}
