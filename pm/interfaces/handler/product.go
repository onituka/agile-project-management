package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase"
)

type productHandler struct {
	productUsecase productusecase.ProductUsecase
}

func NewProductHandler(productUsecase productusecase.ProductUsecase) *productHandler {
	return &productHandler{
		productUsecase: productUsecase,
	}
}

func (h *productHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var in productusecase.CreateProductInput
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

func (h *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := mux.Vars(r)["productID"]

	in := productusecase.UpdateProductInput{
		ID: productID,
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		setAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.productUsecase.UpdateProduct(r.Context(), &in)
	if err != nil {
		setAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusCreated, out)

}
