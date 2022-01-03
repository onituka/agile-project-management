package productusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

type UpdateProductUsecase interface {
	UpdateProduct(ctx context.Context, in *UpdateProductInput) (*UpdateProductOutput, error)
}

type updateProductUsecase struct {
	productRepository productdm.ProductRepository
}

func NewUpdateProductUsecase(ProductRepository productdm.ProductRepository) *updateProductUsecase {
	return &updateProductUsecase{
		productRepository: ProductRepository,
	}
}

func (u *updateProductUsecase) UpdateProduct(ctx context.Context, in *UpdateProductInput) (*UpdateProductOutput, error) {
	productIDVo, err := productdm.NewProductID(in.ID)
	if err != nil {
		return nil, err
	}

	productDm, err := u.productRepository.FetchProductByIDForUpdate(ctx, productIDVo)
	if err != nil {
		return nil, err
	}

	nameVo, err := productdm.NewName(in.Name)
	if err != nil {
		return nil, err
	}

	productDm.ChangeName(nameVo)

	leaderIDVo, err := userdm.NewUserID(in.LeaderID)
	if err != nil {
		return nil, err
	}

	productDm.ChangeLeaderID(leaderIDVo)

	productDm.ChangeUpdatedAt()

	productDomainService := productdm.NewProductDomainService(u.productRepository)

	exist, err := productDomainService.ExistsProductForUpdate(ctx, productDm)
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return nil, err
	} else if exist {
		return nil, apperrors.Conflict
	}

	if err = u.productRepository.UpdateProduct(ctx, productDm); err != nil {
		return nil, err
	}

	return &UpdateProductOutput{
		ID:        productDm.ID().Value(),
		GroupID:   productDm.GroupID().Value(),
		Name:      productDm.Name().Value(),
		LeaderID:  productDm.LeaderID().Value(),
		CreatedAt: productDm.CreatedAt(),
		UpdatedAt: productDm.UpdatedAt(),
	}, nil
}
