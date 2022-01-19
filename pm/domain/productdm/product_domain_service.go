package productdm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type productDomainService struct {
	productRepository ProductRepository
}

func NewProductDomainService(productRepository ProductRepository) *productDomainService {
	return &productDomainService{
		productRepository: productRepository,
	}
}

func (s *productDomainService) ExistsProductByIDForUpdate(ctx context.Context, productIDVo ProductID) (*Product, error) {
	return s.productRepository.FetchProductByIDForUpdate(ctx, productIDVo)
}

func (s *productDomainService) ExistsProductForCreate(ctx context.Context, productDm *Product) (bool, error) {
	existingProductDm, err := s.productRepository.FetchProductByGroupIDAndName(ctx, productDm.GroupID(), productDm.Name())
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return false, err
	} else if existingProductDm != nil {
		return true, nil
	}

	return false, err
}

func (s *productDomainService) ExistsProductForUpdate(ctx context.Context, productDm *Product) (bool, error) {
	oldProductDm, err := s.productRepository.FetchProductByID(ctx, productDm.ID())
	if err != nil {
		return false, err
	}

	productDmByName, err := s.productRepository.FetchProductByGroupIDAndName(ctx, productDm.GroupID(), productDm.Name())
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return false, err
	}

	if productDmByName != nil {
		if productDm.Name().Equals(oldProductDm.Name()) {
			return false, apperrors.NotFound
		}

		return true, nil
	}

	return false, apperrors.NotFound
}
