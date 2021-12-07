package producthandler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase"
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
	productID := mux.Vars(r)["productID"]

	in := productusecase.FetchProductByIDInput{
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
