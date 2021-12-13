package productusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/productdm"
)

type FetchProductsUsecase interface {
	FetchProducts(ctx context.Context) (FetchProductsOutput, error)
}

type fetchProductsUsecase struct {
	productRepository productdm.ProductRepository
}

func NewFetchProductsUsecase(FetchProductsRepository productdm.ProductRepository) *fetchProductsUsecase {
	return &fetchProductsUsecase{
		productRepository: FetchProductsRepository,
	}
}

func (u *fetchProductsUsecase) FetchProducts(ctx context.Context) (FetchProductsOutput, error) {
	productsDm, err := u.productRepository.FetchProducts(ctx)
	if err != nil {
		return nil, err
	}

	productsDto := make(FetchProductsOutput, len(productsDm))
	for i, productDm := range productsDm {
		productsDto[i] = &Product{
			ID:        productDm.ID().Value(),
			GroupID:   productDm.GroupID().Value(),
			Name:      productDm.Name().Value(),
			LeaderID:  productDm.LeaderID().Value(),
			CreatedAt: productDm.CreatedAt(),
			UpdatedAt: productDm.UpdatedAt(),
		}
	}

	return productsDto, nil
}
