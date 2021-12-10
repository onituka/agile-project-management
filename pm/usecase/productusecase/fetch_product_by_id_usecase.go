package productusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/productdm"
)

type FetchProductByIDUsecase interface {
	FetchProductByID(ctx context.Context, in *FetchProductByIDInput) (*FetchProductByIDOutput, error)
}

type fetchProductByIDUsecase struct {
	productRepository productdm.ProductRepository
}

func NewFetchProductByIDUsecase(FetchProductByIDRepository productdm.ProductRepository) *fetchProductByIDUsecase {
	return &fetchProductByIDUsecase{
		productRepository: FetchProductByIDRepository,
	}
}

func (u *fetchProductByIDUsecase) FetchProductByID(ctx context.Context, in *FetchProductByIDInput) (*FetchProductByIDOutput, error) {
	productIDVo, err := productdm.NewProductID(in.ID)
	if err != nil {
		return nil, err
	}

	productDm, err := u.productRepository.FetchProductByID(ctx, productIDVo)
	if err != nil {
		return nil, err
	}

	return &FetchProductByIDOutput{
		ID:        productDm.ID().Value(),
		GroupID:   productDm.GroupID().Value(),
		Name:      productDm.Name().Value(),
		LeaderID:  productDm.LeaderID().Value(),
		CreatedAt: productDm.CreatedAt(),
		UpdatedAt: productDm.UpdatedAt(),
	}, nil
}
