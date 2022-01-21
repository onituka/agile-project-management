package productnotehandler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
)

type fetchProductNotesHandler struct {
	productNoteUsecase productnoteusecase.FetchProductNotesUsecase
}

func NewFetchProductNotesHandler(fetchProductNotesUsecase productnoteusecase.FetchProductNotesUsecase) *fetchProductNotesHandler {
	return &fetchProductNotesHandler{
		productNoteUsecase: fetchProductNotesUsecase,
	}
}

func (h *fetchProductNotesHandler) FetchProductNotes(w http.ResponseWriter, r *http.Request) {
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

	page, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 32)
	if err != nil {
		handler.SetAppErrorToCtx(r, apperrors.InvalidParameter)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	limit, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 32)
	if err != nil {
		handler.SetAppErrorToCtx(r, apperrors.InvalidParameter)
		presenter.ErrorJSON(w, apperrors.InvalidParameter)
		return
	}

	in := productnoteinput.FetchProductNotesInput{
		ProductID: productID,
		Page:      uint32(page),
		Limit:     uint32(limit),
	}

	out, err := h.productNoteUsecase.FetchProductNotes(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
