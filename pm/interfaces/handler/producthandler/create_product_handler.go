package producthandler

import (
	"encoding/json"
	"net/http"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productinput"
)

type createProductHandler struct {
	productUsecase productusecase.CreateProductUsecase
}

func NewCreateProductHandler(createProductUsecase productusecase.CreateProductUsecase) *createProductHandler {
	return &createProductHandler{
		productUsecase: createProductUsecase,
	}
}

func (h *createProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var in productinput.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.productUsecase.CreateProduct(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusCreated, out)
}
