package product

import "context"

type Service interface {
	GetAllProducts(ctx context.Context) ([]Product, error)
}

type service struct {
	product Repository
}

func NewService(
	product Repository,
) Service {
	return &service{
		product: product,
	}
}

func (s *service) GetAllProducts(ctx context.Context) ([]Product, error) {
	return s.product.FindAll(ctx)
}
