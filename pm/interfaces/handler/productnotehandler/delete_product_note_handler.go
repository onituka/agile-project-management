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

type deleteProductNoteHandler struct {
	productNoteUsecase productnoteusecase.DeleteProductNoteUsecase
}

func NewDeleteProductNoteHandler(deleteProductNoteUsecase productnoteusecase.DeleteProductNoteUsecase) *deleteProductNoteHandler {
	return &deleteProductNoteHandler{
		productNoteUsecase: deleteProductNoteUsecase,
	}
}

func (h *deleteProductNoteHandler) DeleteProductNote(w http.ResponseWriter, r *http.Request) {
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

	in := productnoteinput.DeleteProductNoteInput{
		ID:        productNoteID,
		ProductID: productID,
	}

	if err := h.productNoteUsecase.DeleteProductNote(r.Context(), &in); err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusNoContent, nil)
}
