package productusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
)

type CreateProductUsecase interface {
	CreateProduct(ctx context.Context, in *productinput.CreateProductInput) (*productoutput.CreateProductOutput, error)
}

type createProductUsecase struct {
	productRepository productdm.ProductRepository
}

func NewCreateProductUsecase(ProductRepository productdm.ProductRepository) *createProductUsecase {
	return &createProductUsecase{
		productRepository: ProductRepository,
	}
}

func (u *createProductUsecase) CreateProduct(ctx context.Context, in *productinput.CreateProductInput) (*productoutput.CreateProductOutput, error) {
	groupIDVo, err := groupdm.NewGroupID(in.GroupID)
	if err != nil {
		return nil, err
	}

	nameVo, err := productdm.NewName(in.Name)
	if err != nil {
		return nil, err
	}

	leaderIDVo, err := userdm.NewUserID(in.LeaderID)
	if err != nil {
		return nil, err
	}

	productDm, err := productdm.GenProductForCreate(
		groupIDVo,
		nameVo,
		leaderIDVo,
	)
	if err != nil {
		return nil, err
	}

	productDomainService := productdm.NewProductDomainService(u.productRepository)

	exist, err := productDomainService.ExistsProductForCreate(ctx, productDm)
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return nil, err
	} else if exist {
		return nil, apperrors.Conflict
	}

	if err = u.productRepository.CreateProduct(ctx, productDm); err != nil {
		return nil, err
	}

	return &productoutput.CreateProductOutput{
		ID:        productDm.ID().Value(),
		GroupID:   productDm.GroupID().Value(),
		Name:      productDm.Name().Value(),
		LeaderID:  productDm.LeaderID().Value(),
		CreatedAt: productDm.CreatedAt(),
		UpdatedAt: productDm.UpdatedAt(),
	}, nil
}
