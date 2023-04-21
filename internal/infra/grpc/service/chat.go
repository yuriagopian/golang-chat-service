package service

import "github.com/yuriagopian/golang-chat-service/internal/infra/grpc/pb"

type ChatService struct {
  pb.UnimplementedChatServiceServer
}