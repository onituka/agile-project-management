package productusecse

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/timemanager"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, in *CreateProductInput) (*CreateProductOutput, error)
	UpdateProduct(ctx context.Context, in *UpdateProductInput) (*UpdateProductOutput, error)
}

type productUsecase struct {
	productRepository productdm.ProductRepository
	timeManager       timemanager.TimeManager
}

func NewProductUsecase(ProductRepository productdm.ProductRepository, timeManager timemanager.TimeManager) *productUsecase {
	return &productUsecase{
		productRepository: ProductRepository,
		timeManager:       timeManager,
	}
}

func (u *productUsecase) CreateProduct(ctx context.Context, in *CreateProductInput) (*CreateProductOutput, error) {
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

func (u *productUsecase) UpdateProduct(ctx context.Context, in *UpdateProductInput) (*UpdateProductOutput, error) {
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

	productDm.ChangeUpdatedAt(u.timeManager.Now())

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
