package productusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/usecase/timemanager"
)

type TrashedProductUsecase interface {
	TrashedProduct(ctx context.Context, in *TrashedProductInput) (*TrashedProductOutput, error)
}

type trashedProductUsecase struct {
	productRepository productdm.ProductRepository
	projectRepository projectdm.ProjectRepository
	timeManager       timemanager.TimeManager
}

func NewTrashedProductUsecase(productRepository productdm.ProductRepository, projectRepository projectdm.ProjectRepository, timeManager timemanager.TimeManager) *trashedProductUsecase {
	return &trashedProductUsecase{
		productRepository: productRepository,
		projectRepository: projectRepository,
		timeManager:       timeManager,
	}
}

func (u *trashedProductUsecase) TrashedProduct(ctx context.Context, in *TrashedProductInput) (*TrashedProductOutput, error) {
	productIDVo, err := productdm.NewProductID(in.ID)
	if err != nil {
		return nil, err
	}

	productDm, err := u.productRepository.FetchProductByIDForUpdate(ctx, productIDVo)
	if err != nil {
		return nil, err
	} else if productDm.IsTrashed() {
		return nil, apperrors.Conflict
	}

	now := u.timeManager.Now()

	productDm.MoveToTrash()

	productDm.ChangeUpdatedAt(now)

	if err = u.productRepository.UpdateProduct(ctx, productDm); err != nil {
		return nil, err
	}

	projectDms, err := u.projectRepository.FetchProjectsNotInTrashByProductID(ctx, productDm.ID())
	if err != nil {
		return nil, err
	}

	// プロダクトに紐付くプロジェクト数をNとした時、N回分DBアクセスする事を許容する
	// プロジェクト数は多くならない想定
	for i, _ := range projectDms {
		projectDm := projectDms[i]

		if err = projectDm.ChangeTrashedAt(&now); err != nil {
			return nil, err
		}

		projectDm.ChangeUpdatedAt(now)

		if err = u.projectRepository.UpdateProject(ctx, projectDm); err != nil {
			return nil, err
		}
	}

	return &TrashedProductOutput{
		ID:        productDm.ID().Value(),
		GroupID:   productDm.GroupID().Value(),
		Name:      productDm.Name().Value(),
		LeaderID:  productDm.LeaderID().Value(),
		TrashedAt: productDm.TrashedAt(),
		CreatedAt: productDm.CreatedAt(),
		UpdatedAt: productDm.UpdatedAt(),
	}, nil
}
