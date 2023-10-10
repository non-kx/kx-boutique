package service

import (
	"context"

	pb "kx-boutique/api/shop/v1"
	"kx-boutique/app/shop/internal/biz"
)

type CartsService struct {
	uc *biz.CartUsecase
	pb.UnimplementedCartsServer
}

func NewCartsService(uc *biz.CartUsecase) *CartsService {
	return &CartsService{uc: uc}
}

func (s *CartsService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartReply, error) {
	cart, err := s.uc.GetUserCart(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	if cart == nil {
		return &pb.GetCartReply{}, nil
	}

	replyItems := make([]*pb.Item, 0)
	for _, item := range cart.Items {
		replyItems = append(replyItems, &pb.Item{
			ProductId: item.ProductId,
			Qty:       item.Qty,
		})
	}

	return &pb.GetCartReply{
		Cart: &pb.Cart{
			UserId: cart.UserId,
			Items:  replyItems,
		},
	}, nil
}
func (s *CartsService) AddToCart(ctx context.Context, req *pb.AddToCartReqeust) (*pb.AddToCartReply, error) {
	item := biz.Item{
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Qty,
	}

	err := s.uc.AddToCart(ctx, item, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.AddToCartReply{
		Message: "OK",
	}, nil
}
