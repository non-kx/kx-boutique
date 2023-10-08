package service

import (
	"context"

	"kx-boutique/app/shop/internal/biz"

	pb "kx-boutique/api/shop/v1"
)

type ProductsService struct {
	uc *biz.ProductUsecase
	pb.UnimplementedProductsServer
}

// mustEmbedUnimplementedProductsServer implements v1.ProductsServer.
func (*ProductsService) mustEmbedUnimplementedProductsServer() {
	panic("unimplemented")
}

func NewProductsService(uc *biz.ProductUsecase) *ProductsService {
	return &ProductsService{
		uc: uc,
	}
}

func (s *ProductsService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductReply, error) {
	product, err := s.uc.GetProductById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetProductReply{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		ImageUrl:    &product.ImageUrl,
	}, nil
}
func (s *ProductsService) GetProductsPaginate(ctx context.Context, req *pb.GetProductsPaginateRequest) (*pb.GetProductsPaginateReply, error) {
	pgnate, err := s.uc.GetProductsPaginate(ctx, req.Page, req.Limit)
	if err != nil {
		return nil, err
	}

	reply := &pb.GetProductsPaginateReply{
		PageCount:  pgnate.PageCount,
		TotalCount: pgnate.TotalCount,
	}
	for _, p := range pgnate.Products {
		reply.Products = append(reply.Products, &pb.Product{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			ImageUrl:    &p.ImageUrl,
		})
	}

	return reply, nil
}
