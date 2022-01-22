package productnotehandler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
)

type fetchProductNoteByIDHandler struct {
	productUsecase productnoteusecase.FetchProductNoteByIDUsecase
}

func NewFetchProductNoteByIDHandler(fetchProductNoteByIDUsecase productnoteusecase.FetchProductNoteByIDUsecase) *fetchProductNoteByIDHandler {
	return &fetchProductNoteByIDHandler{
		productUsecase: fetchProductNoteByIDUsecase,
	}
}

func (h *fetchProductNoteByIDHandler) FetchProductNoteByID(w http.ResponseWriter, r *http.Request) {
	rv := mux.Vars(r)
	if rv == nil {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}

	productID, ok := rv["productID"]
	if !ok {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}

	productNoteID, ok := rv["productNoteID"]
	if !ok {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
		return
	}

	in := productnoteinput.FetchProductNoteByIDInput{
		ID:        productNoteID,
		ProductID: productID,
	}

	out, err := h.productUsecase.FetchProductNoteByID(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
