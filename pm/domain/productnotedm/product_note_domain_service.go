package productnotedm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
)

type productNoteDomainService struct {
	productNoteRepository ProductNoteRepository
}

func NewProductNoteDomainService(productNoteRepository ProductNoteRepository) *productNoteDomainService {
	return &productNoteDomainService{
		productNoteRepository: productNoteRepository,
	}
}

func (s *productNoteDomainService) ExistsProductNoteByIDForUpdate(ctx context.Context, productNoteIDVo ProductNoteID, productIDVo productdm.ProductID) (*ProductNote, error) {
	return s.productNoteRepository.FetchProductNoteByIDForUpdate(ctx, productNoteIDVo, productIDVo)
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

func (s *productNoteDomainService) ExistsProductNoteForUpdate(ctx context.Context, productNoteDm *ProductNote) (bool, error) {
	oldProductNoteDm, err := s.productNoteRepository.FetchProductNoteByID(ctx, productNoteDm.ID(), productNoteDm.productID)
	if err != nil {
		return false, err
	}

	productNoteDmByTitle, err := s.productNoteRepository.FetchProductNoteByProductIDAndTitle(ctx, productNoteDm.productID, productNoteDm.Title())
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return false, err
	}

	if productNoteDmByTitle != nil {
		if productNoteDm.Title().Equals(oldProductNoteDm.Title()) {
			return false, apperrors.NotFound
		}

		return true, nil
	}

	return false, apperrors.NotFound
}
