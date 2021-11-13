package handler

import (
	"encoding/json"
	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecse"
	"net/http"
)

type productHandler struct {
	productUsecase productusecse.ProductUsecase
}

func NewProductHandler(productUsecase productusecse.ProductUsecase) *productHandler {
	return &productHandler{
		productUsecase: productUsecase,
	}
}

func (h *productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var in productusecse.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		setAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.productUsecase.CreateProduct(r.Context(), &in)
	if err != nil {
		setAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusCreated, out)
}
