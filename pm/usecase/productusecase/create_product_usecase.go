package productusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/timemanager"
)

type CreateProductUsecase interface {
	CreateProduct(ctx context.Context, in *CreateProductInput) (*CreateProductOutput, error)
}

type createProductUsecase struct {
	productRepository productdm.ProductRepository
	timeManager       timemanager.TimeManager
}

func NewCreateProductUsecase(ProductRepository productdm.ProductRepository, timeManager timemanager.TimeManager) *createProductUsecase {
	return &createProductUsecase{
		productRepository: ProductRepository,
		timeManager:       timeManager,
	}
}

func (u *createProductUsecase) CreateProduct(ctx context.Context, in *CreateProductInput) (*CreateProductOutput, error) {
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

	now := u.timeManager.Now()

	productDm, err := productdm.GenProductForCreate(
		groupIDVo,
		nameVo,
		leaderIDVo,
		now,
		now,
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

	return &CreateProductOutput{
		ID:        productDm.ID().Value(),
		GroupID:   productDm.GroupID().Value(),
		Name:      productDm.Name().Value(),
		LeaderID:  productDm.LeaderID().Value(),
		CreatedAt: productDm.CreatedAt(),
		UpdatedAt: productDm.UpdatedAt(),
	}, nil
}
