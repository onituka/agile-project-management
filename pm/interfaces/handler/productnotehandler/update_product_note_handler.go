package productnotehandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
)

type updateProductNoteHandler struct {
	productNoteUsecase productnoteusecase.UpdateProductNoteUsecase
}

func NewUpdateProductNoteHandler(updateProductNoteUsecase productnoteusecase.UpdateProductNoteUsecase) *updateProductNoteHandler {
	return &updateProductNoteHandler{
		productNoteUsecase: updateProductNoteUsecase,
	}
}

func (h *updateProductNoteHandler) UpdateProductNote(w http.ResponseWriter, r *http.Request) {
	rv := mux.Vars(r)
	if rv == nil {
		handler.SetAppErrorToCtx(r, apperrors.InternalServerError)
		presenter.ErrorJSON(w, apperrors.InternalServerError)
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

	in := productnoteinput.UpdateProductNoteInput{
		ProductID: productID,
		ID:        productNoteID,
	}

	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	out, err := h.productNoteUsecase.UpdateProductNote(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusCreated, out)
}
