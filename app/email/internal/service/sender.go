package service

import (
	"context"

	pb "kx-boutique/api/email/v1"
)

type SenderService struct {
	pb.UnimplementedSenderServer
}

func NewSenderService() *SenderService {
	return &SenderService{}
}

func (s *SenderService) SendMail(ctx context.Context, req *pb.SendMailRequest) (*pb.SendMailReply, error) {
	return &pb.SendMailReply{}, nil
}
