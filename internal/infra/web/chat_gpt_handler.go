package web

import "github.com/yuriagopian/golang-chat-service/internal/usecases/chatcompletion"

type WebChatGPTHandler struct {
	CompletionUseCase chatcompletion.ChatCompletionUseCase
	Config            chatcompletion.ChatCompletionConfigInputDTO
	AuthToken         string
}
