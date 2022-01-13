package producthandler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productinput"
)

type fetchProductByIDHandler struct {
	productUsecase productusecase.FetchProductByIDUsecase
}

func NewFetchProductByIDHandler(fetchProductByIDUsecase productusecase.FetchProductByIDUsecase) *fetchProductByIDHandler {
	return &fetchProductByIDHandler{
		productUsecase: fetchProductByIDUsecase,
	}
}

func (h *fetchProductByIDHandler) FetchProductByID(w http.ResponseWriter, r *http.Request) {
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

	in := productinput.FetchProductByIDInput{
		ID: productID,
	}

	out, err := h.productUsecase.FetchProductByID(r.Context(), &in)
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
