package service

import (
	"context"

	pb "kx-boutique/api/user/v1"
	"kx-boutique/app/user/internal/biz"
)

type AuthService struct {
	uc *biz.AuthUsecase
	pb.UnimplementedAuthServer
}

func NewAuthService(uc *biz.AuthUsecase) *AuthService {
	return &AuthService{uc: uc}
}

func (s *AuthService) SignIn(ctx context.Context, req *pb.SignInRequest) (*pb.SignInReply, error) {
	token, err := s.uc.SignIn(ctx, &biz.SignInRequest{
		UserName: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &pb.SignInReply{
		Token: token,
	}, nil
}
func (s *AuthService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpReply, error) {
	token, err := s.uc.SignUp(ctx, &biz.SignUpRequest{
		UserName: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})

	if err != nil {
		return nil, err
	}

	return &pb.SignUpReply{
		Token: token,
	}, nil
}
