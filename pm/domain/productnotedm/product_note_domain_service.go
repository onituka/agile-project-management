package productnotedm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type productNoteDomainService struct {
	productNoteRepository ProductNoteRepository
}

func NewProductNoteDomainService(productNoteRepository ProductNoteRepository) *productNoteDomainService {
	return &productNoteDomainService{
		productNoteRepository: productNoteRepository,
	}
}

func (s *productNoteDomainService) ExistsProductNoteForCreate(ctx context.Context, productNoteDm *ProductNote) (bool, error) {
	existingProductNoteDm, err := s.productNoteRepository.FetchProductNoteByProductIDAndTitle(ctx, productNoteDm.ProductID(), productNoteDm.Title())
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return false, err
	} else if existingProductNoteDm != nil {
		return true, nil
	}

	return false, err
}
