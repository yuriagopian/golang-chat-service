package service

import (
	"github.com/yuriagopian/golang-chat-service/internal/infra/grpc/pb"
	"github.com/yuriagopian/golang-chat-service/internal/usecases/chatcompletionstream"
)

type ChatService struct {
  pb.UnimplementedChatServiceServer
  ChatCompletionStreamUseCase chatcompletionstream.ChatCompletionUseCase
  ChatConfigStream chatcompletionstream.ChatCompletionConfigInputDTO
}