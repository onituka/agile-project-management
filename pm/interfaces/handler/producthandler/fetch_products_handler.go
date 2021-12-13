package producthandler

import (
	"net/http"

	"github.com/onituka/agile-project-management/project-management/interfaces/handler"
	"github.com/onituka/agile-project-management/project-management/interfaces/presenter"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase"
)

type fetchProductsHandlere struct {
	productUsecase productusecase.FetchProductsUsecase
}

func NewFetchProductsHandler(fetchProductsUsecase productusecase.FetchProductsUsecase) *fetchProductsHandlere {
	return &fetchProductsHandlere{
		productUsecase: fetchProductsUsecase,
	}
}

func (h *fetchProductsHandlere) FetchProducts(w http.ResponseWriter, r *http.Request) {
	out, err := h.productUsecase.FetchProducts(r.Context())
	if err != nil {
		handler.SetAppErrorToCtx(r, err)
		presenter.ErrorJSON(w, err)
		return
	}

	presenter.JSON(w, http.StatusOK, out)
}
