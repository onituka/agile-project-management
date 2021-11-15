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

func (s *productDomainService) ExistsProductForCreate(ctx context.Context, productDm *Product) (bool, error) {
	existingProductDm, err := s.productRepository.FetchProductByGroupIDAndName(ctx, productDm.GroupID(), productDm.Name())
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return false, err
	} else if existingProductDm != nil {
		return true, nil
	}

	return false, err
}
